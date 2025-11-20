package domain

import "github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"

type UserRepository interface {
	Save(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
