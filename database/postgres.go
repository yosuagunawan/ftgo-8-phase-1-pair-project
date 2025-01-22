// internal/database/postgres.go
package database

import (
	"database/sql"
	"fmt"
	"ftgo-8-phase-1-pair-project/config"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s sslmode=require",
		cfg.Name,
		cfg.User,
		cfg.Password,
		cfg.Host,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}
