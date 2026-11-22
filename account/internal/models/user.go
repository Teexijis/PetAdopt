package models

type User struct {
	ID           int64
	FullName     string
	Email        string
	Phone        *string
	PasswordHash string
	Role         string
}