package main

import (
	"backend/db"
	userv1 "backend/gen/user/v1"
	"backend/gen/user/v1/userv1connect"
	"backend/interceptor"
	"context"
	"database/sql"
	"errors"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	ServiceName    = "UserService-Version"
	ServiceVersion = "v1"
)

type UserServer struct {
	app  *firebase.App
	auth *auth.Client
	db   *sql.DB
}

func (s *UserServer) GetAccount(_ context.Context, req *connect.Request[userv1.GetAccountRequest]) (*connect.Response[userv1.GetAccountResponse], error) {
	// 処理要求アカウントのUserIdを取得
	requesterUserId := req.Header().Get("X-Peg-UserId")

	// DBからアカウント情報取り出し
	account_, err := db.GetAccount(s.db, requesterUserId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			// 一致するものがいない
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("account not found: %s", requesterUserId))
		default:
			// 不明なエラー
			// todo : log via discord
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error"))
		}
	}

	// レスポンス作成
	res := connect.NewResponse(&userv1.GetAccountResponse{Account: account_})
	res.Header().Set(ServiceName, ServiceVersion)

	return res, nil
}

func (s *UserServer) UpdateAccount(_ context.Context, _ *connect.Request[userv1.UpdateAccountRequest]) (*connect.Response[userv1.UpdateAccountResponse], error) {
	// emailを変更可能にするかどうか。
	// user_nameを別の関数として実装する予定なので、消しても良いかもしれない。
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("UpdateAccount inimplemented"))
}

func (s *UserServer) GetProfile(_ context.Context, req *connect.Request[userv1.GetProfileRequest]) (*connect.Response[userv1.GetProfileResponse], error) {
	// 対象アカウントのUserIdを取得
	targetUserId := req.Msg.UserId
	// DBからProfile情報の取得
	prof, err := db.GetProfile(s.db, targetUserId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			// 一致するものがいない
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("profile not found: %s", targetUserId))
		default:
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error"))
		}
	}

	res := connect.NewResponse(&userv1.GetProfileResponse{Profile: prof})
	res.Header().Set(ServiceName, ServiceVersion)

	return res, nil
}

func (s *UserServer) UpdateProfile(_ context.Context, req *connect.Request[userv1.UpdateProfileRequest]) (*connect.Response[userv1.UpdateProfileResponse], error) {
	requesterUserId := req.Header().Get("X-Peg-UserId")

	// optional, nili以外をアップデートします。
	newProfile, err := db.UpdateProfile(s.db, requesterUserId, req.Msg)
	if err != nil {
		log.Printf("%v", err)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("profile not found: %s", requesterUserId))
		case errors.Is(err, err.(*mysql.MySQLError)):
			mysqlErr := err.(*mysql.MySQLError)
			switch mysqlErr.Number {
			case 4025:
				return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("database error: CONSTRAINT"))
			default:
				return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("database error: %v", mysqlErr.Number))
			}
		case errors.Is(err, err.(*connect.Error)):
			return nil, err
		default:
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error"))
		}
	}

	res := connect.NewResponse(&userv1.UpdateProfileResponse{Profile: newProfile})
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

	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MARIADB_USER"),
		os.Getenv("MARIADB_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("MARIADB_DATABASE")))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	userServer := &UserServer{
		app:  app,
		auth: client,
		db:   database,
	}
	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(
		interceptor.NewFirebaseAuthInterceptor(client))
	mux.Handle(userv1connect.NewUserServiceHandler(
		userServer,
		interceptors,
	))

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("USER_SERVICE_PORT"))

	if err = http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
