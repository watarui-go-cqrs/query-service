package main

import (
	"query-service/internal/presentation"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		presentation.QueryDepend,
	).Run()
}
