package models

import (
	"database/sql"
)


type User struct {
    ID        uint      `json:"ID"`
    Username  string    `json:"username"`
    Password  string    `json:"password"` // Password should be stored hashed
    IsAdmin   bool      `json:"admin"` // For admin identification
}

func FindUserByUsername(db *sql.DB, username string) (*User, error) {
    var user User
    row := db.QueryRow("SELECT * FROM users WHERE username=$1", username)
    
    err := row.Scan(&user.Username, &user.Password,)
    
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func SaveUser(db *sql.DB, user *User) error {
	_,err := db.Exec("INSERT INTO users (username, public_key, private_key) VALUES ($1, $2, $3)",user.Username,user.Password)
	return err
}