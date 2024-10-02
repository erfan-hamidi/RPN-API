package models

import "time"

type RequestLog struct {
    ID         int		`json:"ID"`
	Request    string	`json:"requset"`
    UserID     int      `json:"userid"`	// Foreign key to User
    RequestedAt time.Time 				// Timestamp of the request
}
