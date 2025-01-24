package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReportHandler struct untuk menyimpan koneksi database
type ReportHandler struct {
	DB sql.DB
}

func (hReportHandler) ShowReportsMenu() {
	reader := bufio.NewReader(os.Stdin) // Reader hanya dibuat sekali di awal
	for {
		// Cetak menu
		fmt.Println("Reports Menu:")
		fmt.Println("1. Low Stock Alert Report")
		fmt.Println("2. Customer Purchase Frequency")
		fmt.Println("3. Sales Performance by Category")
		fmt.Println("4. Recent Game Releases Performance")
		fmt.Println("5. Average Order Value by Month")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		// Baca input pengguna
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		// Bersihkan input dari spasi dan newline
		input = strings.TrimSpace(input)

		// Tangani input kosong
		if input == "" {
			fmt.Println("Input cannot be empty. Please enter a number between 1 and 6.")
			continue
		}

		// Konversi input ke integer
		_, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 6.")
			continue
		}
	}
}

// LowStockAlertReportCLI menangani laporan stok rendah untuk CLI
func (h *ReportHandler) LowStockAlertReportCLI() {
	query := `
		SELECT title, stock, price
		FROM games
		WHERE stock < 10
		ORDER BY stock ASC;
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return
	}
	defer rows.Close()

	fmt.Println("Low Stock Alert Report")
	fmt.Println("========================")
	fmt.Printf("%-30s %-10s %-10s\n", "Title", "Stock", "Price")
	fmt.Println("--------------------------------------------")

	for rows.Next() {
		var title string
		var stock int
		var price float64
		if err := rows.Scan(&title, &stock, &price); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return
		}
		fmt.Printf("%-30s %-10d %-10.2f\n", title, stock, price)
	}
}

// CustomerPurchaseFrequencyCLI menangani analisis frekuensi pembelian pelanggan untuk CLI
func (h *ReportHandler) CustomerPurchaseFrequencyCLI() {
	query := `
		SELECT u.email, COUNT(o.id) AS order_count, SUM(o.total) AS total_spent
		FROM users u
		JOIN orders o ON u.id = o.user_id
		GROUP BY u.email
		ORDER BY order_count DESC;
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return
	}
	defer rows.Close()

	fmt.Println("Customer Purchase Frequency Analysis")
	fmt.Println("=====================================")
	fmt.Printf("%-30s %-15s %-15s\n", "Email", "Order Count", "Total Spent")
	fmt.Println("-------------------------------------------------------------")

	for rows.Next() {
		var email string
		var orderCount int
		var totalSpent float64
		if err := rows.Scan(&email, &orderCount, &totalSpent); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return
		}
		fmt.Printf("%-30s %-15d %-15.2f\n", email, orderCount, totalSpent)
	}
}

// SalesPerformanceByCategoryCLI menangani performa penjualan berdasarkan kategori untuk CLI
func (h *ReportHandler) SalesPerformanceByCategoryCLI() {
	query := `
		SELECT c.name, COUNT(o.id) AS orders, SUM(o.total) AS revenue
		FROM categories c
		JOIN games g ON c.id = g.category_id
		JOIN orders o ON g.id = o.game_id
		GROUP BY c.name
		ORDER BY revenue DESC;
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return
	}
	defer rows.Close()

	fmt.Println("Sales Performance by Category")
	fmt.Println("=============================")
	fmt.Printf("%-20s %-10s %-15s\n", "Category Name", "Orders", "Revenue")
	fmt.Println("----------------------------------------------")

	for rows.Next() {
		var categoryName string
		var orders int
		var revenue float64
		if err := rows.Scan(&categoryName, &orders, &revenue); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return
		}
		fmt.Printf("%-20s %-10d %-15.2f\n", categoryName, orders, revenue)
	}
}

// RecentGameReleasesPerformanceCLI menangani performa game baru untuk CLI
func (h *ReportHandler) RecentGameReleasesPerformanceCLI() {
	query := `
		SELECT g.title, g.release_date, COUNT(o.id) AS orders
		FROM games g
		LEFT JOIN orders o ON g.id = o.game_id
		WHERE g.release_date >= NOW() - INTERVAL '30 days'
		GROUP BY g.title, g.release_date
		ORDER BY orders DESC;
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return
	}
	defer rows.Close()

	fmt.Println("Recent Game Releases Performance")
	fmt.Println("================================")
	fmt.Printf("%-30s %-15s %-10s\n", "Title", "Release Date", "Orders")
	fmt.Println("--------------------------------------------------")

	for rows.Next() {
		var title, releaseDate string
		var orders int
		if err := rows.Scan(&title, &releaseDate, &orders); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return
		}
		fmt.Printf("%-30s %-15s %-10d\n", title, releaseDate, orders)
	}
}

// AverageOrderValueByMonthCLI menangani laporan nilai rata-rata order bulanan untuk CLI
func (h *ReportHandler) AverageOrderValueByMonthCLI() {
	query := `
		SELECT 
			DATE_TRUNC('month', o.created_at) AS month,
			ROUND(AVG(o.total), 2) AS avg_order_value,
			COUNT(o.id) AS total_orders
		FROM orders o
		GROUP BY DATE_TRUNC('month', o.created_at)
		ORDER BY month DESC;
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return
	}
	defer rows.Close()

	fmt.Println("Average Order Value by Month")
	fmt.Println("============================")
	fmt.Printf("%-15s %-20s %-15s\n", "Month", "Avg Order Value", "Total Orders")
	fmt.Println("-----------------------------------------------------")

	for rows.Next() {
		var month string
		var avgOrderValue float64
		var totalOrders int
		if err := rows.Scan(&month, &avgOrderValue, &totalOrders); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return
		}
		fmt.Printf("%-15s %-20.2f %-15d\n", month, avgOrderValue, totalOrders)
	}
}
