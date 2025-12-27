package main

import (
	"log"

	httpAdapter "github.com/anfastk/MERGESPACE/internal/api-gateway/adapter/inbound/http"
	grpcAdapter "github.com/anfastk/MERGESPACE/internal/api-gateway/adapter/outbound/grpc"
	"github.com/anfastk/MERGESPACE/internal/api-gateway/infrastructure/config"
)

func main() {
	cfg := config.Load()

	conn := grpcAdapter.NewConn(cfg.AuthGRPC)
	authClient := grpcAdapter.NewAuthClient(conn)

	authHandler := httpAdapter.NewAuthHandler(authClient.Client)
	router := httpAdapter.NewRouter(authHandler)

	log.Println("API Gateway running on :" + cfg.HTTPPort)
	router.Run(":" + cfg.HTTPPort)
}
