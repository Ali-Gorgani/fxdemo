package server

import (
	"context"

	"example.com/fxdemo/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Server struct {
	app *fiber.App
	logger *logger.Logger
}

func NewServer(app *fiber.App, logger *logger.Logger) *Server {
	return &Server{
		app: app, 
		logger: logger,
	}
}
func (srv *Server) SetupServer(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			srv.logger.Info("Starting server")
			go srv.app.Listen(":3000")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.logger.Info("Shutting down server")
			return srv.app.Shutdown()
		},
	})
}
