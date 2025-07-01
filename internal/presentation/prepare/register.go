package prepare

import (
	"context"
	"query-service/internal/presentation/interceptor"

	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/grpc"
)

type QueryServer struct {
	Server *grpc.Server
}

func chainUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	newHandler := func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
		return interceptor.UUIDValidationInterceptor(currentCtx, currentReq, info, handler)
	}
	return interceptor.LoggingInterceptor(ctx, req, info, newHandler)
}

func NewQueryServer(category pb.CategoryQueryServer, product pb.ProductQueryServer) *QueryServer {
	serverOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(chainUnaryInterceptor),
	}
	server := grpc.NewServer(serverOpts...)
	pb.RegisterCategoryQueryServer(server, category)
	pb.RegisterProductQueryServer(server, product)
	return &QueryServer{
		Server: server,
	}
}
