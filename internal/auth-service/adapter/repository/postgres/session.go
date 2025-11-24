package postgres

import "time"

type SessionModel struct {
	SessionID    string    `gorm:"primaryKey;type:uuid;column:session_id"`
	UserID       uint64    `gorm:"not null"`
	AccessToken  string    `gorm:"size:500"`
	RefreshToken string    `gorm:"size:500"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	IPAddress    string    `gorm:"type:inet"`
	UserAgent    string
}

func (SessionModel) TableName() string {
	return "user_sessions"
}
