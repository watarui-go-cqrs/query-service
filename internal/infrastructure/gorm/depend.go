package gorm

import (
	"query-service/internal/infrastructure/gorm/adapter"
	"query-service/internal/infrastructure/gorm/handler"
	"query-service/internal/infrastructure/gorm/repository"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var DBModule = fx.Provide(func() (*gorm.DB, error) {
	return handler.ConnectDB()
})

var RepDepend = fx.Options(
	DBModule,
	fx.Provide(
		adapter.NewCategoryAdapterImpl,
		adapter.NewProductAdapterImpl,

		repository.NewCategoryRepositoryGORM,
		repository.NewProductRepositoryGORM,
	),
)
