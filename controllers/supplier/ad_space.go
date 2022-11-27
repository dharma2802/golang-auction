package supplier

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
)

func GetAdSpaces(w http.ResponseWriter, r *http.Request) {
	var adSpaces []entities.AdSpace
	isActive := true
	res, err := db.Instance.Query("SELECT * FROM ad_spaces where is_active = ?", isActive)

	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = "Cannot fetch ad spaces"
		json.NewEncoder(w).Encode(response)
		return
	}

	for res.Next() {
		var adSpace entities.AdSpace
		err := res.Scan(&adSpace.Id, &adSpace.Name, &adSpace.Position, &adSpace.Width, &adSpace.Height, &adSpace.Price, &adSpace.IsActive, &adSpace.CreatedAt, &adSpace.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			response.Message = "Couldnt scan ad space"
			json.NewEncoder(w).Encode(response)
			return
		}
		adSpaces = append(adSpaces, adSpace)
	}
	if adSpaces != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(adSpaces)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response.Status = http.StatusNotFound
		response.Message = "Couldnt find ad spaces"
		json.NewEncoder(w).Encode(response)
	}
	fmt.Println("Ad Space Details are fetched")
}

func InsertAdSpaces(w http.ResponseWriter, r *http.Request) {
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
