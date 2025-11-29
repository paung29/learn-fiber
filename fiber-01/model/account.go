package model

import "time"


type Account struct {
	ID uint 
	Email string
	Password string
	User []User
	CreatedAt time.Time
	UpdatedAt time.Time
}