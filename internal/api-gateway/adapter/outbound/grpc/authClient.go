package grpc

import (
	authpb "github.com/anfastk/MERGESPACE/api/proto/v1"
	"google.golang.org/grpc"
)

type AuthClient struct {
	Client authpb.AuthServiceClient
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	return &AuthClient{
		Client: authpb.NewAuthServiceClient(conn),
	}
}
