package dto

import "time"

type AuthTokens struct {
	AccessToken      string
	RefreshToken     string
	SessionID        string
	AccessExpiresAt  time.Time
	RefreshExpiresAt time.Time
}
