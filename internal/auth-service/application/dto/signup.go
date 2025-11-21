package dto

type SignupInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Username string `json:"username"`
}

type SignupOutput struct {
    UserID int64  `json:"user_id"`
    Email  string `json:"email"`
}
