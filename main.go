package main

import (
	"database/sql"
	"fmt"
	"ftgo-8-phase-1-pair-project/config"
	"ftgo-8-phase-1-pair-project/database"
	"ftgo-8-phase-1-pair-project/entity"
	"ftgo-8-phase-1-pair-project/handler"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	db          *sql.DB
	currentUser *entity.User
)

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
		if currentUser == nil {
			showAuthMenu()
		} else {
			showMainMenu()
		}
		choice := getUserChoice()
		handleChoice(choice)
	}
}

func showAuthMenu() {
	color.Cyan("\n=== Games Store CLI ===")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
}

func getUserChoice() int {
	var choice int
	fmt.Print("\nEnter your choice: ")
	fmt.Scan(&choice)
	return choice
}

func showMainMenu() {
	color.Cyan("\n=== Games Store CLI ===")
	color.Yellow("Logged in as: %s (%s)", currentUser.Email, currentUser.Role)

	if currentUser.Role == "admin" {
		fmt.Println("1. Print Admin Email")
		fmt.Println("2. Logout")
		fmt.Println("3. Exit")
	} else {
		fmt.Println("1. Print Customer Email")
		fmt.Println("2. My Orders")
		fmt.Println("3. Logout")
		fmt.Println("4. Exit")
	}
}

func handleChoice(choice int) {
	if currentUser == nil {
		switch choice {
		case 1:
			handler.HandleUserRegistration(db)
		case 2:
			if user := handler.HandleUserLogin(db); user != nil {
				currentUser = user
			}
		case 3:
			os.Exit(0)
		}
		return
	}

	if currentUser.Role == "admin" {
		handleAdminChoice(choice)
	} else {
		handleCustomerChoice(choice)
	}
}

func handleAdminChoice(choice int) {
	switch choice {
	case 1:
		color.Magenta("Email: %s", currentUser.Email)
	case 2:
		currentUser = nil
	case 3:
		os.Exit(0)
	}
}

func handleCustomerChoice(choice int) {
	switch choice {
	case 1:
		color.Magenta("Email: %s", currentUser.Email)
	case 2:
		handler.HandleCustomerOrderMenu(db, currentUser.ID)
	case 3:
		currentUser = nil
	case 4:
		os.Exit(0)
	}
}

func game() {
	// Buat koneksi database
	db, err := sql.Open("postgres", "your_connection_string_here")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	for {
		fmt.Println("\n=== Game Management ===")
		fmt.Println("1. Add Game")
		fmt.Println("2. List Games")
		fmt.Println("3. Update Game")
		fmt.Println("4. Delete Game")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handler.AddGame(db)
		case 2:
			handler.ListGames(db)
		case 3:
			handler.UpdateGame(db)
		case 4:
			handler.DeleteGame(db)
		case 5:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
