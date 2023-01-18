package model

import "time"

type User struct {
	ID       int       `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"-" db:"password"`
	Email    string    `json:"email" db:"email"`
	Created  time.Time `json:"created_at" db:"created_at"`
	Updated  time.Time `json:"updated_at" db:"updated_at"`
}
