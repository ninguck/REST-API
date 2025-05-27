package models

import (
	"database/sql"
	"time"
)

//Wallet represents a users financial wallet
type Wallet struct {
	ID 			int64 		`json:"id"`
	UserID		int64 		`json:"user_id"`
	Balance 	float64 	`json:"balance"`
	Currency 	string 		`json:"currency"`
	CreatedAt 	time.Time 	`json:"create_at`
}

//CreateWallet inserts a new wallet into the database
func CreateWallet(db *sql.DB, wallet Wallet) (Wallet, error) {
	query := `
		INSERT INTO wallets (user_id, balance, currency, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, balance, currency, created_at;
	`

	now := time.Now()
	row := db.QueryRow(query, wallet.UserID, wallet.Balance, wallet.Currency, now)

	var created Wallet
	err := row.Scan(&created.ID, &created.UserID, &created.Balance, &created.Currency, &created.CreatedAt)
	if err != nil {
		return Wallet{}, err
	}

	return created, nil
}

//GetWalletByID retrieves a wallet from teh DB by ID
func GetWalletByID(db *sql.DB, id int64) (Wallet, error) {
	query := `
		SELECT id, user_id, balance, currency, created_at
		FROM wallets
		WHERE id = $1;
	`

	var wallet Wallet
	err := db.QueryRow(query, id).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.Balance,
		&wallet.Currency,
		&wallet.CreatedAt,
	)
	if err != nil {
		return Wallet{}, err
	}

	return wallet, nil
}