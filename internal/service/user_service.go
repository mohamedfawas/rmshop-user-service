package service

import (
	"context"

	userv1 "github.com/mohamedfawas/rmshop-proto/gen/v1/user"
	"github.com/mohamedfawas/rmshop-user-service/internal/domain"
	"github.com/mohamedfawas/rmshop-user-service/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	repo domain.UserRepository
	userv1.UnimplementedUserServiceServer
}

func NewUserService(repo domain.UserRepository) userv1.UserServiceServer {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "missing required fields")
	}

	// Hash password
	passwordHash, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	user := &domain.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &userv1.CreateUserResponse{
		UserId: user.ID,
	}, nil
}
