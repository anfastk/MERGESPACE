package grpc

import (
	"context"

	authpb "github.com/anfastk/MERGESPACE/api/proto/v1"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/service"
)

type SignupHandler struct {
	authpb.UnimplementedAuthServiceServer
	usecase service.AuthService
}

func NewSignupHandler(usecase service.AuthService) *SignupHandler {
	return &SignupHandler{usecase: usecase}
}

func (h *SignupHandler) InitiateSignup(ctx context.Context, req *authpb.SignUpRequest) (*authpb.SignUpResponce, error) {
	res, err := h.usecase.InitiateSignup(ctx, dto.InitiateSignUpRequest{
		Email:     req.Email,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Password:  req.Password,
	})
	if err != nil {
		return nil, mapError(err)
	}

	return &authpb.SignUpResponce{
		SignupSessionId: res.SignupSessionID,
		Status:          authpb.SignupStatus(dto.SignupStatusOTPSent),
	}, nil
}
