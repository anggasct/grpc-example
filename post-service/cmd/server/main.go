package main

import (
	"fmt"
	"log"
	"net"

	"github.com/anggasct/post-service/internal/config"
	"github.com/anggasct/post-service/internal/db"
	"github.com/anggasct/post-service/internal/service"
	pb "github.com/anggasct/post-service/pb/post"

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

	grpcServer := grpc.NewServer()
	postService := service.NewPostService(database, cfg.GetUserServiceAddress())
	pb.RegisterPostServiceServer(grpcServer, postService)

	log.Printf("Starting gRPC server on port %s", cfg.ServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
