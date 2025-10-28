package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/maehiyu/tollo/gen/go/protos/chatservice"
	"github.com/maehiyu/tollo/internal/adapter/repository"
	"github.com/maehiyu/tollo/internal/chatservice"
)

func main() {
	chatRepo := repository.NewChatRepositoryMock()
	chatUsecase := chatservice.NewUsecase(chatRepo)
	chatServiceServer := chatservice.NewServer(chatUsecase)

	port := ":50052"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, chatServiceServer)

	reflection.Register(grpcServer)

	go func() {
		log.Printf("gRPC server is running on port %s", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gRPC server...")

	grpcServer.GracefulStop()
	log.Println("gRPC server gracefully stopped")
}
