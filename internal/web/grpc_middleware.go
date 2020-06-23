package web

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// NewGRPCUnaryServerInterceptor ...
func NewGRPCUnaryServerInterceptor(nextTraceID func() string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (
		resp interface{},
		err error,
	) {
		// 尝试获取traceID
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			tid := md.Get(string(traceIDKey))
			if len(tid) == 1 && tid[0] != "" {
				log.Printf("current call stack: %s", tid[0])
				return handler(ctx, req)
			}
		}

		// 按需创建新traceID
		tid := nextTraceID()
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(string(traceIDKey), tid))
		log.Printf("new call stack: %s", tid)

		return handler(ctx, req)
	}

}
