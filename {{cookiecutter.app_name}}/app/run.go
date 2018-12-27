package app

import (
	"github.com/gofunct/common/runtime"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/app/server"
	"golang.org/x/tools/go/internal/gccgoimporter/testdata"
)

// Run starts the runtime server.
func Run() error {
	r := runtime.New(
		runtime.WithDefaultLogger(),
		runtime.WithServers(
			server.New{{cookiecutter.service}}ServiceServer()
		),
	)
	return r.Serve()
}
