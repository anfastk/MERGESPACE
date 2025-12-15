package avro

type SignupOTPEventAvro struct {
	SignupSessionID string `avro:"signup_session_id"`
	Email           string `avro:"email"`
	OTP             string `avro:"otp"`
	CreatedAt       int64  `avro:"created_at"`
}