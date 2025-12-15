package dto

type SignupStatus int

const (
	SignupStatusUnspecified SignupStatus = iota
	SignupStatusOTPSent
	SignupStatusRateLimited
	SignupStatusInvalidPhone
	SignupStatusInternalError
)

type InitiateSignUpRequest struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

type InitiateSignUpResponse struct {
	SignupSessionID string
	Status          SignupStatus
}
