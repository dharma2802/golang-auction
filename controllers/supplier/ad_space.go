package supplier

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seller-app/auction/db"
	"github.com/seller-app/auction/entities"
	"github.com/seller-app/auction/utils"
)

func GetAdSpaces(w http.ResponseWriter, r *http.Request) {
	var adSpaces []entities.AdSpace
	isActive := true
	res, err := db.Instance.Query("SELECT * FROM ad_spaces where is_active = ?", isActive)

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Cannot fetch ad spaces")
		return
	}

	for res.Next() {
		var adSpace entities.AdSpace
		err := res.Scan(&adSpace.Id, &adSpace.Name, &adSpace.Position, &adSpace.Width, &adSpace.Height, &adSpace.Price, &adSpace.IsActive, &adSpace.CreatedAt, &adSpace.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			utils.InternalError(w, "Couldnt scan ad space")
			return
		}
		adSpaces = append(adSpaces, adSpace)
	}
	if adSpaces != nil {
		utils.Ok(w)
		json.NewEncoder(w).Encode(adSpaces)
	} else {
		utils.NotFound(w, "Couldnt find ad spaces")
	}
	fmt.Println("Ad Space Details are fetched")
}

func InsertAdSpaces(w http.ResponseWriter, r *http.Request) {
	var adSpace entities.AdSpace
	json.NewDecoder(r.Body).Decode(&adSpace)
	query := "INSERT INTO ad_spaces (`name`, `position`, `width`, `height`, `price`) VALUES(?, ?, ?, ?, ?);"
	inserResult, err := db.Instance.ExecContext(context.Background(), query, adSpace.Name, adSpace.Position, adSpace.Width, adSpace.Height, adSpace.Price)

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt insert Ad Space data")
		return
	}
	adSpaceId, err := inserResult.LastInsertId()

	if err != nil {
		fmt.Println(err)
		utils.InternalError(w, "Couldnt retrive Ad Space ID")
	} else {
		adSpace.Id = int(adSpaceId)

		utils.Ok(w)
		json.NewEncoder(w).Encode(adSpace)
		fmt.Println("Ad Space Created")
	}

}
