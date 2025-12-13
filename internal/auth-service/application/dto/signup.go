package dto

type SignupStatus string

const (
	SignupStatusUnspecified   SignupStatus = "UNSPECIFIED"
	SignupStatusOtpSent       SignupStatus = "OTP_SENT"
	SignupStatusRateLimited   SignupStatus = "RATE_LIMITED"
	SignupStatusInvalidPhone  SignupStatus = "INVALID_PHONE"
	SignupStatusInternalError SignupStatus = "INTERNAL_ERROR"
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
