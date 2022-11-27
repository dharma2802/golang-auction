package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seller-app/auction/controllers/demand"
	"github.com/seller-app/auction/controllers/supplier"
	"github.com/seller-app/auction/db"
)

func main() {

	db.Connect()

	router := mux.NewRouter().StrictSlash(true)
	RegisterProductRoutes(router)
	log.Println(fmt.Sprintf("Starting Server on port %s", 7010))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 7010), router))
}

func RegisterProductRoutes(router *mux.Router) {
	// ad spaces routes
	router.HandleFunc("/sa/v1/ad-space", supplier.InsertAdSpaces).Methods("POST")
	router.HandleFunc("/sa/v1/ad-spaces", supplier.GetAdSpaces).Methods("GET")

	// auctions routes
	router.HandleFunc("/sa/v1/auction", supplier.InsertAuction).Methods("POST")
	router.HandleFunc("/sa/v1/auctions", supplier.GetAuctions).Methods("GET")

	// bidder routes
	router.HandleFunc("/sa/v1/bidder", demand.InsertBidder).Methods("POST")
	router.HandleFunc("/sa/v1/bidders", demand.GetBidders).Methods("GET")

	// bidding routes
	router.HandleFunc("/sa/v1/bidding", demand.InsertBidding).Methods("POST")
	router.HandleFunc("/sa/v1/bidding/{auctionId}", demand.GetBiddingsByAution).Methods("GET")

}
