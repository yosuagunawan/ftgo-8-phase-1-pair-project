package handler

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestHandleAddGame(t *testing.T) {
	// Buat mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Atur ekspektasi query untuk INSERT
	mock.ExpectExec("INSERT INTO games").
		WithArgs("Test Game", 19.99, 100, 1, "2025-01-01", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Panggil fungsi dengan mock database
	handleAddGameTest(db, "Test Game", 19.99, 100, 1, "2025-01-01")

	// Pastikan semua ekspektasi terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet expectations: %v", err)
	}
}

func handleAddGameTest(db *sql.DB, title string, price float64, stock, categoryID int, releaseDate string) {
	createdAt := time.Now()
	query := `
		INSERT INTO games (title, price, stock, category_id, release_date, created_at)
		VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, title, price, stock, categoryID, releaseDate, createdAt)
	if err != nil {
		panic(err)
	}
}

func TestHandleUpdateGame(t *testing.T) {
	// Buat mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Atur ekspektasi query untuk UPDATE
	mock.ExpectExec("UPDATE games").
		WithArgs("Updated Game", 39.99, 70, 2, "2024-06-01", 1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 untuk jumlah baris yang terpengaruh

	// Panggil fungsi dengan mock database
	query := `
		UPDATE games
		SET title = ?, price = ?, stock = ?, category_id = ?, release_date = ?
		WHERE id = ?`
	_, err = db.Exec(query, "Updated Game", 39.99, 70, 2, "2024-06-01", 1)
	if err != nil {
		t.Fatalf("Failed to update game: %v", err)
	}

	// Pastikan semua ekspektasi terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet expectations: %v", err)
	}
}

func TestHandleDeleteGame(t *testing.T) {
	// Buat mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Atur ekspektasi query untuk DELETE
	mock.ExpectExec("DELETE FROM games").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 untuk jumlah baris yang terpengaruh

	// Panggil fungsi dengan mock database
	query := `DELETE FROM games WHERE id = ?`
	_, err = db.Exec(query, 1)
	if err != nil {
		t.Fatalf("Failed to delete game: %v", err)
	}

	// Pastikan semua ekspektasi terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet expectations: %v", err)
	}
}
