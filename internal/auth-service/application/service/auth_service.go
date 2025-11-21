package service

import (
	"errors"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/output"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo  output.UserRepository
	token output.TokenService
}

func NewAuthService(repo output.UserRepository, token output.TokenService) *AuthService {
	return &AuthService{repo: repo, token: token}
}

func (s *AuthService) Signup(input dto.SignupInput) (*dto.SignupOutput, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	user := &entity.User{
		Email:    input.Email,
		Username: input.Username,
		Password: string(hashed),
	}

	err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}

	return &dto.SignupOutput{
		UserID: user.ID,
		Email:  user.Email,
	}, nil
}

func (s *AuthService) Login(input dto.LoginInput) (*dto.LoginOutput, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return nil, errors.New("invalid password")
	}

	access, refresh, _ := s.token.GenerateTokens(user.ID, user.Email)

	return &dto.LoginOutput{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
