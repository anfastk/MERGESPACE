package entity

import "time"

type UserID uint64

type User struct {
	ID            UserID
	Email         string
	Username      string
	PasswordHash  string
	AccountStatus string
	GoogleOAuthID string
	GithubOAuthID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewUser(email, username, passwordHash string) *User {
	return &User{
		Email:         email,
		Username:      username,
		PasswordHash:  passwordHash,
		AccountStatus: "Active",
	}
}
