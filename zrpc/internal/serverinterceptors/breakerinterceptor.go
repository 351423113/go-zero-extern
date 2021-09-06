package serverinterceptors

import (
	"context"

	"github.com/lukebull/go-zero-extern/core/breaker"
	"github.com/lukebull/go-zero-extern/zrpc/internal/codes"
	"google.golang.org/grpc"
)

// StreamBreakerInterceptor is an interceptor that acts as a circuit breaker.
func StreamBreakerInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) (err error) {
	breakerName := info.FullMethod
	return breaker.DoWithAcceptable(breakerName, func() error {
		return handler(srv, stream)
	}, codes.Acceptable)
}

// UnaryBreakerInterceptor is an interceptor that acts as a circuit breaker.
func UnaryBreakerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		breakerName := info.FullMethod
		err = breaker.DoWithAcceptable(breakerName, func() error {
			var err error
			resp, err = handler(ctx, req)
			return err
		}, codes.Acceptable)

		return resp, err
	}
}
