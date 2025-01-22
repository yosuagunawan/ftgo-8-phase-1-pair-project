package handler

import (
	"database/sql"
	"fmt"
	"ftgo-8-phase-1-pair-project/entity"

	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
)

func HandleUserMenu(db *sql.DB) {
	for {
		color.Cyan("\n=== User Management ===")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Back")

		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handleUserRegistration(db)
		case 2:
			handleUserLogin(db)
		case 3:
			return
		default:
			color.Red("Invalid choice!")
		}
	}
}

func handleUserRegistration(db *sql.DB) {
	var email, password string
	fmt.Print("Enter email: ")
	fmt.Scan(&email)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		color.Red("Error hashing password:", err)
		return
	}

	_, err = db.Exec(
		"INSERT INTO users (email, password, role) VALUES ($1, $2, 'customer')",
		email, hashedPassword)
	if err != nil {
		color.Red("Error registering user:", err)
		return
	}

	color.Green("Registration successful! Please login.")
}

func handleUserLogin(db *sql.DB) *entity.User {
	var email, password string
	fmt.Print("Enter email: ")
	fmt.Scan(&email)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	var user entity.User
	var storedPassword string
	err := db.QueryRow(
		"SELECT id, email, password, role FROM users WHERE email = $1",
		email).Scan(&user.ID, &user.Email, &storedPassword, &user.Role)
	if err != nil {
		color.Red("Invalid credentials")
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		color.Red("Invalid credentials")
		return nil
	}

	color.Green("Login successful!")
	return &user
}
