package handler

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestLowStockAlertReportCLI(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	query := `
		SELECT title, stock, price
		FROM games
		WHERE stock < 10
		ORDER BY stock ASC;
	`

	// Mocking expected rows
	rows := sqlmock.NewRows([]string{"title", "stock", "price"}).
		AddRow("Game A", 5, 19.99).
		AddRow("Game B", 3, 29.99)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	handler := &ReportHandler{DB: db}
	handler.LowStockAlertReportCLI()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations not met: %v", err)
	}
}

func TestCustomerPurchaseFrequencyCLI(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	query := `
		SELECT u.email, COUNT(o.id) AS order_count, SUM(o.total) AS total_spent
		FROM users u
		JOIN orders o ON u.id = o.user_id
		GROUP BY u.email
		ORDER BY order_count DESC;
	`

	// Mocking expected rows
	rows := sqlmock.NewRows([]string{"email", "order_count", "total_spent"}).
		AddRow("user1@example.com", 5, 100.00).
		AddRow("user2@example.com", 3, 60.00)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	handler := &ReportHandler{DB: db}
	handler.CustomerPurchaseFrequencyCLI()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations not met: %v", err)
	}
}

func TestSalesPerformanceByCategoryCLI(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	query := `
		SELECT c.name, COUNT(o.id) AS orders, SUM(o.total) AS revenue
		FROM categories c
		JOIN games g ON c.id = g.category_id
		JOIN orders o ON g.id = o.game_id
		GROUP BY c.name
		ORDER BY revenue DESC;
	`

	// Mocking expected rows
	rows := sqlmock.NewRows([]string{"name", "orders", "revenue"}).
		AddRow("Action", 10, 500.00).
		AddRow("Adventure", 8, 300.00)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	handler := &ReportHandler{DB: db}
	handler.SalesPerformanceByCategoryCLI()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations not met: %v", err)
	}
}

func TestRecentGameReleasesPerformanceCLI(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	query := `
		SELECT g.title, g.release_date, COUNT(o.id) AS orders
		FROM games g
		LEFT JOIN orders o ON g.id = o.game_id
		WHERE g.release_date >= NOW() - INTERVAL '30 days'
		GROUP BY g.title, g.release_date
		ORDER BY orders DESC;
	`

	// Mocking expected rows
	rows := sqlmock.NewRows([]string{"title", "release_date", "orders"}).
		AddRow("Game A", "2023-12-01", 20).
		AddRow("Game B", "2023-12-15", 10)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	handler := &ReportHandler{DB: db}
	handler.RecentGameReleasesPerformanceCLI()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations not met: %v", err)
	}
}

func TestAverageOrderValueByMonthCLI(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	query := `
		SELECT 
			DATE_TRUNC('month', o.created_at) AS month,
			ROUND(AVG(o.total), 2) AS avg_order_value,
			COUNT(o.id) AS total_orders
		FROM orders o
		GROUP BY DATE_TRUNC('month', o.created_at)
		ORDER BY month DESC;
	`

	// Mocking expected rows
	rows := sqlmock.NewRows([]string{"month", "avg_order_value", "total_orders"}).
		AddRow("2023-12-01", 50.00, 100).
		AddRow("2023-11-01", 45.00, 90)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	handler := &ReportHandler{DB: db}
	handler.AverageOrderValueByMonthCLI()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations not met: %v", err)
	}
}
