package user

import (
	"fmt"
	user_pb "github.com/gofunct/example/gen/user"
	"context"
)



type Service struct{}

func New() user_pb.UserServiceServer { return &Service{} }

func (svc *Service) CreateUser(ctx context.Context, in *user_pb.CreateUserRequest) (*user_pb.CreateUserResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) GetUser(ctx context.Context, in *user_pb.GetUserRequest) (*user_pb.GetUserResponse, error) {
	return nil, fmt.Errorf("not implemented")
}