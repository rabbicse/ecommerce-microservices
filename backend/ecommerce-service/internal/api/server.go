package api

import (
	"ecommerce-service/configs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I'm breathing",
	})
}
