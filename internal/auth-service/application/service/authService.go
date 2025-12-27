package service

import (
	"context"
	"log"
	"time"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/outbound"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/errs"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/valueobject"
	"github.com/anfastk/MERGESPACE/shared/ratelimiter/limiter"
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
	otpPublisher   outbound.OTPEventPublisher
	rateLimiter    *limiter.Limiter
}

func NewAuthService(user outbound.UserRepository, usernameGen *UsernameGenerator, otpGen outbound.OTPGenerator, idGen outbound.IDGenerator, pendingSignups outbound.PendingSignupRepository, pub outbound.OTPEventPublisher, rateLimiter *limiter.Limiter) *AuthService {
	return &AuthService{
		user:           user,
		usernameGen:    usernameGen,
		otpGen:         otpGen,
		idGen:          idGen,
		pendingSignups: pendingSignups,
		otpPublisher:   pub,
		rateLimiter:    rateLimiter,
	}
}

func (s *AuthService) InitiateSignup(ctx context.Context, req dto.InitiateSignUpRequest) (*dto.InitiateSignUpResponse, error) {
	/* 	ok, _, err := s.rateLimiter.Allow(
	   		ctx,
	   		limiter.SignupIPRule,
	   		clientIP,
	   	)
	   	if err != nil || !ok {
	   		return nil, errs.ErrTooManyRequests
	   	} */

	ok, _, err := s.rateLimiter.Allow(
		ctx,
		limiter.SignupEmailRule,
		req.Email,
	)
	if err != nil || !ok {
		return nil, errs.ErrTooManyRequests
	}

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
	log.Println("1. Redis save start")

	if err := s.pendingSignups.Save(ctx, pending); err != nil {
		return nil, err
	}
	log.Println("1. Redis save end", err)

	event := dto.SignupOTPEvent{

		SignupSessionID: string(tempID),
		Email:           email.String(),
		OTP:             otp,
		CreatedAt:       time.Now().Unix(),
	}
	log.Println("2. Kafka publish start")

	if err = s.otpPublisher.PublishOTPEvent(ctx, event); err != nil {
		return nil, err
	}
	log.Println("2. Kafka publish end", err)

	return &dto.InitiateSignUpResponse{
		SignupSessionID: string(tempID),
		Status:          dto.SignupStatusOTPSent,
	}, nil
}
