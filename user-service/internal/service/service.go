package service

import (
	"context"
	"fmt"

	"github.com/anggasct/user-service/internal/db"
	pb "github.com/anggasct/user-service/pb/user"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	db *db.Database
}

func NewUserService(database *db.Database) *UserService {
	return &UserService{
		db: database,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.db.GetUser(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:   user.ID,
			Name: user.Name,
		},
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	user, err := s.db.CreateUser(req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:   user.ID,
			Name: user.Name,
		},
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, totalCount, err := s.db.ListUsers(req.PageSize, req.Page)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %v", err)
	}

	pbUsers := make([]*pb.User, len(users))
	for i, user := range users {
		pbUsers[i] = &pb.User{
			Id:   user.ID,
			Name: user.Name,
		}
	}

	return &pb.ListUsersResponse{
		Users:      pbUsers,
		TotalCount: totalCount,
	}, nil
}
