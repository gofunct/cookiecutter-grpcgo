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

	session_svc "github.com/gofunct/example/services"
	session_endpoints "github.com/gofunct/example/services/session/gen/endpoints"
	session_pb "github.com/gofunct/example/services/session/gen/pb"
	session_grpctransport "github.com/gofunct/example/services/session/gen/transports/grpc"
	session_httptransport "github.com/gofunct/example/services/session/gen/transports/http"

	sprint_svc "github.com/gofunct/example/services/sprint"
	sprint_endpoints "github.com/gofunct/example/services/sprint/gen/endpoints"
	sprint_pb "github.com/gofunct/example/services/sprint/gen/pb"
	sprint_grpctransport "github.com/gofunct/example/services/sprint/gen/transports/grpc"
	sprint_httptransport "github.com/gofunct/example/services/sprint/gen/transports/http"

	user_svc "github.com/gofunct/example/services/user"
	user_endpoints "github.com/gofunct/example/services/user/gen/endpoints"
	user_pb "github.com/gofunct/example/services/user/gen/pb"
	user_grpctransport "github.com/gofunct/example/services/user/gen/transports/grpc"
	user_httptransport "github.com/gofunct/example/services/user/gen/transports/http"
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

	// initialize services
	{
		svc := session_svc.New()
		endpoints := session_endpoints.MakeEndpoints(svc)
		srv := session_grpctransport.MakeGRPCServer(endpoints)
		session_pb.RegisterSessionServiceServer(s, srv)
		session_httptransport.RegisterHandlers(svc, mux, endpoints)
	}
	{
		svc := sprint_svc.New()
		endpoints := sprint_endpoints.MakeEndpoints(svc)
		srv := sprint_grpctransport.MakeGRPCServer(endpoints)
		sprint_pb.RegisterSprintServiceServer(s, srv)
		sprint_httptransport.RegisterHandlers(svc, mux, endpoints)
	}
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
		logger := log.With(logger, "transport", "HTTP")
		logger.Log("addr", ":"+viper.GetString("http_port"))
		errc <- http.ListenAndServe(viper.GetString(":"+"http_port"), handlers.LoggingHandler(os.Stderr, mux))
	}()

	go func() {
		logger := log.With(logger, "transport", "gRPC")
		ln, err := net.Listen("tcp", defaultConfig.)
		if err != nil {
			errc <- err
			return
		}
		logger.Log("addr", ":"+viper.GetString("grpc_port"))
		errc <- s.Serve(ln)
	}()

	logger.Log("exit", <-errc)
}
