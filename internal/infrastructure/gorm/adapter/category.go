package adapter

import (
	"query-service/internal/domain/models/categories"
	"query-service/internal/errs"
	"query-service/internal/infrastructure/gorm/models"
)

type categoryAdapterImpl struct{}

func NewCategoryAdapterImpl() categories.CategoryAdapter {
	return &categoryAdapterImpl{}
}

func (a *categoryAdapterImpl) Convert(source *categories.Category) any {
	return &models.Category{
		ObjId: source.Id(),
		Name:  source.Name(),
	}
}

func (a *categoryAdapterImpl) Rebuild(source any) (dest *categories.Category, err error) {
	if c, ok := source.(*models.Category); ok {
		dest = categories.NewCategory(c.ObjId, c.Name)
	} else {
		err = errs.NewInternalError("category adapter rebuild error: source is not *models.Category")
	}
	return
}
