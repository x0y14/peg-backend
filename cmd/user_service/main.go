package main

import (
	typesv1 "backend/gen/types/v1"
	userv1 "backend/gen/user/v1"
	"backend/gen/user/v1/userv1connect"
	"backend/interceptor"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

const (
	ServiceName    = "UserService-Version"
	ServiceVersion = "v1"
)

type UserServer struct{}

func (s *UserServer) Ping(ctx context.Context, req *connect.Request[userv1.PingRequest]) (*connect.Response[userv1.PingResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&userv1.PingResponse{Text: fmt.Sprintf("re: %s", req.Msg.Text)})
	// sample に　あったからつけてみてる。いる?
	res.Header().Set(ServiceName, ServiceVersion)
	return res, nil
}

func (s *UserServer) GetAccount(ctx context.Context, req *connect.Request[userv1.GetAccountRequest]) (*connect.Response[userv1.GetAccountResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&userv1.GetAccountResponse{Account: &typesv1.Account{
		UserId:   "",
		Email:    "",
		Username: "",
	}})
	res.Header().Set(ServiceName, ServiceVersion)
	return res, nil
}

func (s *UserServer) GetProfile(ctx context.Context, req *connect.Request[userv1.GetProfileRequest]) (*connect.Response[userv1.GetProfileResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&userv1.GetProfileResponse{Profile: &typesv1.Profile{
		UserId:      "",
		DisplayName: "",
		IconUrl:     "",
		Metadata:    "",
	}})
	res.Header().Set(ServiceName, ServiceVersion)
	return res, nil
}

func main() {
	// 環境変数読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading the .env file: %v", err)
	}
	// Firebase appの初期化
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// Firebase Authの初期化
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	userServer := &UserServer{}
	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(
		interceptor.NewFirebaseAuthInterceptor(client))
	mux.Handle(userv1connect.NewUserServiceHandler(
		userServer,
		interceptors,
	))
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
