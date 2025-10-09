package userservice

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "github.com/maehiyu/tollo/gen/go/protos/userservice"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer() *Server {
	return &Server{}
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error) {
	return nil, nil
}
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfo, error) {
	return nil, nil
}
func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserInfo, error) {
	return nil, nil
}
func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, nil
}