package entity

import (
	"time"

	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/valueobject"
)

type SessionID string
type DeviceID string

type Session struct {
	SessionID SessionID
	UserID    valueobject.UserID
	DeviceID  DeviceID

	AccessToken  string 
	RefreshToken string 
		
	ExpiresAt    time.Time
	CreatedAt    time.Time
	RevokedAt    *time.Time
}
