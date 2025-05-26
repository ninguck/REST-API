package models

import (
	"database/sql"
	"time"
)

//User represents a system user
type User struct {
	ID 			int64 		`json:"id"`
	Name 		string 		`json:"name"`
	Email 		string		`string:"email"`
	CreatedAt 	time.Time	`json:"created_at"`
}

//CreateUser inserts a new userinto the data base and returns the created record
func CreateUser(db *sql.DB, user User) (User, error) {
	query := `
		INSERT INTO users (name, email, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, created_at
		`

		now := time.Now()
		row := db.QueryRow(query, user.Name, user.Email, now)

		var created User
		err := row.Scan(&created.ID, &created.Name, &created.Email, &created.CreatedAt)
		if err != nil {
			return User{}, err
		}

		return created, nil
}


//GetUserByID fetches a user from the data base by thier ID
func GetUserByID(db *sql.DB, id int64) (User, error) {
	query := `
	SELECT id, name, email, created_at
	FROM users
	WHERE id = $1;
	`

	var user User
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}