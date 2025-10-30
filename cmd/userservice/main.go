package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/maehiyu/tollo/gen/go/protos/userservice"
	"github.com/maehiyu/tollo/internal/adapter/repository"
	"github.com/maehiyu/tollo/internal/userservice"
)

func main() {
	userRepo := repository.NewMockUserRepository()
	userUsecase := userservice.NewUsecase(userRepo)
	userServiceServer := userservice.NewUserServiceServer(userUsecase)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServiceServer)

	reflection.Register(grpcServer)

	go func() {
		log.Println("gRPC server is running on port :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gRPC server...")

	grpcServer.GracefulStop()
	log.Println("gRPC server gracefully stopped")
}
