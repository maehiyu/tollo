package userservice

import (
	"context"
	"errors"

	pb "github.com/maehiyu/tollo/gen/go/protos/userservice"
	"github.com/maehiyu/tollo/internal/adapter/converter"
	"github.com/maehiyu/tollo/internal/userservice/domain/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	usecase Usecase
}

func NewUserServiceServer(uc Usecase) *Server {
	return &Server{usecase: uc}
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error) {
	email, err := s.emailFromString(req.GetEmail())
	if err != nil {
		return nil, err
	}

	profile, err := s.profileFromCreateRequest(req)
	if err != nil {
		return nil, err
	}

	input := &CreateUserInput{
		Name:    req.GetName(),
		Email:   email,
		Profile: profile,
	}

	createdUser, err := s.usecase.CreateUser(ctx, input)
	if err != nil {
		if errors.Is(err, user.ErrEmailAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return converter.ToUserInfo(createdUser), nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfo, error) {
	var foundUser *user.User
	var err error

	switch v := req.GetLookupBy().(type) {
	case *pb.GetUserRequest_Id:
		id := v.Id
		if id == "" {
			return nil, status.Errorf(codes.InvalidArgument, "user id is required")
		}
		foundUser, err = s.usecase.GetUserByID(ctx, id)

	case *pb.GetUserRequest_Email:
		emailStr := v.Email
		email, err_conv := s.emailFromString(emailStr)
		if err_conv != nil {
			return nil, err_conv
		}
		foundUser, err = s.usecase.GetUserByEmail(ctx, email)

	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "lookup field (id or email) is required")

	default:
		return nil, status.Errorf(codes.Internal, "unknown lookup type")
	}
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	return converter.ToUserInfo(foundUser), nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserInfo, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user id is required")
	}
	if req.GetData() == nil || req.GetUpdateMask() == nil || len(req.GetUpdateMask().GetPaths()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update data and mask are required")
	}
	input := &UpdateUserInput{}

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "name":
			name := req.GetData().GetName()
			input.Name = &name
		case "profile":
			profile, err := s.profileFromUserInfo(req.GetData())
			if err != nil {
				return nil, err
			}
			input.Profile = &profile
		}
	}
	updatedUser, err := s.usecase.UpdateUser(ctx, id, input)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}
	return converter.ToUserInfo(updatedUser), nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user id is required")
	}

	err := s.usecase.DeleteUser(ctx, id)

	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) profileFromCreateRequest(req *pb.CreateUserRequest) (user.Profile, error) {
	switch v := req.GetProfile().(type) {
	case *pb.CreateUserRequest_Professional:
		return &user.ProfessionalProfile{
			ProBadgeURL: v.Professional.ProBadgeUrl,
			Biography:   v.Professional.Biography,
		}, nil
	case *pb.CreateUserRequest_General:
		return &user.GeneralProfile{
			Points:       v.General.Points,
			Introduction: v.General.Introduction,
		}, nil
	case nil:
		return &user.GeneralProfile{}, nil
	default:
		return nil, status.Errorf(codes.Internal, "unhandled profile type: %T", v)
	}
}

func (s *Server) profileFromUserInfo(userInfo *pb.UserInfo) (user.Profile, error) {
	switch v := userInfo.GetProfile().(type) {
	case *pb.UserInfo_ProfessionalProfile:
		return &user.ProfessionalProfile{
			ProBadgeURL: v.ProfessionalProfile.ProBadgeUrl,
			Biography:   v.ProfessionalProfile.Biography,
		}, nil
	case *pb.UserInfo_GeneralProfile:
		return &user.GeneralProfile{
			Points:       v.GeneralProfile.Points,
			Introduction: v.GeneralProfile.Introduction,
		}, nil
	case nil:
		return &user.GeneralProfile{}, nil
	default:
		return nil, status.Errorf(codes.Internal, "unhandled profile type: %T", v)
	}
}

func (s *Server) emailFromString(emailStr string) (user.Email, error) {
	email, err := user.NewEmail(emailStr)
	if err != nil {
		if errors.Is(err, user.ErrInvalidEmail) {
			return "", status.Errorf(codes.InvalidArgument, err.Error())
		}
		return "", status.Errorf(codes.Internal, "unexpected error during email validation: %v", err)
	}
	return email, nil
}
