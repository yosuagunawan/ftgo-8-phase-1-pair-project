package handler

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Game struct {
	ID          int
	Title       string
	Price       float64
	Stock       int
	CategoryID  int
	CreatedAt   time.Time
	ReleaseDate string
}

func HandleGameMenu(db *sql.DB) {
	for {
		color.Cyan("\n=== Game Management ===")
		fmt.Println("1. Add Game")
		fmt.Println("2. List Games")
		fmt.Println("3. Update Game")
		fmt.Println("4. Delete Game")
		fmt.Println("5. Back")

		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handleAddGame(db)
		case 2:
			HandleListGames(db)
		case 3:
			handleUpdateGame(db)
		case 4:
			handleDeleteGame(db)
		case 5:
			return
		default:
			color.Red("Invalid choice!")
		}
	}
}

func handleAddGame(db *sql.DB) {
	var title, releaseDate string
	var price float64
	var stock, categoryID int

	fmt.Print("Enter game title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter game price: ")
	fmt.Scan(&price)
	fmt.Print("Enter game stock: ")
	fmt.Scan(&stock)
	fmt.Print("Enter category ID: ")
	fmt.Scan(&categoryID)
	fmt.Print("Enter release date (YYYY-MM-DD): ")
	fmt.Scan(&releaseDate)
	createdAt := time.Now()

	query := `
		INSERT INTO games (title, price, stock, category_id, release_date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(query, title, price, stock, categoryID, releaseDate, createdAt)
	if err != nil {
		color.Red("Error adding game:", err)
		return
	}

	color.Green("Game added successfully!")
}

func HandleListGames(db *sql.DB) {
	query := `
		SELECT id, title, price, stock, category_id, release_date 
		FROM games ORDER BY id`

	rows, err := db.Query(query)
	if err != nil {
		color.Red("Error fetching games:", err)
		return
	}
	defer rows.Close()

	color.Cyan("\n=== Available Games ===")
	fmt.Printf("%-5s %-30s %-10s %-10s %-15s %-20s\n",
		"ID", "Title", "Price", "Stock", "Category ID", "Release Date")

	for rows.Next() {
		var game Game
		err := rows.Scan(&game.ID, &game.Title, &game.Price, &game.Stock, &game.CategoryID, &game.ReleaseDate)
		if err != nil {
			color.Red("Error reading game data:", err)
			continue
		}
		fmt.Printf("%-5d %-30s $%-9.2f %-10d %-15d %-20s\n",
			game.ID, game.Title, game.Price, game.Stock, game.CategoryID, game.ReleaseDate)
	}
}

func handleUpdateGame(db *sql.DB) {
	var id int
	var title, releaseDate string
	var price float64
	var stock, categoryID int

	fmt.Print("Enter game ID to update: ")
	fmt.Scan(&id)

	fmt.Print("Enter new title: ")
	fmt.Scan(&title)
	fmt.Print("Enter new price: ")
	fmt.Scan(&price)
	fmt.Print("Enter new stock: ")
	fmt.Scan(&stock)
	fmt.Print("Enter new category ID: ")
	fmt.Scan(&categoryID)
	fmt.Print("Enter new release date (YYYY-MM-DD): ")
	fmt.Scan(&releaseDate)

	query := `
		UPDATE games 
		SET title = $1, price = $2, stock = $3, category_id = $4, release_date = $5 
		WHERE id = $6`

	_, err := db.Exec(query, title, price, stock, categoryID, releaseDate, id)
	if err != nil {
		color.Red("Error updating game:", err)
		return
	}

	color.Green("Game updated successfully!")
}

func handleDeleteGame(db *sql.DB) {
	var id int
	fmt.Print("Enter game ID to delete: ")
	fmt.Scan(&id)

	query := `
		DELETE FROM games WHERE id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		color.Red("Error deleting game:", err)
		return
	}

	color.Green("Game deleted successfully!")
}
