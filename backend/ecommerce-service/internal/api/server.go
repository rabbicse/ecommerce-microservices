package api

import (
	"ecommerce-service/configs"
	"ecommerce-service/internal/api/rest"
	"ecommerce-service/internal/api/rest/handlers"
	"ecommerce-service/internal/domain"
	"ecommerce-service/internal/helper"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error: %v\n", err)
	}
	log.Println("database connected!")

	// run migration
	db.AutoMigrate(&domain.User{})

	auth := helper.SetupAuth(config.AppSecret)

	app.Get("/health", HealthCheck)

	restHandler := &rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: auth,
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
