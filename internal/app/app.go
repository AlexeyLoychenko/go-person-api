package app

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/AlexeyLoychenko/person_api/config"
	"github.com/AlexeyLoychenko/person_api/internal/controller/httpv1"
	"github.com/AlexeyLoychenko/person_api/internal/repo"
	"github.com/AlexeyLoychenko/person_api/internal/usecase"
	"github.com/AlexeyLoychenko/person_api/internal/webapi"
	"github.com/AlexeyLoychenko/person_api/pkg/logger"
	"github.com/AlexeyLoychenko/person_api/pkg/postgres"
	"github.com/go-chi/chi/v5"
)

func Run(cfg *config.Config) {
	//Logger
	log := logger.NewSlogWrapper(cfg.LogLevel, cfg.LogHandlerType)
	log.Info("Starting app with:", "config", *cfg)
	log.Debug("Debug-level messages are on")

	//Repository
	pgPool, err := postgres.New(cfg.PostgresDSN)
	if err != nil {
		log.Error("Postgres init failed", slog.Attr{Key: "Error", Value: slog.StringValue(err.Error())})
		os.Exit(1)
	}
	defer pgPool.Close()
	log.Info("pgPool created")

	// Use cases
	personUseCase := usecase.New(
		repo.New(pgPool),
		webapi.New(time.Duration(cfg.HttpClientTimeout)*time.Second),
	)

	// HTTP Server
	router := chi.NewRouter()
	httpv1.NewRouter(router, log, personUseCase)
	httpServer := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.AppPort),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.HttpServerReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.HttpServerWriteTimeout) * time.Second,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Error("Error starting server", "error", err)
	}
}
