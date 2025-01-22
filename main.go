package main

import (
	"database/sql"
	"fmt"
	"ftgo-8-phase-1-pair-project/config"
	"ftgo-8-phase-1-pair-project/database"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err = database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM games")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var price float64
		var stock int
		var releaseDate string
		var created_at string
		var categoryID int
		err = rows.Scan(&id, &title, &price, &stock, &releaseDate, &created_at, &categoryID)
		if err != nil {
			panic(err)
		}

		fmt.Println(title, price)
	}
}
