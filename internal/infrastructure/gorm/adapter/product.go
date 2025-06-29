package adapter

import (
	"query-service/internal/domain/models/categories"
	"query-service/internal/domain/models/products"
	"query-service/internal/errs"
	"query-service/internal/infrastructure/gorm/models"
)

type productAdapterImpl struct{}

func NewProductAdapterImpl() products.ProductAdapter {
	return &productAdapterImpl{}
}

func (a *productAdapterImpl) Convert(source *products.Product) any {
	category := source.Category()
	return &models.Product{
		ObjId:        source.Id(),
		Name:         source.Name(),
		Price:        source.Price(),
		CategoryId:   (&category).Id(),
		CategoryName: (&category).Name(),
	}
}

func (a *productAdapterImpl) Rebuild(source any) (dest *products.Product, err error) {
	if p, ok := source.(*models.Product); ok {
		c := categories.NewCategory(p.CategoryId, p.CategoryName)
		dest = products.NewProduct(p.ObjId, p.Name, p.Price, *c)
	} else {
		err = errs.NewInternalError("product adapter rebuild error: source is not *models.Product")
	}
	return
}
