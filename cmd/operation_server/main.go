package main

import (
	"backend/db"
	operationv1 "backend/gen/operation/v1"
	"backend/gen/operation/v1/operationv1connect"
	"backend/gen/supervisor/v1/supervisorv1connect"
	typesv1 "backend/gen/types/v1"
	"backend/interceptor"
	"context"
	"database/sql"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/bufbuild/connect-go"
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

type OperationServer struct {
	app              *firebase.App
	auth             *auth.Client
	db               *sql.DB
	supervisorClient *supervisorv1connect.SupervisorServiceClient
	pulsarClient     *pulsar.Client
}

func (s *OperationServer) FetchOperations(ctx context.Context, req *connect.Request[operationv1.FetchOperationsRequest], stream *connect.ServerStream[operationv1.FetchOperationsResponse]) error {
	requesterUserId := req.Header().Get("X-Peg-UserId")

	log.Printf("X-Peg-UserId: %s", requesterUserId)

	// pulsarと接続
	consumer, err := (*s.pulsarClient).Subscribe(pulsar.ConsumerOptions{
		Topic:            requesterUserId,
		SubscriptionName: "fetch-" + requesterUserId,
		Type:             pulsar.Shared,
	})
	if err != nil {
		return connect.NewError(connect.CodeUnknown, err)
	}
	defer consumer.Close()

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			return connect.NewError(connect.CodeUnknown, err)
		}
		operationId, err := strconv.ParseInt(string(msg.Payload()), 10, 64)
		log.Printf("Receive OpId From Pulsar: %v\n", operationId)
		consumer.Ack(msg)
		if err != nil {
			return connect.NewError(connect.CodeUnknown, err)
		}

		op, err := db.GetOperationWithOperationId(s.db, operationId)
		if err != nil {
			log.Printf("GetOperationWithOperationId: %v", err)
			return connect.NewError(connect.CodeUnknown, err)
		}

		var opMsg *typesv1.Message
		if op.Type == typesv1.OperationType_OPERATION_TYPE_SEND_MESSAGE || op.Type == typesv1.OperationType_OPERATION_TYPE_SEND_MESSAGE_RECV {
			m_, err := db.GetMessageWithMessageId(s.db, *op.Param1)
			if err != nil {
				log.Printf("GetMessageWithMessageId: %v", err)
				return connect.NewError(connect.CodeUnknown, err)
			}
			opMsg = m_
		}

		err = stream.Send(&operationv1.FetchOperationsResponse{
			Operation: op,
			Message:   opMsg,
		})
		if err != nil {
			log.Println(err)
			return connect.NewError(connect.CodeUnknown, err)
		}
	}
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

	// supervisor authClient
	supervisorClient := supervisorv1connect.NewSupervisorServiceClient(
		http.DefaultClient,
		fmt.Sprintf("http://%s:%s", os.Getenv("SUPERVISOR_SERVICE_HOST"), os.Getenv("SUPERVISOR_SERVICE_PORT")),
	)

	// pulsar
	pl, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               fmt.Sprintf("pulsar://%s:%s", os.Getenv("PULSAR_HOST"), os.Getenv("PULSAR_PORT")),
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar authClient: %v", err)
	}
	defer pl.Close()

	operationServer := &OperationServer{
		app:              app,
		auth:             authClient,
		db:               database,
		supervisorClient: &supervisorClient,
		pulsarClient:     &pl,
	}

	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(
		//interceptor.NewFirebaseAuthInterceptor(authClient),
		interceptor.NewAuthInterceptor(authClient),
	)
	mux.Handle(operationv1connect.NewOperationServiceHandler(
		operationServer,
		interceptors))

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("OPERATION_SERVICE_PORT"))
	if err = http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}

}
