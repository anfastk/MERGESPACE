package dto

type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginOutput struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}
