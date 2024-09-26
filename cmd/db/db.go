// db/db.go
package db

import (
    "database/sql"
    "log"
)

var conn *sql.DB

// Init initializes the database connection
func Init() {
	dataSourceName := ""
    var err error
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