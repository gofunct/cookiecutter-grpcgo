package cmd

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/app"
	"google.golang.org/grpc/grpclog"
	"os"
	"github.com/spf13/cobra"
)

func init() {
}

// gatewayCmd represents the gateway command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the grpc/http runtime server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := app.Run(); err != nil {
			grpclog.Errorf("server was shutdown with errors: %v", err)
			os.Exit(1)
		}
	},
}

