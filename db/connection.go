package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Instance *sql.DB

func Connect() {

	database, dbError := sql.Open("mysql", "root:rootroot@tcp(0.0.0.0:3306)/auctions")

	// database, dbError := gorm.Open("postgres", "dbname=auctions sslmode=disable")
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	Instance = database
	log.Println("Connected to Database")

	// migrate()
}

func Migrate() {

	// Instance.AutoMigrate(&entities.AdSpace)
	// Instance.AutoMigrate(&entities.Auction)
	// Instance.AutoMigrate(&entities.Bidder)
	// Instance.AutoMigrate(&entities.Bidding)

	log.Println("Database migration completed")
}
