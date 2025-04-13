package api

import (
	"ecommerce-service/configs"
	"ecommerce-service/internal/api/rest"
	"ecommerce-service/internal/api/rest/handlers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	// app.Get("/health", HealthCheck)

	restHandler := &rest.RestHandler{
		App: app,
	}

	setupRoutes(restHandler)

	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I'm breathing",
	})
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}
