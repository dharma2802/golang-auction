package demand

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
)

func InsertBidder(w http.ResponseWriter, r *http.Request) {
	var bidder entities.Bidder
	json.NewDecoder(r.Body).Decode(&bidder)
	query := "INSERT INTO bidder (`name`, `email`) VALUES(?, ?);"
	inserResult, err := db.Instance.ExecContext(context.Background(), query, bidder.Name, bidder.Email)

	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = "Couldnt insert Bidder data"
		json.NewEncoder(w).Encode(response)
	} else {
		bidderId, err := inserResult.LastInsertId()

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Couldnt retrive Bidder ID"
			json.NewEncoder(w).Encode(response)
		} else {
			bidder.Id = int(bidderId)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bidder)
			fmt.Println("Ad Space Created")
		}

	}

}

func GetBidders(w http.ResponseWriter, r *http.Request) {
	var bidders []entities.Bidder
	isActive := true
	res, err := db.Instance.Query("SELECT * FROM bidder where is_active = ?", isActive)

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
		var bidder entities.Bidder
		err := res.Scan(&bidder.Id, &bidder.Name, &bidder.Email, &bidder.CreatedAt, &bidder.UpdatedAt, &bidder.IsActive)
		fmt.Println(bidder)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Couldnt scan the bidder"
			json.NewEncoder(w).Encode(response)
			return
		}
		bidders = append(bidders, bidder)

	}
	if bidders != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bidders)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response.Status = http.StatusNotFound
		response.Message = "Couldnt find bidders"
		json.NewEncoder(w).Encode(response)
	}

	fmt.Println("Bidders Details are fetched")

}
