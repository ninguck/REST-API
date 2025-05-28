package tests

import (
	"os"
	"testing"
	"time"
	"fmt"
	

	"github.com/ninguck/wallet-api/internal/db"
	"github.com/ninguck/wallet-api/internal/models"
	
)

func setupDB(t *testing.T) {
	os.Setenv("DB_USER", "wallet_user")
	os.Setenv("DB_PASSWORD", "wallet_pass")
	os.Setenv("DB_NAME", "wallet_db")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")

	if err := db.InitDB(); err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}
}

func generateRandomUser () models.User {
	return models.User {
		Name: fmt.Sprintf("TestUser_%d", time.Now().UnixNano()),
		Email: fmt.Sprintf("test_%d@example.com", time.Now().UnixNano()),
	}
}


func generateWalletForUser(userID int64) models.Wallet {
	return models.Wallet {
		UserID: userID,
		Balance: 500.00,
		Currency: "AUD",
	}
}
