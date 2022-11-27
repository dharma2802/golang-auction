package demand

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
)

func InsertBidding(w http.ResponseWriter, r *http.Request) {
	var bidding entities.Bidding
	json.NewDecoder(r.Body).Decode(&bidding)
	query := "INSERT INTO biddings (`bidder_id`, `auction_id`, `amount`) VALUES(?, ?, ?);"
	inserResult, err := db.Instance.ExecContext(context.Background(), query, bidding.BidderId, bidding.AuctionId, bidding.Amount)

	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = "Couldnt insert Bidding data"
		json.NewEncoder(w).Encode(response)
	} else {
		biddingId, err := inserResult.LastInsertId()

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Couldnt retrive Bidding ID"
			json.NewEncoder(w).Encode(response)
		} else {
			bidding.Id = int(biddingId)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bidding)
			fmt.Println("Bidding Created")
		}

	}

}

func GetBiddingsByAution(w http.ResponseWriter, r *http.Request) {
	autionId := mux.Vars(r)["auctionId"]
	var biddings []entities.Bidding
	isActive := true
	res, err := db.Instance.Query("SELECT * FROM biddings where is_active = ? and auction_id = ? order by amount desc", isActive, autionId)

	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = "Couldnt retrive bidders"
		json.NewEncoder(w).Encode(response)
		return
	}
	for res.Next() {
		var bidding entities.Bidding
		err := res.Scan(&bidding.Id, &bidding.BidderId, &bidding.AuctionId, &bidding.IsActive, &bidding.Amount, &bidding.CreatedAt, &bidding.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Couldnt scan bidders"
			json.NewEncoder(w).Encode(response)
			return
		}
		biddings = append(biddings, bidding)
	}
	if biddings != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(biddings)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response.Status = http.StatusNotFound
		response.Message = "Couldnt find biddings"
		json.NewEncoder(w).Encode(response)
	}
	fmt.Println("Biddings Details are fetched")

}
