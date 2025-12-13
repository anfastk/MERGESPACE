package entity

import (
	"time"

	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/valueobject"
)

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusPending   UserStatus = "pending"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusDeleted   UserStatus = "deleted"
)

type AuthProvider string

const (
	AuthProviderLocal  AuthProvider = "email"
	AuthProviderGoogle AuthProvider = "google"
	AuthProviderGithub AuthProvider = "github"
)

type User struct {
	UserID       valueobject.UserID
	FirstName    valueobject.Name
	LastName     valueobject.Name
	Username     valueobject.Username
	Email        valueobject.Email
	AuthProvider AuthProvider
	Password     *valueobject.Password
	ProviderID   *string
	Status       UserStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func NewLocalUser(id valueobject.UserID, email valueobject.Email, username valueobject.Username, password valueobject.Password, now time.Time) *User {
	return &User{
		UserID:       id,
		Email:        email,
		Username:     username,
		AuthProvider: AuthProviderLocal,
		Password:     &password,
		Status:       UserStatusActive,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func NewOAuthUser(id valueobject.UserID, email valueobject.Email, username valueobject.Username, provider AuthProvider, providerID string, now time.Time) *User {
	return &User{
		UserID:       id,
		Email:        email,
		Username:     username,
		AuthProvider: provider,
		ProviderID:   &providerID,
		Status:       UserStatusActive,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
