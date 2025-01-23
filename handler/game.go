package handler

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Game represents the structure of a game
type Game struct {
	ID          int
	Title       string
	Price       float64
	Stock       int
	Genre       string
	CategoryID  int
	CreatedAt   time.Time
	ReleaseDate string
}

// AddGame manually adds a new game to the database
func AddGame(db *sql.DB) {
	var title, genre, releaseDate string
	var price float64
	var stock, categoryID int

	fmt.Print("Enter game title: ")
	fmt.Scan(&title)
	fmt.Print("Enter game price: ")
	fmt.Scan(&price)
	fmt.Print("Enter game stock: ")
	fmt.Scan(&stock)
	fmt.Print("Enter game genre: ")
	fmt.Scan(&genre)
	fmt.Print("Enter category ID: ")
	fmt.Scan(&categoryID)
	fmt.Print("Enter release date (YYYY-MM-DD): ")
	fmt.Scan(&releaseDate)

	createdAt := time.Now()

	query := `
		INSERT INTO games (title, price, stock, genre, category_id, release_date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Exec(query, title, price, stock, genre, categoryID, releaseDate, createdAt)
	if err != nil {
		color.Red("Error adding game:", err)
		return
	}

	color.Green("Game added successfully!")
}

// ListGames manually retrieves all games and displays them
func ListGames(db *sql.DB) {
	query := `
		SELECT id, title, price, stock, genre, category_id, release_date, created_at 
		FROM games ORDER BY id`

	rows, err := db.Query(query)
	if err != nil {
		color.Red("Error fetching games:", err)
		return
	}
	defer rows.Close()

	color.Cyan("\n=== Available Games ===")
	fmt.Printf("%-5s %-30s %-10s %-10s %-20s %-15s %-20s %-20s\n",
		"ID", "Title", "Price", "Stock", "Genre", "Category ID", "Release Date", "Created At")

	for rows.Next() {
		var game Game
		err := rows.Scan(&game.ID, &game.Title, &game.Price, &game.Stock, &game.Genre, &game.CategoryID, &game.ReleaseDate, &game.CreatedAt)
		if err != nil {
			color.Red("Error reading game data:", err)
			continue
		}
		fmt.Printf("%-5d %-30s $%-9.2f %-10d %-20s %-15d %-20s %-20s\n",
			game.ID, game.Title, game.Price, game.Stock, game.Genre, game.CategoryID, game.ReleaseDate, game.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

// UpdateGame manually updates a game's details
func UpdateGame(db *sql.DB) {
	var id int
	var title, genre, releaseDate string
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
	fmt.Print("Enter new genre: ")
	fmt.Scan(&genre)
	fmt.Print("Enter new category ID: ")
	fmt.Scan(&categoryID)
	fmt.Print("Enter new release date (YYYY-MM-DD): ")
	fmt.Scan(&releaseDate)

	query := `
		UPDATE games 
		SET title = $1, price = $2, stock = $3, genre = $4, category_id = $5, release_date = $6 
		WHERE id = $7`

	_, err := db.Exec(query, title, price, stock, genre, categoryID, releaseDate, id)
	if err != nil {
		color.Red("Error updating game:", err)
		return
	}

	color.Green("Game updated successfully!")
}

// DeleteGame manually deletes a game
func DeleteGame(db *sql.DB) {
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
