package supplier

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
)

func GetAdSpaces(w http.ResponseWriter, r *http.Request) {
	var adSpaces []entities.AdSpace
	isActive := true
	res, err := db.Instance.Query("SELECT * FROM ad_spaces where is_active = ?", isActive)

	if err != nil {
		log.Fatal(err)
		panic("Cannot fetch ad spaces")
	}

	for res.Next() {
		var adSpace entities.AdSpace
		err := res.Scan(&adSpace.Id, &adSpace.Name, &adSpace.Position, &adSpace.Width, &adSpace.Height, &adSpace.Price, &adSpace.IsActive, &adSpace.CreatedAt, &adSpace.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}
		adSpaces = append(adSpaces, adSpace)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(adSpaces)
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
