package model

type User struct {
	ID uint
	Name string
	AccountId uint
	Account Account
	Posts []Post
}