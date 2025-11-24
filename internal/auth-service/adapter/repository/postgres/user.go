package postgres

import "time"

type UserModel struct {
	UserID        uint64    `gorm:"primaryKey;autoIncrement;column:user_id"`
	Email         string    `gorm:"unique;not null"`
	Username      string    `gorm:"unique;not null"`
	PasswordHash  string    `gorm:"not null"`
	AccountStatus string    `gorm:"default:'active'"`
	CreatedAt     time.Time `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time `gorm:"default:current_timestamp"`
}

func (UserModel) TableName() string {
	return "users"
}
