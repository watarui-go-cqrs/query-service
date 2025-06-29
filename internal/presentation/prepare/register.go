package prepare

import (
	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/grpc"
)

type QueryServer struct {
	Server *grpc.Server
}

func NewQueryServer(category pb.CategoryQueryServer, product pb.ProductQueryServer) *QueryServer {
	server := grpc.NewServer()
	pb.RegisterCategoryQueryServer(server, category)
	pb.RegisterProductQueryServer(server, product)
	return &QueryServer{
		Server: server,
	}
}
