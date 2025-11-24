package database

import (
	"github.com/anfastk/MERGESPACE/internal/auth-service/adapter/repository/postgres"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&postgres.UserModel{},
		&postgres.SessionModel{},
	)
}
