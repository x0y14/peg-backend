package interceptor

import (
	"context"
	"github.com/bufbuild/connect-go"
)

func NewHeaderSetInterceptor(serviceName string, version string) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			res, err := next(ctx, request)
			if err != nil {
				return nil, err
			}
			res.Header().Set(serviceName, version)
			return res, nil
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
