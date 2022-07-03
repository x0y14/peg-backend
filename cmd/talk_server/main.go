package main

import (
	"backend/db"
	talkv1 "backend/gen/talk/v1"
	"backend/gen/talk/v1/talkv1connect"
	"backend/interceptor"
	"backend/util"
	"context"
	"database/sql"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/bufbuild/connect-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/xid"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	HeaderKey   = "TalkService-Version"
	HeaderValue = "talkv1"
)

type TalkServer struct {
	app  *firebase.App
	auth *auth.Client
	db   *sql.DB
}

func (s *TalkServer) SendMessage(ctx context.Context, req *connect.Request[talkv1.SendMessageRequest]) (*connect.Response[talkv1.SendMessageResponse], error) {
	senderUserId := req.Header().Get("X-Peg-UserId")

	msg := req.Msg.Message
	msgId := fmt.Sprintf("ms|%s", xid.New().String())
	msg.Id = msgId
	msg.From = senderUserId // 強制付け替え

	if !util.IsDbJSON(msg.Metadata) {

	}

	// 送信先のチェック
	switch {
	case strings.Contains(msg.To, "gr|"):
		// groupか?
		return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("group is not implemented"))
	case strings.Contains(msg.To, "di|") && strings.Contains(msg.To, ".") && strings.Contains(msg.To, senderUserId):
		receiverUserId := strings.Replace(msg.To, "di|", "", 1)
		receiverUserId = strings.Replace(receiverUserId, ".", "", 1)
		receiverUserId = strings.Replace(receiverUserId, senderUserId, "", 1)

		// 自分宛ではないか
		if receiverUserId == senderUserId {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid receiver"))
		}

		// 存在確認
		if !db.IsAccountExists(s.db, receiverUserId) {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid receiver"))
		}

		// 必ずしも適正なidとは限らないので、再構築
		msg.To = util.CreateDirectChatId(msg.From, receiverUserId)

	default:
		// 自分宛ではないか
		if msg.To == senderUserId {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid receiver"))
		}

		// 存在確認
		if !db.IsAccountExists(s.db, msg.To) {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid receiver"))
		}

		msg.To = util.CreateDirectChatId(msg.From, msg.To)
	}

	resMsg, err := db.CreateMessage(s.db, msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	// メッセージの保存は完了.

	res := connect.NewResponse(&talkv1.SendMessageResponse{Message: resMsg})

	return res, nil
}

func (s *TalkServer) SendReadReceipt(ctx context.Context, req *connect.Request[talkv1.SendReadReceiptRequest]) (*connect.Response[talkv1.SendReadReceiptResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("sorry"))
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

	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		"root",
		os.Getenv("MARIADB_ROOT_PASSWORD"),
		os.Getenv("MARIADB_DATABASE")))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	talkServer := &TalkServer{
		app:  app,
		auth: client,
		db:   database,
	}

	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(
		interceptor.NewFirebaseAuthInterceptor(client))
	mux.Handle(talkv1connect.NewTalkServiceHandler(
		talkServer,
		interceptors))

	if err = http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
