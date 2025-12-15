package dto

type SignupOTPEvent struct {
	SignupSessionID string
	Email           string
	OTP             string
	CreatedAt       int64
}
	