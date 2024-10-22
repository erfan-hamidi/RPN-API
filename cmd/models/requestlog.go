package models

import (
	"RPN/cmd/db"
	"time"
)

type RequestLog struct {
    ID         int		`json:"ID"`
	Request    string	`json:"requset"`
    UserID     int      `json:"userid"`	// Foreign key to User
    RequestedAt time.Time 				// Timestamp of the request
}

func logRequest(userID int) error {
    db := db.GetDB()
    query := `INSERT INTO request_logs (user_id, requested_at) VALUES ($1, $2)`
    _, err := db.Exec(query, userID, time.Now())
    return err
}