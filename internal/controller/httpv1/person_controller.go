package httpv1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AlexeyLoychenko/person_api/internal/entity"
	"github.com/AlexeyLoychenko/person_api/internal/model"
	"github.com/AlexeyLoychenko/person_api/internal/usecase"
	"github.com/AlexeyLoychenko/person_api/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type PersonController struct {
	l  logger.Logger
	uc usecase.UseCase
}

func NewPersonController(logger logger.Logger, usecase usecase.UseCase) *PersonController {
	return &PersonController{
		l:  logger,
		uc: usecase,
	}
}

func (c *PersonController) RegisterRoutes(router *chi.Mux) {
	router.Get("/person/{id}", c.GetById())
	router.Get("/person", c.GetList())
	router.Delete("/person/{id}", c.Delete())
	router.Post("/person/{id}", c.Update())
	router.Put("/person", c.Create())
}

func (c *PersonController) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := c.l.With(
			"handler", "PersonController - GetById()",
			"request_id", middleware.GetReqID(r.Context()),
		)

		paramId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("Failed to decode request", "request", r, "error", err)
			render.JSON(w, r, Error("Failed to decode request"))
			return
		}

		resp, err := c.uc.GetById(paramId)
		if err != nil {
			log.Error("Failed to perform request", "request", r, "error", err)
			render.JSON(w, r, Error("Failed to perform request"))
			return
		}

		render.JSON(w, r, model.WebResponse[entity.Person]{Data: resp})
		log.Info("Request completed")
	}
}

func (c *PersonController) GetList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := c.l.With(
			"handler", "PersonController - GetList()",
			"request_id", middleware.GetReqID(r.Context()),
		)

		pageId, _ := strconv.Atoi(r.URL.Query().Get("page_id"))
		age, _ := strconv.Atoi(r.URL.Query().Get("age"))
		personRequest := model.GetPersonRequest{
			Id:          r.URL.Query().Get("id"),
			Name:        r.URL.Query().Get("name"),
			Surname:     r.URL.Query().Get("surname"),
			Patronymic:  r.URL.Query().Get("patronymic"),
			Gender:      r.URL.Query().Get("gender"),
			Age:         age,
			Nationality: r.URL.Query().Get("nationality"),
			PageId:      pageId,
		}
		log.Debug("GetPersonRequest", "value", personRequest)

		data, page, err := c.uc.GetList(personRequest)
		if err != nil {
			log.Error("Failed to decode request", "request", r, "error", err)
			render.JSON(w, r, Error("Failed to decode request"))
			return
		}

		render.JSON(w, r, model.WebResponse[[]entity.Person]{
			Data: data,
			Paging: &model.PageMetadata{
				PageId:      page.PageId,
				HasNextPage: page.HasNextPage,
			}})
		log.Info("Request completed")
	}
}

func (c *PersonController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := c.l.With(
			"handler", "PersonController - Delete()",
			"request_id", middleware.GetReqID(r.Context()),
		)

		paramId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("Failed to decode request", "request", r, "error", err)
			render.JSON(w, r, Error("Failed to decode request"))
			return
		}

		res, err := c.uc.Delete(paramId)
		if err != nil {
			log.Error("Failed to delete record", "request", r, "error", err)
			render.JSON(w, r, Error("Failed to delete record"))
			return
		}

		render.JSON(w, r, model.WebResponse[bool]{Data: res})
		log.Info("Request completed")
	}
}

func (c *PersonController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := c.l.With(
			"handler", "PersonController - Update()",
			"request_id", middleware.GetReqID(r.Context()),
		)

		paramId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("Failed to decode parameter: id", "request", r, "error", err)
			render.JSON(w, r, Error("Bad request"))
			return
		}

		var updateReq model.UpdatePersonRequest
		err = json.NewDecoder(r.Body).Decode(&updateReq)
		log.Debug("UpdatePersonRequest", "value", updateReq)
		if err != nil {
			log.Error("Failed to decode request body", "requestBody", r.Body, "error", err)
			render.JSON(w, r, Error("Bad request body"))
			return
		}

		updateReq.Id = paramId
		res, err := c.uc.Update(updateReq)
		if err != nil {
			log.Error("Failed to update record", "request", r, "error", err)
			render.JSON(w, r, Error("Failed to update record"))
			return
		}

		render.JSON(w, r, model.WebResponse[bool]{Data: res})
		log.Info("Request completed")
	}
}

func (c *PersonController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := c.l.With(
			"handler", "PersonController - Create()",
			"request_id", middleware.GetReqID(r.Context()),
		)

		var createReq model.CreatePersonRequest
		err := json.NewDecoder(r.Body).Decode(&createReq)
		log.Debug("CreatePersonRequest", "value", createReq)
		if err != nil {
			log.Error("Failed to decode request Body", "requestBody", r.Body, "error", err)
			render.JSON(w, r, Error("Bad request body"))
			return
		}

		res, err := c.uc.Create(createReq)
		if err != nil {
			log.Error("Failed to create record", "request", r, "error", err)
			render.JSON(w, r, Error("Failed to create record"))
			return
		}

		render.JSON(w, r, model.WebResponse[entity.Person]{Data: res})
		log.Info("Request completed")
	}
}
