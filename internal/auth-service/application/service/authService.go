package service

import (
	"context"
	"time"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/outbound"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/errs"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/valueobject"
	"golang.org/x/crypto/bcrypt"
)

const (
	otpTTLMinutes     = 5
	maxOTPAttempts    = 5
	maxOTPResendCount = 3
	sessionTTLDays    = 30
)

type AuthService struct {
	user           outbound.UserRepository
	usernameGen    *UsernameGenerator
	otpGen         outbound.OTPGenerator
	idGen          outbound.IDGenerator
	pendingSignups outbound.PendingSignupRepository
}

func NewAuthService(user outbound.UserRepository, usernameGen *UsernameGenerator, otpGen outbound.OTPGenerator, idGen outbound.IDGenerator, pendingSignups outbound.PendingSignupRepository) *AuthService {
	return &AuthService{
		user:           user,
		usernameGen:    usernameGen,
		otpGen:         otpGen,
		idGen:          idGen,
		pendingSignups: pendingSignups,
	}	
}

func (s *AuthService) InitiateSignup(ctx context.Context, req dto.InitiateSignUpRequest) (*dto.InitiateSignUpResponse, error) {

	email, err := valueobject.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}

	firstname, err := valueobject.NewName(req.FirstName)
	if err != nil {
		return nil, err
	}

	lastname, err := valueobject.NewName(req.LastName)
	if err != nil {
		return nil, err
	}

	username, err := s.usernameGen.Generate(ctx, req.FirstName, req.LastName)
	if err != nil {
		return nil, err
	}

	userName, err := valueobject.NewUsername(username)
	if err != nil {
		return nil, err
	}

	password, err := valueobject.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}

	if u, err := s.user.FindByEmail(ctx, email.String()); err != nil {
		return nil, err
	} else if u != nil {
		return nil, errs.ErrEmailAlreadyExists
	}

	if u, err := s.user.FindByUsername(ctx, userName.String()); err != nil {
		return nil, err
	} else if u != nil {
		return nil, errs.ErrUsernameExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password.String()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	tempID := entity.PendingSignupID(s.idGen.NewID())
	otp, err := s.otpGen.Generate()
	if err != nil {
		return nil, err
	}

	pending := &entity.PendingSignup{
		ID:           tempID,
		Email:        email.String(),
		FirstName:    firstname.String(),
		LastName:     lastname.String(),
		Username:     userName.String(),
		PasswordHash: string(passwordHash),
		OTP:          otp,
		CreatedAt:    now,
		ExpiresAt:    now.Add(otpTTLMinutes * time.Minute),
	}

	if err := s.pendingSignups.Save(ctx, pending); err != nil {
		return nil, err
	}

	return &dto.InitiateSignUpResponse{
		SignupSessionID: string(tempID),
		Status:          dto.SignupStatusOtpSent,
	}, nil
}
