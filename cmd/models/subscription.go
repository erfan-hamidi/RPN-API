package models

import "time"

type Subscription struct {
    ID         int
    UserID     int       `json:"ID"`		// Foreign key to User
    Capacity   int       `json:"capacity"`	// Maximum allowed requests
    Remaining  int       `json:"remainig"`	// Requests left
    ExpiryDate time.Time `json:"expiry"`	// Expiry date of the subscription
    IsActive   bool      `json:"active"`	// True if subscription is still active
}
