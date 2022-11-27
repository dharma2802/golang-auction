package supplier

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
	"github.com/seller-app/auction/utils"
)

func GetAuctions(w http.ResponseWriter, r *http.Request) {
	var auctions []entities.Auction
	isActive := true
	res, err := db.Instance.Query("SELECT * FROM auctions where is_active = ?", isActive)

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Cannot fetch auctions")
		return
	}

	for res.Next() {
		var auction entities.Auction
		err := res.Scan(&auction.Id, &auction.AdSpaceId, &auction.EndTime, &auction.Status, &auction.IsActive, &auction.CreatedAt, &auction.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			utils.InternalError(w, "Couldnt scan auctions")
			return
		}
		auctions = append(auctions, auction)
	}
	if auctions != nil {
		utils.Ok(w)
		json.NewEncoder(w).Encode(auctions)
	} else {
		utils.NotFound(w, "Couldnt find auctions")
	}
	fmt.Println("Auctions Details are fetched")
}

func GetAuctionsById(id int) []entities.Auction {
	var auctions []entities.Auction
	isActive := true

	res, err := db.Instance.Query("SELECT * FROM auctions where is_active = ? and id = ?", isActive, id)

	if err != nil {
		fmt.Println(err)
		return auctions
	}

	for res.Next() {
		var auction entities.Auction
		err := res.Scan(&auction.Id, &auction.AdSpaceId, &auction.EndTime, &auction.Status, &auction.IsActive, &auction.CreatedAt, &auction.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return auctions
		}
		auctions = append(auctions, auction)
	}

	return auctions
}

func InsertAuction(w http.ResponseWriter, r *http.Request) {
	var auction entities.Auction
	json.NewDecoder(r.Body).Decode(&auction)

	currentTimeUTC := time.Now().Local()
	auction.EndTime = auction.EndTime.Local()
	difference := auction.EndTime.Sub(currentTimeUTC)

	if difference <= 0 {
		utils.BadRequest(w, "End Time should not be less than current time")
		return
	}
	auction.Status = "PENDING"
	query := "INSERT INTO auctions (`ad_space_id`, `end_time`, `status`) VALUES(?, ?, ?);"
	inserResult, err := db.Instance.ExecContext(context.Background(), query, auction.AdSpaceId, auction.EndTime, auction.Status)

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt insert Auction data")
		return
	}
	auctionId, err := inserResult.LastInsertId()

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt retrive Auction ID")
		return
	}
	auction.Id = int(auctionId)

	utils.Ok(w)
	json.NewEncoder(w).Encode(auction)
	fmt.Println("Auction Created")

}

func UpdateAuctionStatus(id int, status string) bool {
	_, err := db.Instance.Exec("update auctions set status = ? where id = ?", status, id)

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
