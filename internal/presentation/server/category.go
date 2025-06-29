package server

import (
	"context"
	"query-service/internal/domain/models/categories"
	"query-service/internal/presentation/builder"

	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryServer struct {
	repository categories.CategoryRepository
	builder    builder.ResultBuilder
	pb.UnimplementedCategoryQueryServer
}

func NewCategoryServer(repository categories.CategoryRepository, builder builder.ResultBuilder) pb.CategoryQueryServer {
	return &categoryServer{
		repository: repository,
		builder:    builder,
	}
}

func (s *categoryServer) List(ctx context.Context, param *emptypb.Empty) (*pb.CategoriesResult, error) {
	categories, err := s.repository.List(ctx)
	if err != nil {
		return s.builder.BuildCategoriesResult(err), nil
	}
	return s.builder.BuildCategoriesResult(categories), nil
}

func (s *categoryServer) ById(ctx context.Context, param *pb.CategoryParam) (*pb.CategoryResult, error) {
	category, err := s.repository.FindByCategoryId(ctx, param.GetId())
	if err != nil {
		return s.builder.BuildCategoryResult(err), nil
	}
	return s.builder.BuildCategoryResult(category), nil
}
