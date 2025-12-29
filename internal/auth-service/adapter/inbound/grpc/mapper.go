package grpc

import (
	domainErr "github.com/anfastk/MERGESPACE/internal/auth-service/domain/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func mapError(err error) error {
	switch err {
	case domainErr.ErrEmailAlreadyExists:
		return status.Error(codes.AlreadyExists, err.Error())
	case domainErr.ErrInvalidEmail:
		return status.Error(codes.InvalidArgument, err.Error())
	case domainErr.ErrTooManyRequests:
		return status.Error(codes.ResourceExhausted, err.Error())
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
