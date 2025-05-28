package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	//Verify connection
	if err := DB.Ping(); err != nil {
		return err
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			created_at TIMESTAMP NOT NULL
		);
	`)

	if err != nil {
		return fmt.Errorf("failed ot create users table: %w", err)
	}


	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS wallets (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			balance NUMERIC NOT NULL,
			currency TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL
		);
	`)

	if err != nil {
		return fmt.Errorf("failed to create wallets table: %w", err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS transactions (
		id SERIAL PRIMARY KEY,
		wallet_id INT NOT NULL REFERENCES wallets(id),
		amount NUMERIC NOT NULL,
		type TEXT CHECK (type IN ('credit', 'debit')) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		return fmt.Errorf("failed to create transactions table: %w", err)
	}

	return nil
}