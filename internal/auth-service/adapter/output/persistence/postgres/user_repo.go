package postgres

import (
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/output"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) output.UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var u entity.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
