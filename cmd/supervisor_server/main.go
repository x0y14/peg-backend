package main

import (
	"backend/db"
	supervisorv1 "backend/gen/supervisor/v1"
	"backend/gen/supervisor/v1/supervisorv1connect"
	"backend/interceptor"
	"backend/util"
	"context"
	"database/sql"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/bufbuild/connect-go"
	"github.com/bwmarrin/snowflake"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	defaultDisplayName = "unknown_v1"
	defaultIconPath    = "image_v1"
)

type SupervisorServer struct {
	db           *sql.DB
	auth         *auth.Client
	id           *snowflake.Node
	pulsarClient *pulsar.Client
}

func (s *SupervisorServer) CreateAccount(_ context.Context,
	req *connect.Request[supervisorv1.CreateAccountRequest]) (*connect.Response[supervisorv1.CreateAccountResponse], error) {
	// 権限確認
	if !util.ParseBool(req.Header().Get("X-Peg-Admin")) {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("admin only"))
	}

	// firebaseにユーザーが存在するかを確認
	user, err := s.auth.GetUser(context.Background(), req.Msg.Account.UserId)
	if err != nil {
		if auth.IsUserNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	// databaseに作成
	account, err := db.CreateAccount(s.db, user.UID, user.Email)
	if err != nil {
		// todo : already exists
		// Error 1062: Duplicate entry
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	res := connect.NewResponse(&supervisorv1.CreateAccountResponse{Account: account})

	return res, nil
}
func (s *SupervisorServer) CreateProfile(_ context.Context,
	req *connect.Request[supervisorv1.CreateProfileRequest]) (*connect.Response[supervisorv1.CreateProfileResponse], error) {
	// 権限確認
	if !util.ParseBool(req.Header().Get("X-Peg-Admin")) {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("admin only"))
	}

	// firebaseにユーザーが存在するか
	user, err := s.auth.GetUser(context.Background(), req.Msg.Profile.UserId)
	if err != nil {
		if auth.IsUserNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	// databaseに作成
	profile, err := db.CreateProfile(s.db, user.UID, defaultDisplayName, defaultIconPath)
	if err != nil {
		// todo : already exists
		// Error 1062: Duplicate entry
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	res := connect.NewResponse(&supervisorv1.CreateProfileResponse{Profile: profile})

	return res, nil
}
func (s *SupervisorServer) RecordOperation(_ context.Context,
	req *connect.Request[supervisorv1.RecordOperationRequest]) (*connect.Response[supervisorv1.RecordOperationResponse], error) {
	// 権限確認
	if !util.ParseBool(req.Header().Get("X-Peg-Admin")) {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("admin only"))
	}

	// supervisorは、お願いされたものを記録するだけ。
	for _, op := range req.Msg.Operations {
		opId := s.id.Generate().Int64()
		// op本体の記録
		_, err := db.CreateOperation(s.db, opId, op.Type, op.Source)
		if err != nil {
			return nil, connect.NewError(connect.CodeUnknown, err)
		}

		// destinationを記録する前にopを記録しているので、宛先が出鱈目でもopだけは保存されてしまう。
		// destinationのop_idはfkなので、先に入れることはできない。
		// 先にdestをチェックするか、、、

		// 宛先の記録
		err = db.CreateOperationDestination(s.db, opId, op.Destination)
		if err != nil {
			return nil, connect.NewError(connect.CodeUnknown, err)
		}

		for _, dest := range op.Destination {
			// pulsar
			producer, err := (*s.pulsarClient).CreateProducer(pulsar.ProducerOptions{
				Topic: dest,
			})
			if err != nil {
				return nil, connect.NewError(connect.CodeUnknown, err)
			}
			_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
				Payload: []byte(strconv.FormatInt(opId, 10)),
			})
			if err != nil {
				return nil, connect.NewError(connect.CodeUnknown, err)
			}
			producer.Close()
		}
	}

	res := connect.NewResponse(&supervisorv1.RecordOperationResponse{})

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
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth authClient: %v\n", err)
	}

	// database準備
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

	// id generator
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}

	pl, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               fmt.Sprintf("pulsar://%s:%s", os.Getenv("PULSAR_HOST"), os.Getenv("PULSAR_PORT")),
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar authClient: %v", err)
	}
	defer pl.Close()

	// 渡してあげる
	supervisorServer := &SupervisorServer{
		db:           database,
		auth:         authClient,
		id:           node,
		pulsarClient: &pl,
	}

	// インターセプター埋め込んであげる
	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(
		interceptor.NewFirebaseAuthInterceptor(authClient))
	mux.Handle(supervisorv1connect.NewSupervisorServiceHandler(
		supervisorServer,
		interceptors,
	))

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("SUPERVISOR_SERVICE_PORT"))

	// 起動
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
