package main

import (
	supervisorv1 "backend/gen/supervisor/v1"
	"backend/gen/supervisor/v1/supervisorv1connect"
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type SupervisorServer struct{}

func (s *SupervisorServer) Ping(ctx context.Context, req *connect.Request[supervisorv1.PingRequest]) (*connect.Response[supervisorv1.PingResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&supervisorv1.PingResponse{Message: fmt.Sprintf("re: %s", req.Msg.Message)})
	res.Header().Set("Supervisor-Version", "v1")
	return res, nil
}

func (s *SupervisorServer) ServerStreamPing(ctx context.Context, req *connect.Request[supervisorv1.ServerStreamPingRequest], stream *connect.ServerStream[supervisorv1.ServerStreamPingResponse]) error {
	for i := 0; i < 5; i++ {
		if err := stream.Send(&supervisorv1.ServerStreamPingResponse{
			Message: fmt.Sprintf("re(%d): %s", i, req.Msg.Message),
		}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	//if err := godotenv.Load(); err != nil {
	//	log.Fatalf("error loading the .env file: %v", err)
	//}

	sv := &SupervisorServer{}
	mux := http.NewServeMux()
	path, handler := supervisorv1connect.NewSupervisorServiceHandler(sv)
	mux.Handle(path, handler)
	//mux.Handle(path, middleware.EnsureValidToken()(handler))
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
