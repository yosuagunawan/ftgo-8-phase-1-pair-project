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
	db            *sql.DB
	currentUser   *entity.User
	reportHandler *handler.ReportHandler // Tambahkan reportHandler
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

	// Inisialisasi ReportHandler
	reportHandler = &handler.ReportHandler{DB: db}

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
		fmt.Println("1. Game Management")
		fmt.Println("2. Reports") // Tambahkan opsi Reports
		fmt.Println("3. Logout")
		fmt.Println("4. Exit")
	} else {
		fmt.Println("1. Browse Games")
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
		handler.HandleGameMenu(db)
	case 2:
		showReportsMenu() // Tambahkan opsi untuk Reports
	case 3:
		currentUser = nil
	case 4:
		os.Exit(0)
	}
}

func handleCustomerChoice(choice int) {
	switch choice {
	case 1:
		handler.HandleListGames(db)
	case 2:
		handler.HandleCustomerOrderMenu(db, currentUser.ID)
	case 3:
		currentUser = nil
	case 4:
		os.Exit(0)
	}
}

func showReportsMenu() {
	color.Cyan("\n=== Reports Menu ===")
	fmt.Println("1. Low Stock Alert Report")
	fmt.Println("2. Customer Purchase Frequency")
	fmt.Println("3. Sales Performance by Category")
	fmt.Println("4. Recent Game Releases Performance")
	fmt.Println("5. Average Order Value by Month")
	fmt.Println("6. Back to Main Menu")

	choice := getUserChoice()
	switch choice {
	case 1:
		reportHandler.LowStockAlertReportCLI()
	case 2:
		reportHandler.CustomerPurchaseFrequencyCLI()
	case 3:
		reportHandler.SalesPerformanceByCategoryCLI()
	case 4:
		reportHandler.RecentGameReleasesPerformanceCLI()
	case 5:
		reportHandler.AverageOrderValueByMonthCLI()
	case 6:
		return
	default:
		fmt.Println("Invalid choice, please try again.")
	}
}
