package interceptor

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/bufbuild/connect-go"
	"strings"
)

func NewFirebaseAuthInterceptor(cl *auth.Client) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			// インターセプター（前処理）
			// "Bearer e..."
			idTokenRaw := request.Header().Get("Authorization")
			if idTokenRaw == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("need Bearer token"))
			}
			// "e..."
			idToken := strings.Replace(idTokenRaw, "Bearer ", "", 1)
			// 検証
			token, err := cl.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			// 本来の処理で使用するのでデータをくっつけてあげる。
			request.Header().Set("X-User-ID", token.UID)

			// 本来の処理
			res, err := next(ctx, request)
			if err != nil {
				return nil, err
			}
			return res, nil
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
