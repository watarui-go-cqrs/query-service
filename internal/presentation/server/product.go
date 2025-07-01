package server

import (
	"context"
	"query-service/internal/domain/models/products"
	"query-service/internal/presentation/builder"

	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type productServer struct {
	repository products.ProductRepository
	builder    builder.ResultBuilder
	pb.UnimplementedProductQueryServer
}

func NewProductServer(repository products.ProductRepository, builder builder.ResultBuilder) pb.ProductQueryServer {
	return &productServer{
		repository: repository,
		builder:    builder,
	}
}

func (s *productServer) ListStream(param *emptypb.Empty, stream pb.ProductQuery_ListStreamServer) error {
	results, err := s.repository.List(context.Background())
	if err != nil {
		return err
	}
	products := s.builder.BuildProductsResult(results)
	for _, product := range products.GetProducts() {
		if err := stream.Send(product); err != nil {
			return err
		}
	}
	return nil
}

func (s *productServer) List(ctx context.Context, param *emptypb.Empty) (*pb.ProductsResult, error) {
	products, err := s.repository.List(ctx)
	if err != nil {
		return s.builder.BuildProductsResult(err), nil
	}
	return s.builder.BuildProductsResult(products), nil
}

func (s *productServer) ById(ctx context.Context, param *pb.ProductParam) (*pb.ProductResult, error) {
	product, err := s.repository.FindByProductId(ctx, param.GetId())
	if err != nil {
		return s.builder.BuildProductResult(err), nil
	}
	return s.builder.BuildProductResult(product), nil
}

func (s *productServer) ByKeyword(ctx context.Context, param *pb.ProductParam) (*pb.ProductsResult, error) {
	products, err := s.repository.FindByProductNameLike(ctx, param.GetKeyword())
	if err != nil {
		return s.builder.BuildProductsResult(err), nil
	}
	return s.builder.BuildProductsResult(products), nil
}
