package tests

import (
	"os"
	"testing"
	

	"github.com/ninguck/wallet-api/internal/db"
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
