package builder

import "github.com/watarui-go-cqrs/pb/pb"

type ResultBuilder interface {
	BuildCategoryResult(source any) *pb.CategoryResult
	BuildCategoriesResult(source any) *pb.CategoriesResult
	BuildProductResult(source any) *pb.ProductResult
	BuildProductsResult(source any) *pb.ProductsResult
	BuildErrorResult(source any) *pb.Error
}
