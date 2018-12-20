package cmd


import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/handlers"
	"google.golang.org/grpc"

	user_pb "github.com/gofunct/example/gen/user"
	user_endpoints "github.com/gofunct/example/gen/user/endpoints"
	user_svc "github.com/gofunct/example/services/user"
	user_grpctransport "github.com/gofunct/example/gen/user/transports/grpc"
	user_httptransport "github.com/gofunct/example/gen/user/transports/http"
)


var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "start a grpc and http server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func runServer() {
	mux := http.NewServeMux()
	errc := make(chan error)
	s := grpc.NewServer()

	// initialize services here...
	{
		svc := user_svc.New()
		endpoints := user_endpoints.MakeEndpoints(svc)
		srv := user_grpctransport.MakeGRPCServer(endpoints)
		user_pb.RegisterUserServiceServer(s, srv)
		user_httptransport.RegisterHandlers(svc, mux, endpoints)
	}

	// start servers
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger := log.With(kitLog, "transport", "HTTP")
		logger.Log("addr", ":"+viper.GetString("http_port"))
		errc <- http.ListenAndServe(viper.GetString(":"+"http_port"), handlers.LoggingHandler(os.Stderr, mux))
	}()

	go func() {
		logger := log.With(kitLog, "transport", "gRPC")
		ln, err := net.Listen("tcp", defaultConfig.GetString("grpc_port"))
		if err != nil {
			errc <- err
			return
		}
		logger.Log("addr", ":"+defaultConfig.GetString("grpc_port"))
		errc <- s.Serve(ln)
	}()

	kitLog.Log("exit", <-errc)
}
