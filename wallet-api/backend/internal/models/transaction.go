package models

import (
	"time"
	"database/sql"
)

type Transaction struct {
	ID        int64     `json:"id"`
	WalletID  int64     `json:"wallet_id"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"` // "credit" or "debit"
	CreatedAt time.Time `json:"created_at"`
}

func CreateTransaction(db *sql.DB, txn Transaction) (Transaction, error) {
	query := `
	INSERT INTO transactions (wallet_id, amount, type, created_at) 
	VALUES ($1, $2, $3, $4) RETURNING id
	`

	err := db.QueryRow(query, txn.WalletID, txn.Amount, txn.Type, time.Now()).Scan(&txn.ID)
	txn.CreatedAt = time.Now()
	return txn, err
}

func GetTransactionByID(db *sql.DB, id int64) (Transaction, error) {
	var txn Transaction
	query := `SELECT id, wallet_id, amount, type, created_at FROM transactions WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&txn.ID, &txn.WalletID, &txn.Amount, &txn.Type, &txn.CreatedAt)
	return txn, err
}