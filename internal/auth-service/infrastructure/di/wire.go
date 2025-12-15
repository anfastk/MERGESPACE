package di

import (
	"log"
	"os"

	authpb "github.com/anfastk/MERGESPACE/api/proto/v1"
	"github.com/anfastk/MERGESPACE/internal/auth-service/adapter/inbound/grpc"
	"github.com/anfastk/MERGESPACE/internal/auth-service/adapter/outbound/idgen"
	"github.com/anfastk/MERGESPACE/internal/auth-service/adapter/outbound/kafka"
	"github.com/anfastk/MERGESPACE/internal/auth-service/adapter/outbound/otp"
	"github.com/anfastk/MERGESPACE/internal/auth-service/adapter/outbound/postgres"
	adapterRedis "github.com/anfastk/MERGESPACE/internal/auth-service/adapter/outbound/redis"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/service"
	"github.com/anfastk/MERGESPACE/internal/auth-service/infrastructure/config"
	"github.com/anfastk/MERGESPACE/internal/auth-service/infrastructure/database"
	"github.com/anfastk/MERGESPACE/internal/auth-service/infrastructure/redis"
	"github.com/anfastk/MERGESPACE/shared/kafka/producer"
)

type Container struct {
	AuthService authpb.AuthServiceServer
}

func InitContainer() *Container {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.NewPostgres(database.PostgresConfig{
		DSN: cfg.DatabaseURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	userRepo := postgres.NewUserRepository(pool)

	idGen := idgen.NewUUIDGenerator()
	otpGen := otp.NewCryptoOTPGenerator()
	usernameGen := service.NewUsernameGenerator(userRepo)

	schemaBytes, err := os.ReadFile("schemas/user_signup.avsc")
	if err != nil {
		log.Fatal(err)
	}
	schemaStr := string(schemaBytes)

	prod, err := producer.New(
		cfg.Kafka.Brokers,
		cfg.Kafka.SchemaRegistryURL,
		cfg.Kafka.UserSignupTopic,
		schemaStr,
	)
	if err != nil {
		log.Fatal(err)
	}

	otpPublisher := kafka.NewSignupEventProducer(prod)

	redisClient := redis.NewRedis(redis.RedisConfig{
		Addr: cfg.Redis.Addr,
		DB:   0,
	})

	pendingSignupRepo := adapterRedis.NewPendingSignupRepository(redisClient)

	authService := service.NewAuthService(
		userRepo,
		usernameGen,
		otpGen,
		idGen,
		pendingSignupRepo,
		otpPublisher,
	)
	authHandler := grpc.NewSignupHandler(*authService)

	return &Container{
		AuthService: authHandler,
	}
}
