package main

import (
	"context"
	"log"

	"github.com/anfastk/MERGESPACE/internal/notification-service/adapter/inbound/kafka"
	"github.com/anfastk/MERGESPACE/internal/notification-service/application/service"
)

func main() {
	ctx := context.Background()

	handler := service.NewOTPHandler()

	consumer, err := kafka.NewSignupConsumer(
		[]string{"localhost:29092"},
		"notification-service",
		"user.signup",
		"http://localhost:8081",
		handler,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Kafka consumer started")
	consumer.Start(ctx)
}
