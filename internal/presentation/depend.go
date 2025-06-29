package presentation

import (
	"query-service/internal/infrastructure/gorm"
	"query-service/internal/presentation/builder"
	"query-service/internal/presentation/prepare"
	"query-service/internal/presentation/server"

	"go.uber.org/fx"
)

var QueryDepend = fx.Options(
	gorm.RepDepend,
	fx.Provide(
		builder.NewResultBuilderImpl,
		server.NewCategoryServer,
		server.NewProductServer,
		prepare.NewQueryServer,
	),
	fx.Invoke(
		prepare.QueryServiceLifecycle,
	),
)
