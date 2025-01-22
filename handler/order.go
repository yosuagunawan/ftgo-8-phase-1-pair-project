package handler

import (
	"database/sql"
	"fmt"

	"github.com/fatih/color"
)

func HandleOrderMenu(db *sql.DB) {
	for {
		color.Cyan("\n=== Order Management ===")
		fmt.Println("1. View Orders")
		fmt.Println("2. Place Order")
		fmt.Println("3. Back")

		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handleViewOrders(db)
		case 2:
			handlePlaceOrder(db)
		case 3:
			return
		default:
			color.Red("Invalid choice!")
		}
	}
}

func handleViewOrders(db *sql.DB) {
	rows, err := db.Query(`
        SELECT o.id, g.title, o.quantity, o.total, o.created_at
        FROM orders o
        JOIN games g ON o.game_id = g.id
        ORDER BY o.created_at DESC`)
	if err != nil {
		color.Red("Error fetching orders:", err)
		return
	}
	defer rows.Close()

	color.Yellow("\nOrder History:")
	fmt.Printf("%-5s %-30s %-10s %-10s %-20s\n",
		"ID", "Game", "Quantity", "Total", "Date")

	for rows.Next() {
		var id, quantity int
		var title, createdAt string
		var total float64

		rows.Scan(&id, &title, &quantity, &total, &createdAt)
		fmt.Printf("%-5d %-30s %-10d $%-9.2f %-20s\n",
			id, title, quantity, total, createdAt)
	}
}

func handlePlaceOrder(db *sql.DB) {
	// handleListGames(db)

	var gameID, quantity int
	fmt.Print("\nEnter game ID: ")
	fmt.Scan(&gameID)
	fmt.Print("Enter quantity: ")
	fmt.Scan(&quantity)

	tx, err := db.Begin()
	if err != nil {
		color.Red("Error starting transaction:", err)
		return
	}

	var stock int
	var price float64
	err = tx.QueryRow("SELECT stock, price FROM games WHERE id =$1", gameID).
		Scan(&stock, &price)
	if err != nil {
		tx.Rollback()
		color.Red("Error getting game details:", err)
		return
	}

	if stock < quantity {
		tx.Rollback()
		color.Red("Insufficient stock! Available: %d", stock)
		return
	}

	total := price * float64(quantity)
	_, err = tx.Exec(`
        INSERT INTO orders (game_id, quantity, total) 
        VALUES ($1, $2, $3)`,
		gameID, quantity, total)
	if err != nil {
		tx.Rollback()
		color.Red("Error creating order:", err)
		return
	}

	_, err = tx.Exec("UPDATE games SET stock = stock - $1 WHERE id = $2",
		quantity, gameID)
	if err != nil {
		tx.Rollback()
		color.Red("Error updating stock:", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		color.Red("Error completing order:", err)
		return
	}

	color.Green("Order placed successfully!")
}
