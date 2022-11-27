package demand

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
	"github.com/seller-app/auction/utils"
)

func InsertBidder(w http.ResponseWriter, r *http.Request) {
	var bidder entities.Bidder
	json.NewDecoder(r.Body).Decode(&bidder)
	query := "INSERT INTO bidder (`name`, `email`) VALUES(?, ?);"
	inserResult, err := db.Instance.ExecContext(context.Background(), query, bidder.Name, bidder.Email)

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt insert Bidder data")
		return
	}
	bidderId, err := inserResult.LastInsertId()

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt retrive Bidder ID")
		return
	}
	bidder.Id = int(bidderId)
	utils.Ok(w)
	json.NewEncoder(w).Encode(bidder)
	fmt.Println("Ad Space Created")
}

func GetBidders(w http.ResponseWriter, r *http.Request) {
	var bidders []entities.Bidder
	isActive := true
	res, err := db.Instance.Query("SELECT * FROM bidder where is_active = ?", isActive)

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt retrive bidders")
		return
	}
	for res.Next() {
		var bidder entities.Bidder
		err := res.Scan(&bidder.Id, &bidder.Name, &bidder.Email, &bidder.CreatedAt, &bidder.UpdatedAt, &bidder.IsActive)
		fmt.Println(bidder)
		if err != nil {
			fmt.Println(err)
			utils.InternalError(w, "Couldnt scan the bidder")
			return
		}
		bidders = append(bidders, bidder)

	}
	if bidders != nil {
		utils.Ok(w)
		json.NewEncoder(w).Encode(bidders)
	} else {
		utils.NotFound(w, "Couldnt find bidders")
	}
	fmt.Println("Bidders Details are fetched")

}
