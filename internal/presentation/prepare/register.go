package prepare

import (
	"query-service/internal/presentation/interceptor"

	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/grpc"
)

type QueryServer struct {
	Server *grpc.Server
}

func NewQueryServer(category pb.CategoryQueryServer, product pb.ProductQueryServer) *QueryServer {
	serverOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.LoggingInterceptor),
	}
	server := grpc.NewServer(serverOpts...)
	pb.RegisterCategoryQueryServer(server, category)
	pb.RegisterProductQueryServer(server, product)
	return &QueryServer{
		Server: server,
	}
}
