package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	AppModeContainerized = "containerized"
)

type Config struct {
	AppPort                int
	LogLevel               string
	LogHandlerType         string
	HttpServerReadTimeout  int
	HttpServerWriteTimeout int
	HttpClientTimeout      int
	PostgresDSN            string
}

func Load() (*Config, error) {
	var cfg Config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	cfg.AppPort, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal("config - Load(): error parsing AppPort, integer required")
	}
	cfg.LogLevel = os.Getenv("LOG_LEVEL")
	cfg.LogHandlerType = os.Getenv("LOG_HANDLER_TYPE")

	cfg.HttpServerReadTimeout, err = strconv.Atoi(os.Getenv("HTTP_SERVER_READ_TIMEOUT"))
	if err != nil {
		log.Fatal("config - Load(): error parsing HttpServerReadTimeout, integer required")
	}

	cfg.HttpServerWriteTimeout, err = strconv.Atoi(os.Getenv("HTTP_SERVER_WRITE_TIMEOUT"))
	if err != nil {
		log.Fatal("config - Load(): error parsing HttpServerWriteTimeout, integer required")
	}

	cfg.HttpClientTimeout, err = strconv.Atoi(os.Getenv("HTTP_CLIENT_TIMEOUT"))
	if err != nil {
		log.Fatal("config - Load(): error parsing HttpClientTimeout, integer required")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal("config - Load(): error parsing DBPort, integer required")
	}
	dbHost := os.Getenv("POSTGRES_HOST")
	if os.Getenv("APP_MODE") == AppModeContainerized {
		dbHost = os.Getenv("DB_DOCKER_CONTAINER")
	}

	cfg.PostgresDSN = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	return &cfg, nil
}
