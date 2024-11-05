package router

import (
	"example.com/fxdemo/api-gateway/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(handler *handler.Handler) *fiber.App {
	app := fiber.New()

	app.Post("/users", handler.CreateUser)
	app.Get("/users/:id", handler.GetUserByID)

	return app
}
