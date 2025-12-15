package main

import (
	"log"
	"net"

	authpb "github.com/anfastk/MERGESPACE/api/proto/v1"
	"github.com/anfastk/MERGESPACE/internal/auth-service/infrastructure/di"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	container := di.InitContainer()

	grpcServer := grpc.NewServer()

	authpb.RegisterAuthServiceServer(
		grpcServer,
		container.AuthService,
	)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Auth Service gRPC running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
