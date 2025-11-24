package postgres

import "github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"

func (m *UserModel) ToEntity() *entity.User {
	return &entity.User{
		ID:            entity.UserID(m.UserID),
		Email:         m.Email,
		Username:      m.Username,
		PasswordHash:  m.PasswordHash,
		AccountStatus: m.AccountStatus,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

func FromUserEntity(e *entity.User) *UserModel {
	return &UserModel{
		Email:         e.Email,
		Username:      e.Username,
		PasswordHash:  e.PasswordHash,
		AccountStatus: e.AccountStatus,
	}
}
