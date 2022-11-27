package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Instance *sql.DB

func Connect() {

	database, dbError := sql.Open("mysql", "root:rootroot@tcp(db:3306)/auctions?parseTime=true")

	// database, dbError := gorm.Open("postgres", "dbname=auctions sslmode=disable")
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	Instance = database
	log.Println("Connected to Database")
}
