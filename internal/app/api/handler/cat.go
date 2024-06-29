package handler

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/markraiter/spycat/internal/app/service"
	"github.com/markraiter/spycat/internal/domain"
	"github.com/markraiter/spycat/internal/lib/sl"
)

type CatService interface {
	SaveCat(ctx context.Context, cr *domain.CatRequest) (int, error)
	Cat(ctx context.Context, id int) (*domain.Cat, error)
	Cats(ctx context.Context) ([]*domain.Cat, error)
	UpdateCat(ctx context.Context, catID int, cr *domain.CatRequest) error
	DeleteCat(ctx context.Context, id int) error
}

type CatHandler struct {
	log     *slog.Logger
	val     *validator.Validate
	service CatService
}

// @Summary Create cat
// @Description Create cat
// @Security ApiKeyAuth
// @Tags Cat
// @Accept json
// @Produce json
// @Param Create_cat_request body domain.CatRequest true "Cat data"
// @Success 201 {integer} int "Cat ID"
// @Failure 400 {object} domain.Response
// @Failure 403 {object} domain.Response
// @Failure 406 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /cats [post]
func (h *CatHandler) CreateCat(c *fiber.Ctx) error {
	const op = "handler.CreateCat"
	log := h.log.With(slog.String("operation", op))

	var cr domain.CatRequest
	if err := c.BodyParser(&cr); err != nil {
		log.Warn("error while parsing input body", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	if err := h.val.Struct(cr); err != nil {
		log.Warn("validation error", sl.Err(err))
		return c.Status(fiber.StatusNotAcceptable).JSON(domain.Response{Message: err.Error()})
	}

	id, err := h.service.SaveCat(c.Context(), &cr)
	if err != nil {
		if errors.Is(err, service.ErrCatBreedNotFound) {
			log.Warn("cat breed not found", sl.Err(err))
			return c.Status(fiber.StatusNotAcceptable).JSON(domain.Response{Message: err.Error()})
		}
		{
			if errors.Is(err, service.ErrAlreadyExists) {
				log.Warn("cat already exists", sl.Err(err))
				return c.Status(fiber.StatusForbidden).JSON(domain.Response{Message: err.Error()})
			}
		}

		log.Error("error while saving cat", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}

// @Summary Get cat by ID
// @Description Get cat by ID
// @Security ApiKeyAuth
// @Tags Cat
// @Accept json
// @Produce json
// @Param id path int true "Cat ID"
// @Success 200 {object} domain.Cat
// @Failure 400 {object} domain.Response
// @Failure 404 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /cats/{id} [get]
func (h *CatHandler) GetCat(c *fiber.Ctx) error {
	const op = "handler.GetCat"
	log := h.log.With(slog.String("operation", op))

	p := struct {
		ID int `json:"id" validate:"required"`
	}{}

	if err := c.ParamsParser(&p); err != nil {
		log.Warn("error while parsing input parameter", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	cat, err := h.service.Cat(c.Context(), p.ID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Warn("cat not found", sl.Err(err))
			return c.Status(fiber.StatusNotFound).JSON(domain.Response{Message: err.Error()})
		}

		log.Error("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(cat)
}

// @Summary Get all cats
// @Description Get all cats
// @Security ApiKeyAuth
// @Tags Cat
// @Accept json
// @Produce json
// @Success 200 {array} domain.Cat
// @Failure 500 {object} domain.Response
// @Router /cats [get]
func (h *CatHandler) GetCats(c *fiber.Ctx) error {
	const op = "handler.GetCats"
	log := h.log.With(slog.String("operation", op))

	cats, err := h.service.Cats(c.Context())
	if err != nil {
		log.Error("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	if cats == nil {
		cats = []*domain.Cat{}
	}

	return c.Status(fiber.StatusOK).JSON(cats)
}

// @Summary Update cat by ID
// @Description Update cat by ID
// @Security ApiKeyAuth
// @Tags Cat
// @Accept json
// @Produce json
// @Param id path int true "Cat ID"
// @Param Update_cat_request body domain.CatRequest true "Cat data"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 404 {object} domain.Response
// @Failure 406 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /cats/{id} [put]
func (h *CatHandler) UpdateCat(c *fiber.Ctx) error {
	const op = "handler.UpdateCat"
	log := h.log.With(slog.String("operation", op))

	p := struct {
		ID int `json:"id" validate:"required"`
	}{}

	if err := c.ParamsParser(&p); err != nil {
		log.Warn("error while parsing input parameter", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	var cr domain.CatRequest
	if err := c.BodyParser(&cr); err != nil {
		log.Warn("error while parsing input body", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	if err := h.val.Struct(cr); err != nil {
		log.Warn("validation error", sl.Err(err))
		return c.Status(fiber.StatusNotAcceptable).JSON(domain.Response{Message: err.Error()})
	}

	err := h.service.UpdateCat(c.Context(), p.ID, &cr)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Warn("cat not found", sl.Err(err))
			return c.Status(fiber.StatusNotFound).JSON(domain.Response{Message: err.Error()})
		}
		if errors.Is(err, service.ErrCatBreedNotFound) {
			log.Warn("cat breed not found", sl.Err(err))
			return c.Status(fiber.StatusNotAcceptable).JSON(domain.Response{Message: err.Error()})
		}

		log.Error("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{Message: "Cat updated"})
}

// @Summary Delete cat by ID
// @Description Delete cat by ID
// @Security ApiKeyAuth
// @Tags Cat
// @Accept json
// @Produce json
// @Param id path int true "Cat ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 404 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /cats/{id} [delete]
func (h *CatHandler) DeleteCat(c *fiber.Ctx) error {
	const op = "handler.DeleteCat"
	log := h.log.With(slog.String("operation", op))

	p := struct {
		ID int `json:"id" validate:"required"`
	}{}

	if err := c.ParamsParser(&p); err != nil {
		log.Warn("error while parsing input parameter", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	err := h.service.DeleteCat(c.Context(), p.ID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Warn("cat not found", sl.Err(err))
			return c.Status(fiber.StatusNotFound).JSON(domain.Response{Message: err.Error()})
		}

		log.Error("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{Message: "Cat deleted"})
}
