package store

import "time"

// User contain user-related info
type User struct {
	ID        uint64    `json:"id"`
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Admin     bool      `json:"admin"`
	Blocked   bool      `json:"blocked"`
	Age       uint8     `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
