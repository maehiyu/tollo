//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	userpb "github.com/maehiyu/tollo/gen/go/protos/userservice"
	chatpb "github.com/maehiyu/tollo/gen/go/protos/chatservice"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UserClient userpb.UserServiceClient
	ChatClient chatpb.ChatServiceClient
}