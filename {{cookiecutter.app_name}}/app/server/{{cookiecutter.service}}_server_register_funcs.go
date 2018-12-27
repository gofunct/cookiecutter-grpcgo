package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/api"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *{{cookiecutter.service}}ServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterUserServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *{{cookiecutter.service}}ServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterUserServiceHandler(ctx, mux, conn)
}
