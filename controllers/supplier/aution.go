package supplier

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
)

func InsertAuction(w http.ResponseWriter, r *http.Request) {
	var adSpace entities.AdSpace
	json.NewDecoder(r.Body).Decode(&adSpace)
	query := "INSERT INTO ad_spaces (`name`, `position`, `width`, `height`, `price`) VALUES(?, ?, ?, ?, ?);"
	inserResult, err := db.Instance.ExecContext(context.Background(), query, adSpace.Name, adSpace.Position, adSpace.Width, adSpace.Height, adSpace.Price)

	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = "Couldnt insert Ad Space data"
		json.NewEncoder(w).Encode(response)
	} else {
		adSpaceId, err := inserResult.LastInsertId()

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Couldnt retrive Ad Space ID"
			json.NewEncoder(w).Encode(response)
		} else {
			adSpace.Id = int(adSpaceId)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(adSpace)
			fmt.Println("Ad Space Created")
		}

	}

}
