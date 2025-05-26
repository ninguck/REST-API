package db

import (
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	os.Setenv("DB_USER", "wallet_user")
	os.Setenv("DB_PASSWORD", "wallet_pass")
	os.Setenv("DB_NAME", "wallet_db")

	err:= InitDB()
	if err != nil {
		t.Fatalf("failed to connect to DB: %v", err)
	}
}