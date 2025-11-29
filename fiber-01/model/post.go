package model

import "time"

type Post struct {
	ID uint
	UserId uint
	User User
	Title string
	CreatedAt time.Time
	UpdatedAt time.Time
}