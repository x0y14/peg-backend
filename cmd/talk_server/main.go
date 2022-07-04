package main

import (
	"backend/db"
	supervisorv1 "backend/gen/supervisor/v1"
	"backend/gen/supervisor/v1/supervisorv1connect"
	talkv1 "backend/gen/talk/v1"
	"backend/gen/talk/v1/talkv1connect"
	typesv1 "backend/gen/types/v1"
	"backend/interceptor"
	"backend/scripts"
	"backend/util"
	"context"
	"database/sql"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/bufbuild/connect-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/xid"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	app              *firebase.App
	auth             *auth.Client
	db               *sql.DB
	supervisorClient *supervisorv1connect.SupervisorServiceClient
}

func (s *TalkServer) SendMessage(ctx context.Context, req *connect.Request[talkv1.SendMessageRequest]) (*connect.Response[talkv1.SendMessageResponse], error) {
	senderUserId := req.Header().Get("X-Peg-UserId")

	msg := req.Msg.Message
	msgId := fmt.Sprintf("ms|%s", xid.New().String())
	msg.Id = msgId
	msg.From = senderUserId // 強制付け替え

	// opを受け取るユーザーの生のuser_idが入る
	// 自分が送信したことは通知されるべき     複数端末ログインを実装してるから?
	sendOpDest := []string{msg.From}
	// 自分が送信したものを受信できるべき,,,? 参考にしたものはこれも自身で受け取ってる.
	recvOpDest := []string{msg.From}

	// 送信先のチェック
	switch {
	case strings.Contains(msg.To, "gr|"):
		// groupか?
		//opDestination
		return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("group is not implemented"))
	case strings.Contains(msg.To, "di|") && strings.Contains(msg.To, ".") && strings.Contains(msg.To, senderUserId):
		// 余計な部分を一回排除
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

		// opを相手が受け取れるように
		recvOpDest = append(recvOpDest, receiverUserId)

		// direct-chat id再構築
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

		// opを相手が受け取れるように
		// msg.toはこの時点で相手のuser_idが入ってる
		recvOpDest = append(recvOpDest, msg.To)

		// direct-chat idに変更
		msg.To = util.CreateDirectChatId(msg.From, msg.To)
	}

	// 実際にdbに挿入したものを返してあげる
	resMsg, err := db.CreateMessage(s.db, msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	// op配布
	// 管理者トークン発行
	adminToken, err := scripts.GenerateAdminToken()
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	// request作成
	recordReq := connect.NewRequest(&supervisorv1.RecordOperationRequest{Operations: []*typesv1.Operation{
		{
			Id:          0,
			Type:        typesv1.OperationType_OPERATION_TYPE_SEND_MESSAGE,
			Source:      msg.From,
			Destination: sendOpDest,
			Param1:      &msg.Id,
			Param2:      &msg.From,
			Param3:      &msg.To,
			CratedAt:    timestamppb.Now(),
		},
		{
			Id:          0,
			Type:        typesv1.OperationType_OPERATION_TYPE_SEND_MESSAGE_RECV,
			Source:      msg.From,
			Destination: recvOpDest,
			Param1:      &msg.Id,
			Param2:      &msg.From,
			Param3:      &msg.To,
			CratedAt:    timestamppb.Now(),
		},
	}})
	// トークンをくっつけてあげる
	recordReq.Header().Set("Authorization", fmt.Sprintf("Bearer %s", adminToken))
	// リクエスト送信
	go func() {
		_, err = (*s.supervisorClient).RecordOperation(
			context.Background(),
			recordReq,
		)
		if err != nil {
			log.Println(err)
		}
	}()
	// レスポンス作成
	res := connect.NewResponse(&talkv1.SendMessageResponse{Message: resMsg})

	return res, nil
}

func (s *TalkServer) SendReadReceipt(ctx context.Context, req *connect.Request[talkv1.SendReadReceiptRequest]) (*connect.Response[talkv1.SendReadReceiptResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("sorry"))
}

func main() {
	// 環境変数読み込み
	//if err := godotenv.Load(); err != nil {
	//	log.Fatalf("error loading the .env file: %v", err)
	//}

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

	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
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

	// supervisor client
	supervisorClient := supervisorv1connect.NewSupervisorServiceClient(
		http.DefaultClient,
		fmt.Sprintf("http://%s:%s", os.Getenv("SUPERVISOR_SERVICE_HOST"), os.Getenv("SUPERVISOR_SERVICE_PORT")),
	)

	talkServer := &TalkServer{
		app:              app,
		auth:             client,
		db:               database,
		supervisorClient: &supervisorClient,
	}

	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(
		interceptor.NewFirebaseAuthInterceptor(client))
	mux.Handle(talkv1connect.NewTalkServiceHandler(
		talkServer,
		interceptors))

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("TALK_SERVICE_PORT"))
	if err = http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
