package products

import (
	"context"
)

type ProductRepository interface {
	List(ctx context.Context) ([]*Product, error)
	FindByProductId(ctx context.Context, productId string) (*Product, error)
	FindByProductNameLike(ctx context.Context, keyword string) ([]*Product, error)
}
