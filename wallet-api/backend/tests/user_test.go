package test

import (
	"testing"
	"os"
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

func TestCreateAndFetchUser(t *testing.T) {
	setupDB(t)

	//Arrange
	user := models.User {
		Name: "Nicholas",
		Email: "nicholas@example.com",
	}

	//Act
	createdUser, err := models.CreateUser(db.DB, user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	fetchedUser, err := models.GetUserByID(db.DB, createdUser.ID)
	if err != nil {
		t.Fatalf("Failed to fetch user: %v", err)
	}

	if fetchedUser.Email != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, fetchedUser.Email)
	}
}