package repository

import (
	"context"
	"fmt"
	"query-service/internal/domain/models/categories"
	"query-service/internal/errs"
	"query-service/internal/infrastructure/gorm/models"

	"gorm.io/gorm"
)

const (
	CATEGORY_TABLE   = "category"
	CATEGORY_COLUMNS = "id AS c_key,obj_id AS c_id,name AS c_name"
	CATEGORY_WHERE   = "obj_id = ?"
)

type categoryRepositoryGORM struct {
	db      *gorm.DB
	adapter categories.CategoryAdapter
}

func NewCategoryRepositoryGORM(db *gorm.DB, adapter categories.CategoryAdapter) categories.CategoryRepository {
	return &categoryRepositoryGORM{db: db, adapter: adapter}
}

func (c *categoryRepositoryGORM) List(ctx context.Context) ([]*categories.Category, error) {
	var models = []*models.Category{}
	if result := c.db.WithContext(ctx).
		Table(CATEGORY_TABLE).
		Select(CATEGORY_COLUMNS).
		Find(&models); result.Error != nil {
		return nil, errs.NewCRUDError(result.Error.Error())
	}

	var categories []*categories.Category
	for _, model := range models {
		category, err := c.adapter.Rebuild(model)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *categoryRepositoryGORM) FindByCategoryId(ctx context.Context, categoryId string) (*categories.Category, error) {
	var model *models.Category
	if result := c.db.WithContext(ctx).
		Table(CATEGORY_TABLE).
		Select(CATEGORY_COLUMNS).
		Where(CATEGORY_WHERE, categoryId).
		Find(&model); result.Error != nil {
		return nil, result.Error
	}
	if model.ID == 0 {
		return nil, errs.NewCRUDError(fmt.Sprintf("Not found category id %s", categoryId))
	}
	category, err := c.adapter.Rebuild(model)
	if err != nil {
		return nil, err
	}
	return category, nil
}
