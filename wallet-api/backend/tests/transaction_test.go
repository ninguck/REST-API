package tests

import (
	"testing"

    "github.com/ninguck/wallet-api/internal/db"
    "github.com/ninguck/wallet-api/internal/models"
)

func TestCreateAndFetchTransaction(t *testing.T) {
	setupDB(t)

	//Create user
	user := generateRandomUser()
	createdUser, err := models.CreateUser(db.DB, user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	//Create wallet
	wallet := generateWalletForUser(createdUser.ID)
	createdWallet, err := models.CreateWallet(db.DB, wallet)
	if err != nil {
		t.Fatalf("Failed to create wallet: %v", err)
	}

	//Create transaction
	transaction := models.Transaction {
		WalletID: createdWallet.ID,
		Amount: 150.00,
		Type: "credit",
	}

	createdTxn, err := models.CreateTransaction(db.DB, transaction)
	if err != nil {
		t.Fatalf("Failed to fetch transaction: %v", err)
	}

	//Fetch transaction
	fetchedTxn, err:= models.GetTransactionByID(db.DB, createdTxn.ID)
	if err!= nil {
		t.Fatalf("Failed to fetch transaction: %v", err)
	}

	if fetchedTxn.Amount != transaction.Amount || fetchedTxn.Type != transaction.Type {
		t.Errorf("Expected %+v, got %+v", transaction, fetchedTxn)
	}
}