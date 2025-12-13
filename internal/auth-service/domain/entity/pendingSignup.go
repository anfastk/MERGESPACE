package entity

import "time"

type PendingSignupID string

type PendingSignup struct {
	ID           PendingSignupID
	Email        string
	FirstName    string
	LastName     string
	Username     string
	PasswordHash string
	OTP          string
	Attempts     int
	ResendCount  int
	CreatedAt    time.Time
	ExpiresAt    time.Time
}
