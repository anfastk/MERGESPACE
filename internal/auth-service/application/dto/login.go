package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	UserAgent string 
	IPAddress string 
}
