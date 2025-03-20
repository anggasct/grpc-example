package main

import (
	"fmt"
	"log"
	"net"

	"github.com/anggasct/user-service/internal/config"
	"github.com/anggasct/user-service/internal/db"
	"github.com/anggasct/user-service/internal/middleware"
	"github.com/anggasct/user-service/internal/service"
	pb "github.com/anggasct/user-service/pb/user"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	cfg := config.LoadConfig()

	database, err := db.NewDatabase(cfg.GetDSN())
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.ServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.PrettyLoggingInterceptor()),
	)
	userService := service.NewUserService(database)
	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Printf("Starting gRPC server on port %s", cfg.ServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
