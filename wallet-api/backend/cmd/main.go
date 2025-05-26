package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ninguck/wallet-api/internal/db"
	"github.com/ninguck/wallet-api/internal/api"

)

func main() {
	//Load enviroment variables
	os.Setenv("DB_USER", "wallet_user")
	os.Setenv("DB_PASSWORD", "wallet_pass")
	os.Setenv("DB_NAME", "wallet_db")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")


	err:= db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Setup router
	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/ping", api.PingHandler).Methods("GET")

	// User API
	r.HandleFunc("/users", api.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", api.GetUserHandler).Methods("GET")

	// Start server
	fmt.Println("🚀 Server is running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
