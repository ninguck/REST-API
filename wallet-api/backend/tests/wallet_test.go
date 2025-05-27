package tests

import (
	//"os"
	"testing"
	"fmt"
	"time"

	"github.com/ninguck/wallet-api/internal/db"
	"github.com/ninguck/wallet-api/internal/models"
)

/*func setupDB(t *testing.T) {
	os.Setenv("DB_USER", "wallet_user")
	os.Setenv("DB_PASSWORD", "wallet_pass")
	os.Setenv("DB_NAME", "wallet_db")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")

	if err := db.InitDB(); err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}
}*/

func TestCreateAndGetWallet(t *testing.T) {
	setupDB(t)

	// First, create a user to attach the wallet to
	user := models.User{
		Name:  "WalletTester",
		Email: fmt.Sprintf("wallet+%d@test.com", time.Now().UnixNano()),
	}
	createdUser, err := models.CreateUser(db.DB, user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Define a new wallet
	wallet := models.Wallet{
		UserID:   createdUser.ID,
		Balance:  1000.00,
		Currency: "AUD",
	}

	// Create the wallet
	createdWallet, err := models.CreateWallet(db.DB, wallet)
	if err != nil {
		t.Fatalf("Failed to create wallet: %v", err)
	}

	// Fetch the wallet by ID
	fetchedWallet, err := models.GetWalletByID(db.DB, createdWallet.ID)
	if err != nil {
		t.Fatalf("Failed to fetch wallet: %v", err)
	}

	// Verify that values match
	if fetchedWallet.UserID != wallet.UserID || fetchedWallet.Currency != wallet.Currency {
		t.Errorf("Expected wallet %+v, got %+v", wallet, fetchedWallet)
	}
}
