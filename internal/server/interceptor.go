package server

import (
	"time"
	"context"

	"google.golang.org/grpc"
)

func timerInterceptor(lgr Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()
		defer func() {
			lgr.Logf("time in %s: %s\n", info.FullMethod, time.Since(start))
		}()

		return handler(ctx, req)
	}
}