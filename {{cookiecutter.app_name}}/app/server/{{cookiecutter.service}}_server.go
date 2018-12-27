package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gofunct/common/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	api_pb "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/api"
)

// {{cookiecutter.service}}ServiceServer is a composite interface of api_pb.{{cookiecutter.service}}ServiceServer and runtime.Server.
type {{cookiecutter.service}}ServiceServer interface {
	api_pb.{{cookiecutter.service}}ServiceServer
	runtime.Server
}

// New{{cookiecutter.service}}ServiceServer creates a new {{cookiecutter.service}}ServiceServer instance.
func New{{cookiecutter.service}}ServiceServer() {{cookiecutter.service}}ServiceServer {
	return &userServiceServerImpl{}
}

type userServiceServerImpl struct {
}

func (s *userServiceServerImpl) List{{cookiecutter.service}}s(ctx context.Context, req *api_pb.List{{cookiecutter.service}}sRequest) (*api_pb.List{{cookiecutter.service}}sResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) Get{{cookiecutter.service}}(ctx context.Context, req *api_pb.Get{{cookiecutter.service}}Request) (*api_pb.{{cookiecutter.service}}, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) Create{{cookiecutter.service}}(ctx context.Context, req *api_pb.Create{{cookiecutter.service}}Request) (*api_pb.{{cookiecutter.service}}, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) Update{{cookiecutter.service}}(ctx context.Context, req *api_pb.Update{{cookiecutter.service}}Request) (*api_pb.{{cookiecutter.service}}, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) Delete{{cookiecutter.service}}(ctx context.Context, req *api_pb.Delete{{cookiecutter.service}}Request) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
