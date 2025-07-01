package prepare

import (
	"context"
	"crypto/tls"
	"os"
	"query-service/internal/presentation/interceptor"

	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func createCreds() credentials.TransportCredentials {
	certFile := "internal/presentation/prepare/queryservice.pem"
	keyFile := "internal/presentation/prepare/queryservice-key.pem"

	cert, err := os.ReadFile(certFile)
	if err != nil {
		panic(err)
	}
	key, err := os.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}
	certificate, err := tls.X509KeyPair(cert, key)
	if err != nil {
		panic(err)
	}
	creds := credentials.NewServerTLSFromCert(&certificate)
	return creds
}

func NewQueryServer(category pb.CategoryQueryServer, product pb.ProductQueryServer) *QueryServer {
	serverOpts := []grpc.ServerOption{
		grpc.Creds(createCreds()),
		grpc.UnaryInterceptor(chainUnaryInterceptor),
	}
	server := grpc.NewServer(serverOpts...)
	pb.RegisterCategoryQueryServer(server, category)
	pb.RegisterProductQueryServer(server, product)
	return &QueryServer{
		Server: server,
	}
}
