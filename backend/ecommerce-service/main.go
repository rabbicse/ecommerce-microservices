package main

import (
	"ecommerce-service/configs"
	"ecommerce-service/internal/api"
	"log"
)

func main() {
	log.Println("Starting server...")

	cfg, err := configs.SetupEnv()

	if err != nil {
		log.Fatalf("config file is not loaded properly: %v\n", err)
	}

	api.StartServer(cfg)
}
