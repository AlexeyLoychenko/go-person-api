package httpv1

import (
	"github.com/AlexeyLoychenko/person_api/internal/usecase"
	"github.com/AlexeyLoychenko/person_api/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type GetByIdRequest struct {
	Id int `json:"id"`
}

func NewRouter(router *chi.Mux, log logger.Logger, uc usecase.UseCase) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	pc := NewPersonController(log, uc)
	pc.RegisterRoutes(router)
}
