// db/db.go
package db

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var conn *sql.DB

// Init initializes the database connection
func Init() {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	//dbHost := os.Getenv("DB_HOST")
    //dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
	print(dbUser)
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	dbUser, dbPassword, dbName)
    //var err error
    conn, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }

    if err := conn.Ping(); err != nil {
        log.Fatal(err)
    }
}

// GetDB returns the database connection
func GetDB() *sql.DB {
    return conn
}