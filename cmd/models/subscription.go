package models

import (
	"RPN/cmd/db"
	"time"
)

type Subscription struct {
    ID         int
    UserID     int       `json:"ID"`		// Foreign key to User
    Capacity   int       `json:"capacity"`	// Maximum allowed requests
    Remaining  int       `json:"remainig"`	// Requests left
    ExpiryDate time.Time `json:"expiry"`	// Expiry date of the subscription
    IsActive   bool      `json:"active"`	// True if subscription is still active
}


func GetSubscription(userID int) (*Subscription, error) {
    db := db.GetDB()
    row := db.QueryRow("SELECT * FROM subcription WHERE user_id=$1 AND is_active = true LIMIT 1", userID)
    var subcription Subscription
    err := row.Scan(&subcription.ID,&subcription.UserID,&subcription.Capacity, &subcription.Remaining, &subcription.ExpiryDate, &subcription.IsActive)
    if err != nil {
        return nil, err
    }
    return &subcription, nil
}