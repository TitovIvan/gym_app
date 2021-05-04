package store

import "time"

// User contain user-related info
type User struct {
	ID       uint64 `json:"-"`
	Name     string `json:"name"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Age      uint8  `json:"age"`
	Admin    bool   `json:"admin"`
	// Admin     bool      `json:"admin"  binding:"required"`
	Blocked   bool      `json:"blocked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
