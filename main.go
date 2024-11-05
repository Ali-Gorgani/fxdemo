package main

import (
	"example.com/fxdemo/api-gateway/handler"
	"example.com/fxdemo/api-gateway/repository"
	"example.com/fxdemo/api-gateway/router"
	"example.com/fxdemo/api-gateway/server"
	"example.com/fxdemo/api-gateway/usecase"
	"example.com/fxdemo/db"
	"example.com/fxdemo/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		logger.Module,
		db.Module,
		fx.Provide(
			handler.NewHandler,
			router.SetupRouter,
			fx.Annotate(
				repository.NewEntRepository,
				fx.As(new(repository.IRepository)),
			),
			fx.Annotate(
				usecase.NewUsecase,
				fx.As(new(usecase.IUsecase)),
			),
		),
		fx.Invoke(func(lc fx.Lifecycle, app *fiber.App, logger *logger.Logger) {
			server.NewServer(app, logger).SetupServer(lc)
		}),
	)
	app.Run()
}
