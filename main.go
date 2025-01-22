package main

import (
	"database/sql"
	"fmt"
	"ftgo-8-phase-1-pair-project/config"
	"ftgo-8-phase-1-pair-project/database"
	"ftgo-8-phase-1-pair-project/handler"
	"log"
	"os"

	"github.com/fatih/color"
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

	for {
		showMainMenu()
		choice := getUserChoice()
		handleUserChoice(choice)
	}
	// rows, err := db.Query("SELECT * FROM games")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var id int
	// 	var title string
	// 	var price float64
	// 	var stock int
	// 	var releaseDate string
	// 	var created_at string
	// 	var categoryID int
	// 	err = rows.Scan(&id, &title, &price, &stock, &releaseDate, &created_at, &categoryID)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(title, price)
	// }
}

func showMainMenu() {
	color.Cyan("\n=== Games Store CLI ===")
	fmt.Println("1. User Management")
	fmt.Println("2. Game Management")
	fmt.Println("3. Order Management")
	fmt.Println("4. Reports")
	fmt.Println("5. Exit")
}

func getUserChoice() int {
	var choice int
	fmt.Print("\nEnter your choice: ")
	fmt.Scan(&choice)
	return choice
}

func handleUserChoice(choice int) {
	switch choice {
	case 1:
		fmt.Println("User")
	case 2:
		fmt.Println("Games")
	case 3:
		handler.HandleOrderMenu(db)
	case 4:
		fmt.Println("Reports")
	case 5:
		os.Exit(0)
	default:
		color.Red("Invalid choice!")
	}
}
