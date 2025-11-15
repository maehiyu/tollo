package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	chatpb "github.com/maehiyu/tollo/gen/go/protos/chatservice"
	userpb "github.com/maehiyu/tollo/gen/go/protos/userservice"
	"github.com/maehiyu/tollo/internal/gateway/graph"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect to UserService
	userServiceAddr := os.Getenv("USER_SERVICE_ADDR")
	if userServiceAddr == "" {
		userServiceAddr = "localhost:50051"
	}
	userConn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to user service: %v", err)
	}
	defer userConn.Close()
	userClient := userpb.NewUserServiceClient(userConn)

	// Connect to ChatService
	chatServiceAddr := os.Getenv("CHAT_SERVICE_ADDR")
	if chatServiceAddr == "" {
		chatServiceAddr = "localhost:50052"
	}
	chatConn, err := grpc.NewClient(chatServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to chat service: %v", err)
	}
	defer chatConn.Close()
	chatClient := chatpb.NewChatServiceClient(chatConn)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			UserClient: userClient,
			ChatClient: chatClient,
		},
	}))

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		Debug:            true,
	}).Handler(mux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
