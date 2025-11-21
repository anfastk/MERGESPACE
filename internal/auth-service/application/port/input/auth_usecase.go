package input

import "github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"

type AuthUseCase interface {
	Signup(input dto.SignupInput) (*dto.SignupOutput, error)
	Login(input dto.LoginInput) (*dto.LoginOutput, error)
}
