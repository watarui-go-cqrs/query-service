package repository

import (
	"context"
	"fmt"
	"query-service/internal/domain/models/products"
	"query-service/internal/errs"
	"query-service/internal/infrastructure/gorm/handler"
	"query-service/internal/infrastructure/gorm/models"

	"gorm.io/gorm"
)

const (
	PRODUCT_TABLE   = "product"
	PRODUCT_COLUMNS = "product.obj_id AS p_id , product.name AS p_name , product.price AS p_price , product.category_id AS c_id , category.name AS c_name"
	PRODUCT_JOIN    = "JOIN category ON product.category_id = category.obj_id"
	PRODUCT_WHERE   = "product.obj_id = ?"
	PRODUCT_LIKE    = "product.name LIKE ?"
)

type productRepositoryGORM struct {
	db      *gorm.DB
	adapter products.ProductAdapter
}

func NewProductRepositoryGORM(db *gorm.DB, adapter products.ProductAdapter) products.ProductRepository {
	return &productRepositoryGORM{
		db:      db,
		adapter: adapter,
	}
}

func (p *productRepositoryGORM) List(ctx context.Context) ([]*products.Product, error) {
	models := []*models.Product{}
	if result := p.db.WithContext(ctx).
		Table(PRODUCT_TABLE).
		Select(PRODUCT_COLUMNS).
		Joins(PRODUCT_JOIN).
		Find(&models); result.Error != nil {
		return nil, handler.DBErrHandler(result.Error)
	}
	if products, err := p.createSlice(models); err != nil {
		return nil, err
	} else {
		return products, nil
	}
}

func (p *productRepositoryGORM) FindByProductId(ctx context.Context, productId string) (*products.Product, error) {
	model := models.Product{}
	if result := p.db.WithContext(ctx).
		Table(PRODUCT_TABLE).
		Select(PRODUCT_COLUMNS).
		Joins(PRODUCT_JOIN).
		Where(PRODUCT_WHERE, productId).
		Find(&model); result.Error != nil {
		return nil, handler.DBErrHandler(result.Error)
	}
	if model.ObjId == "" {
		return nil, errs.NewCRUDError(fmt.Sprintf("Not found product id %s", productId))
	}
	if product, err := p.adapter.Rebuild(model); err != nil {
		return nil, err
	} else {
		return product, nil
	}
}

func (p *productRepositoryGORM) FindByProductNameLike(ctx context.Context, keyword string) ([]*products.Product, error) {
	models := []*models.Product{}
	if result := p.db.WithContext(ctx).
		Table(PRODUCT_TABLE).
		Select(PRODUCT_COLUMNS).
		Joins(PRODUCT_JOIN).
		Where(PRODUCT_LIKE, "%"+keyword+"%").
		Find(&models); result.Error != nil {
		return nil, handler.DBErrHandler(result.Error)
	}
	if len(models) == 0 {
		return nil, errs.NewCRUDError(fmt.Sprintf("Not found product name like [%s]", keyword))
	}
	if products, err := p.createSlice(models); err != nil {
		return nil, err
	} else {
		return products, nil
	}
}

func (p *productRepositoryGORM) createSlice(results []*models.Product) ([]*products.Product, error) {
	var products []*products.Product
	for _, result := range results {
		product, err := p.adapter.Rebuild(result)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
