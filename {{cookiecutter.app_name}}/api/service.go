package api

import (
	"fmt"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

)



type Service struct{}

func New() {{cookiecutter.service}}ServiceServer { return &Service{} }

func (svc *Service) List{{cookiecutter.service}}s(ctx context.Context, in *List{{cookiecutter.service}}Request, opts ...grpc.CallOption) (*List{{cookiecutter.service}}Response, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Create{{cookiecutter.service}}(ctx context.Context, in *Create{{cookiecutter.service}}Request, opts ...grpc.CallOption) (*{{cookiecutter.service}}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Get{{cookiecutter.service}}(ctx context.Context, in *Get{{cookiecutter.service}}Request, opts ...grpc.CallOption) (*{{cookiecutter.service}}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Update{{cookiecutter.service}}(ctx context.Context, in *Update{{cookiecutter.service}}Request, opts ...grpc.CallOption) (*{{cookiecutter.service}}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Delete{{cookiecutter.service}}(ctx context.Context, in *Delete{{cookiecutter.service}}Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	return nil, fmt.Errorf("not implemented")
}



// List{{cookiecutter.service}}s(ctx context.Context, in *List{{cookiecutter.service}}sRequest, opts ...grpc.CallOption) (*List{{cookiecutter.service}}sResponse, error)
//	Get{{cookiecutter.service}}(ctx context.Context, in *Get{{cookiecutter.service}}Request, opts ...grpc.CallOption) (*{{cookiecutter.service}}, error)
//	Create{{cookiecutter.service}}(ctx context.Context, in *Create{{cookiecutter.service}}Request, opts ...grpc.CallOption) (*{{cookiecutter.service}}, error)
//	Update{{cookiecutter.service}}(ctx context.Context, in *Update{{cookiecutter.service}}Request, opts ...grpc.CallOption) (*{{cookiecutter.service}}, error)
//	Delete{{cookiecutter.service}}(ctx context.Context, in *Delete{{cookiecutter.service}}Request, opts ...grpc.CallOption) (*empty.Empty, error)