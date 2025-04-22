package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	AppSecret  string
}

func SetupEnv() (cfg AppConfig, err error) {
	// if os.Getenv("APP_ENV") == "dev" {
	godotenv.Load()
	// }
	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variable not found!")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("dsn env variable not found!")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("appsecret env variable not found!")
	}

	return AppConfig{ServerPort: httpPort, Dsn: dsn, AppSecret: appSecret}, nil
}
