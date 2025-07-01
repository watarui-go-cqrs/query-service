package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Interceptor Log: %s:%s", info.Server, info.FullMethod)

	resp, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
