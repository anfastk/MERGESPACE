package redis

import "time"

type pendingSignupModel struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	OTP          string    `json:"otp"`
	Attempts     int       `json:"attempts"`
	ResendCount  int       `json:"resend_count"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}


