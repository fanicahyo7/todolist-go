package model

import "time"

type User struct {
	ID       int       `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Email    string    `json:"email" db:"email"`
	Created  time.Time `json:"created_at" db:"created_at"`
	Updated  time.Time `json:"updated_at" db:"updated_at"`
}

type UserRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

type UserResponse struct {
	ID       int       `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
	Created  time.Time `json:"created_at" db:"created_at"`
	Updated  time.Time `json:"updated_at" db:"updated_at"`
}
