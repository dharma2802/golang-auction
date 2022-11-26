package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var Instance *sql.DB
var dbError error

type AdSpace struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Position  string  `json:"position"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	Price     float64 `json:"price"`
	IsActive  bool    `json:"isActive"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func GetAdSpaces(w http.ResponseWriter, r *http.Request) {
	var adSpaces []AdSpace
	isActive := true
	res, err := Instance.Query("SELECT * FROM ad_spaces where is_active = ?", isActive)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
		panic("Cannot fetch ad spaces")
	}

	for res.Next() {
		var adSpace AdSpace
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
	var adSpace AdSpace
	json.NewDecoder(r.Body).Decode(&adSpace)
	query := "INSERT INTO ad_spaces (`name`, `position`, `width`, `height`, `price`) VALUES(?, ?, ?, ?, ?);"
	inserResult, err := Instance.ExecContext(context.Background(), query, adSpace.Name, adSpace.Position, adSpace.Width, adSpace.Height, adSpace.Price)

	w.Header().Set("Content-Type", "application/json")
	var response DefaultResponse
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

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/sa/v1/ad-spaces", GetAdSpaces).Methods("GET")
	router.HandleFunc("/sa/v1/ad-space", InsertAdSpaces).Methods("POST")
}

func migrate() {
	// Instance.AutoMigrate(&AdSpace{})
	// Instance.AutoMigrate(&Auction)
	// Instance.AutoMigrate(&Bidder)
	// Instance.AutoMigrate(&Bidding)

	log.Println("Database migration completed")
}

func connect() {
	database, dbError := sql.Open("mysql", "root:rootroot@tcp(0.0.0.0:3306)/auctions")
	// defer database.Close()

	// database, dbError := gorm.Open("postgres", "dbname=auctions sslmode=disable")
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	Instance = database
	log.Println("Connected to Database")

	migrate()
}

func dbInit() {
	connect()
}

func main() {
	fmt.Println("Welcome to Auction System")
	dbInit()
	router := mux.NewRouter().StrictSlash(true)
	RegisterProductRoutes(router)
	log.Println(fmt.Sprintf("Starting Server on port %s", 4000))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 4000), router))
}
