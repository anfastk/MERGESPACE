package postgres

import (
	"context"
	"errors"

	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/apperrors"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Save(ctx context.Context, user *entity.User) error {
	dbModel := FromUserEntity(user)
	result := r.db.WithContext(ctx).Create(dbModel)
	if result.Error != nil {
		return result.Error
	}
	user.ID = entity.UserID(dbModel.UserID)
	return nil
}

func (r *AuthRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var model UserModel
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, result.Error
	}
	return model.ToEntity(), nil
}

func (r *AuthRepository) IDExists(ctx context.Context, id uint64) (bool, error) {
	var count int64
	r.db.WithContext(ctx).Model(&UserModel{}).Where("user_id = ?", id).Count(&count)
	return count > 0, nil
}

func (r *AuthRepository) CreateSession(ctx context.Context, s *entity.Session) error {
	model := &SessionModel{
		SessionID:    s.SessionID,
		UserID:       s.UserID,
		AccessToken:  s.AccessToken,
		RefreshToken: s.RefreshToken,
		ExpiresAt:    s.ExpiresAt,
		CreatedAt:    s.CreatedAt,
		IPAddress:    s.IPAddress,
		UserAgent:    s.UserAgent,
	}
	return r.db.WithContext(ctx).Create(model).Error
}