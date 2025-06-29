package builder

import (
	"query-service/internal/domain/models/categories"
	"query-service/internal/domain/models/products"
	"query-service/internal/errs"

	"github.com/watarui-go-cqrs/pb/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type resultBuilderImpl struct{}

func NewResultBuilderImpl() ResultBuilder {
	return &resultBuilderImpl{}
}

func (b *resultBuilderImpl) BuildCategoryResult(source any) *pb.CategoryResult {
	result := &pb.CategoryResult{Timestamp: timestamppb.Now()}
	if category, ok := source.(*categories.Category); ok {
		result.Result = &pb.CategoryResult_Category{Category: &pb.Category{Id: category.Id(), Name: category.Name()}}
	} else {
		result.Result = &pb.CategoryResult_Error{Error: b.BuildErrorResult(source)}
	}
	return result
}

func (b *resultBuilderImpl) BuildCategoriesResult(source any) *pb.CategoriesResult {
	result := &pb.CategoriesResult{Timestamp: timestamppb.Now()}
	if categories, ok := source.([]*categories.Category); ok {
		for _, category := range categories {
			result.Categories = append(result.Categories, &pb.Category{Id: category.Id(), Name: category.Name()})
		}
	} else {
		result.Error = b.BuildErrorResult(source)
	}
	return result
}

func (b *resultBuilderImpl) BuildProductResult(source any) *pb.ProductResult {
	result := &pb.ProductResult{Timestamp: timestamppb.Now()}
	if product, ok := source.(*products.Product); ok {
		result.Result = &pb.ProductResult_Product{Product: &pb.Product{
			Id:    product.Id(),
			Name:  product.Name(),
			Price: int32(product.Price()),
			Category: &pb.Category{
				Id:   product.Id(),
				Name: product.Name(),
			},
		}}
	} else {
		result.Result = &pb.ProductResult_Error{Error: b.BuildErrorResult(source)}
	}
	return result
}

func (b *resultBuilderImpl) BuildProductsResult(source any) *pb.ProductsResult {
	result := &pb.ProductsResult{Timestamp: timestamppb.Now()}
	if products, ok := source.([]*products.Product); ok {
		for _, product := range products {
			result.Products = append(result.Products, &pb.Product{
				Id:    product.Id(),
				Name:  product.Name(),
				Price: int32(product.Price()),
				Category: &pb.Category{
					Id:   product.Id(),
					Name: product.Name(),
				},
			})
		}
	} else {
		result.Error = b.BuildErrorResult(source)
	}
	return result
}

func (b *resultBuilderImpl) BuildErrorResult(source any) *pb.Error {
	switch v := source.(type) {
	case *errs.CRUDError:
		return &pb.Error{Type: "CRUD Error", Message: v.Error()}
	case *errs.InternalError:
		return &pb.Error{Type: "Internal Error", Message: v.Error()}
	default:
		return &pb.Error{Type: "Unknown Error", Message: "unknown error"}
	}
}
