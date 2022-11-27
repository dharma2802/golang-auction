package supplier

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
)

func GetAuctions(w http.ResponseWriter, r *http.Request) {
	var auctions []entities.Auction
	isActive := true
	res, err := db.Instance.Query("SELECT * FROM auctions where is_active = ?", isActive)

	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = "Cannot fetch auctions"
		json.NewEncoder(w).Encode(response)
		return
	}

	for res.Next() {
		var auction entities.Auction
		err := res.Scan(&auction.Id, &auction.AdSpaceId, &auction.EndTime, &auction.Status, &auction.IsActive, &auction.CreatedAt, &auction.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Couldnt scan auctions"
			json.NewEncoder(w).Encode(response)
			return
		}
		auctions = append(auctions, auction)
	}
	if auctions != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(auctions)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response.Status = http.StatusNotFound
		response.Message = "Couldnt find auctions"
		json.NewEncoder(w).Encode(response)
	}
	fmt.Println("Auctions Details are fetched")
}

func InsertAuction(w http.ResponseWriter, r *http.Request) {
	var auction entities.Auction
	json.NewDecoder(r.Body).Decode(&auction)

	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse

	currentTimeUTC := time.Now().Local()
	difference := auction.EndTime.Sub(currentTimeUTC)

	if difference <= 0 {
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = "End Time should not be less than current time"
		json.NewEncoder(w).Encode(response)
	} else {
		auction.Status = "PENDING"
		query := "INSERT INTO auctions (`ad_space_id`, `end_time`, `status`) VALUES(?, ?, ?);"
		inserResult, err := db.Instance.ExecContext(context.Background(), query, auction.AdSpaceId, auction.EndTime, auction.Status)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Couldnt insert Auction data"
			json.NewEncoder(w).Encode(response)
		} else {
			auctionId, err := inserResult.LastInsertId()

			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				response.Status = http.StatusInternalServerError
				response.Message = "Couldnt retrive Auction ID"
				json.NewEncoder(w).Encode(response)
			} else {
				auction.Id = int(auctionId)

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(auction)
				fmt.Println("Auction Created")
			}

		}
	}

}
