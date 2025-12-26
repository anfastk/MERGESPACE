package dto

type SignupOTPEvent struct {
	EventID   string
	Email     string
	OTP       string
	CreatedAt int64
}
