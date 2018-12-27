package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/api"
)

func Test_{{cookiecutter.service}}ServiceServer_List{{cookiecutter.service}}s(t *testing.T) {
	svr := New{{cookiecutter.service}}ServiceServer()

	ctx := context.Background()
	req := &api_pb.List{{cookiecutter.service}}sRequest{}

	resp, err := svr.List{{cookiecutter.service}}s(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_{{cookiecutter.service}}ServiceServer_Get{{cookiecutter.service}}(t *testing.T) {
	svr := New{{cookiecutter.service}}ServiceServer()

	ctx := context.Background()
	req := &api_pb.Get{{cookiecutter.service}}Request{}

	resp, err := svr.Get{{cookiecutter.service}}(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_{{cookiecutter.service}}ServiceServer_Create{{cookiecutter.service}}(t *testing.T) {
	svr := New{{cookiecutter.service}}ServiceServer()

	ctx := context.Background()
	req := &api_pb.Create{{cookiecutter.service}}Request{}

	resp, err := svr.Create{{cookiecutter.service}}(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_{{cookiecutter.service}}ServiceServer_Update{{cookiecutter.service}}(t *testing.T) {
	svr := New{{cookiecutter.service}}ServiceServer()

	ctx := context.Background()
	req := &api_pb.Update{{cookiecutter.service}}Request{}

	resp, err := svr.Update{{cookiecutter.service}}(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_{{cookiecutter.service}}ServiceServer_Delete{{cookiecutter.service}}(t *testing.T) {
	svr := New{{cookiecutter.service}}ServiceServer()

	ctx := context.Background()
	req := &api_pb.Delete{{cookiecutter.service}}Request{}

	resp, err := svr.Delete{{cookiecutter.service}}(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
