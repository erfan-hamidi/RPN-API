package models

import (
	"RPN/cmd/db"

	"golang.org/x/crypto/bcrypt"
	//"database/sql"
)


type User struct {
    ID        uint      `json:"ID"`
    Username  string    `json:"username"`
    Password  string    `json:"password"` // Password should be stored hashed
    IsAdmin   bool      `json:"admin"` // For admin identification
}

func FindUserByUsername(username string) (*User, error) {
    var user User
    db := db.GetDB()
    row := db.QueryRow("SELECT * FROM users WHERE username=$1", username)
    
    err := row.Scan(&user.Username, &user.Password,)
    
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func UsernameExist(username string) (bool, error) {
    db := db.GetDB()
    var count int
    err:= db.QueryRow("SELECT COUNT(USERNAME) FROM USERS WHERE USERNAME=$1", username).Scan(&count)
    if err != nil {
        return false, err
    }
    if count > 0 {
        return true, nil
    } else {
        return false, nil
    }
}

func CreateUser(username string, password string) error {
    db := db.GetDB()
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

	db.Exec("INSERT INTO users (username, password, isadmin) VALUES ($1, $2, false)",username, hashedPassword)
	return err
}