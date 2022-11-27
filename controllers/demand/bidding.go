package demand

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/seller-app/auction/controllers/supplier"
	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
	"github.com/seller-app/auction/utils"
)

func InsertBidding(w http.ResponseWriter, r *http.Request) {
	var bidding entities.Bidding
	json.NewDecoder(r.Body).Decode(&bidding)

	if !validateBiddingRequest(bidding, w) {
		utils.BadRequest(w, "Invalid Request")
		return
	}

	query := "INSERT INTO biddings (`bidder_id`, `auction_id`, `amount`) VALUES(?, ?, ?);"
	inserResult, err := db.Instance.ExecContext(context.Background(), query, bidding.BidderId, bidding.AuctionId, bidding.Amount)

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt insert Bidding data")
		return
	}
	biddingId, err := inserResult.LastInsertId()

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt retrive Bidding ID")
		return
	} else {
		utils.Ok(w)
		bidding.Id = int(biddingId)
		json.NewEncoder(w).Encode(bidding)
		fmt.Println("Bidding Created")
	}

}

func validateBiddingRequest(bidding entities.Bidding, w http.ResponseWriter) bool {

	var auctions = supplier.GetAuctionsById(bidding.AuctionId)
	if auctions == nil {
		fmt.Println("No auctions found")
		return false
	}
	currentTimeUTC := time.Now().Local()
	difference := auctions[0].EndTime.Local().Sub(currentTimeUTC)
	if difference >= 0 {
		fmt.Println("Bidding in right time")
		return supplier.UpdateAuctionStatus(bidding.AuctionId, "STARTED")
	} else {
		fmt.Println("Auction ended")
		if !supplier.UpdateAuctionStatus(bidding.AuctionId, "COMPLETED") {
			return false
		}
		return false
	}
}

func GetBiddingsByAution(w http.ResponseWriter, r *http.Request) {
	autionId := mux.Vars(r)["auctionId"]
	var biddings []entities.Bidding
	isActive := true
	res, err := db.Instance.Query("SELECT * FROM biddings where is_active = ? and auction_id = ? order by amount desc", isActive, autionId)

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt retrive bidders")
		return
	}
	for res.Next() {
		var bidding entities.Bidding
		err := res.Scan(&bidding.Id, &bidding.BidderId, &bidding.AuctionId, &bidding.IsActive, &bidding.Amount, &bidding.CreatedAt, &bidding.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			utils.InternalError(w, "Couldnt scan bidders")
			return
		}
		biddings = append(biddings, bidding)
	}
	if biddings != nil {
		utils.Ok(w)
		json.NewEncoder(w).Encode(biddings)
	} else {
		utils.NotFound(w, "Couldnt find biddings")
	}
	fmt.Println("Biddings Details are fetched")

}
