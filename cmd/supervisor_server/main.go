package main

import (
	supervisorv1 "backend/gen/supervisor/v1"
	"backend/gen/supervisor/v1/supervisorv1connect"
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type SupervisorServer struct{}

func (s *SupervisorServer) CreateAccount(_ context.Context,
	req *connect.Request[supervisorv1.CreateAccountRequest]) (*connect.Response[supervisorv1.CreateAccountResponse], error) {
	return nil, nil
}
func (s *SupervisorServer) CreateProfile(_ context.Context,
	req *connect.Request[supervisorv1.CreateProfileRequest]) (*connect.Response[supervisorv1.CreateProfileResponse], error) {
	return nil, nil
}
func (s *SupervisorServer) RecordOperation(_ context.Context,
	req *connect.Request[supervisorv1.RecordOperationRequest]) (*connect.Response[supervisorv1.RecordOperationResponse], error) {
	return nil, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading the .env file: %v", err)
	}

	sv := &SupervisorServer{}
	mux := http.NewServeMux()
	path, handler := supervisorv1connect.NewSupervisorServiceHandler(sv)
	mux.Handle(path, handler)

	if err := http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
