package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitialDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Config("DB_HOST"), Config("DB_PORT"), Config("DB_USER"), Config("DB_PASSWORD"), Config("DB_NAME"))

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	if err := db.Ping(); err != nil {
		defer db.Close()

		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}
	return db, nil
}
