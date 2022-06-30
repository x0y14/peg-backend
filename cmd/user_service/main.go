package main

import (
	typesv1 "backend/gen/types/v1"
	userv1 "backend/gen/user/v1"
	"backend/gen/user/v1/userv1connect"
	"backend/interceptor"
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
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

type UserServer struct {
	app  *firebase.App
	auth *auth.Client
}

func (s *UserServer) Ping(ctx context.Context, req *connect.Request[userv1.PingRequest]) (*connect.Response[userv1.PingResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&userv1.PingResponse{Text: fmt.Sprintf("re: %s", req.Msg.Text)})
	// sample に　あったからつけてみてる。いる?
	res.Header().Set(ServiceName, ServiceVersion)
	return res, nil
}

func (s *UserServer) GetAccount(ctx context.Context, req *connect.Request[userv1.GetAccountRequest]) (*connect.Response[userv1.GetAccountResponse], error) {
	log.Println("Request headers: ", req.Header())

	userId := req.Header().Get("X-User-Id")
	// Firebase Authを使ってユーザーデータを取り出す
	userRecode, err := s.auth.GetUser(context.Background(), userId)
	if err != nil {
		// トークン認証は通ってるので、ものすごい短い間にアカウントを消したか、あるいは。
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&userv1.GetAccountResponse{Account: &typesv1.Account{
		UserId:   userRecode.UID,
		Email:    userRecode.Email,
		UserName: "", // todo : database
	}})
	res.Header().Set(ServiceName, ServiceVersion)

	return res, nil
}

func (s *UserServer) UpdateAccount(ctx context.Context, req *connect.Request[userv1.UpdateAccountRequest]) (*connect.Response[userv1.UpdateAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("UpdateAccount inimplemented"))
}

func (s *UserServer) GetProfile(ctx context.Context, req *connect.Request[userv1.GetProfileRequest]) (*connect.Response[userv1.GetProfileResponse], error) {
	log.Println("Request headers: ", req.Header())

	userId := req.Header().Get("X-User-Id")

	res := connect.NewResponse(&userv1.GetProfileResponse{Profile: &typesv1.Profile{
		UserId:        userId,
		DisplayName:   "", // todo : database
		IconPath:      "",
		StatusMessage: "",
		Metadata:      "",
	}})
	res.Header().Set(ServiceName, ServiceVersion)
	return res, nil
}

func (s *UserServer) UpdateProfile(ctx context.Context, req *connect.Request[userv1.UpdateProfileRequest]) (*connect.Response[userv1.UpdateProfileResponse], error) {
	log.Println("Request headers: ", req.Header())

	//userId := req.Header().Get("X-User-Id")

	//optional string display_name = 2;
	//optional string icon_path = 3;
	//optional string status_message = 4;
	//optional string metadata = 5;

	//s.auth.UpdateUser(context.Background(), userId)
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("UpdateProfile inimplemented"))
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

	userServer := &UserServer{
		app:  app,
		auth: client,
	}
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
