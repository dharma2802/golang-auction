package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/sa/v1/ad-spaces", supplier.GetAdSpaces).Methods("GET")
	router.HandleFunc("/sa/v1/ad-space", supplier.InsertAdSpaces).Methods("POST")
}
