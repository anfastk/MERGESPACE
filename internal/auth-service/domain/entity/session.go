package entity

import "time"

type Session struct {
	SessionID    string
	UserID       uint64
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
	IPAddress    string
	UserAgent    string
}
