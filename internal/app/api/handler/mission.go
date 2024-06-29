package handler

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/markraiter/spycat/internal/app/service"
	"github.com/markraiter/spycat/internal/domain"
	"github.com/markraiter/spycat/internal/lib/sl"
)

type MissionService interface {
	SaveMission(ctx context.Context, mr *domain.MissionRequest) (int, error)
	Missions(ctx context.Context) ([]*domain.Mission, error)
	MissionByID(ctx context.Context, id int) (*domain.Mission, error)
	AssignMissionToCat(ctx context.Context, catID, missionID int) error
	CompleteMission(ctx context.Context, id int) error
}

type MissionHandler struct {
	log     *slog.Logger
	val     *validator.Validate
	service MissionService
}

// @Summary Create mission
// @Description Create mission
// @Security ApiKeyAuth
// @Tags Mission
// @Accept json
// @Produce json
// @Param Create_mission_request body domain.MissionRequest true "Mission data"
// @Success 201 {integer} int "Mission ID"
// @Failure 400 {object} domain.Response
// @Failure 403 {object} domain.Response
// @Failure 406 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /missions [post]
func (h *MissionHandler) CreateMission(c *fiber.Ctx) error {
	const op = "handler.CreateMission"
	log := h.log.With(slog.String("operation", op))

	var mr domain.MissionRequest
	if err := c.BodyParser(&mr); err != nil {
		log.Warn("error while parsing input body", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	id, err := h.service.SaveMission(c.Context(), &mr)
	if err != nil {
		if errors.Is(err, service.ErrTooManyTargets) {
			log.Warn("too many targets", sl.Err(err))
			return c.Status(fiber.StatusForbidden).JSON(domain.Response{Message: err.Error()})
		}

		log.Warn("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}

// @Summary Get missions
// @Description Get missions
// @Security ApiKeyAuth
// @Tags Mission
// @Accept json
// @Produce json
// @Success 200 {array} domain.Mission "Missions"
// @Failure 500 {object} domain.Response
// @Router /missions [get]
func (h *MissionHandler) GetMissions(c *fiber.Ctx) error {
	const op = "handler.GetMissions"
	log := h.log.With(slog.String("operation", op))

	missions, err := h.service.Missions(c.Context())
	if err != nil {
		log.Warn("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(missions)
}

// @Summary Get mission by ID
// @Description Get mission by ID
// @Security ApiKeyAuth
// @Tags Mission
// @Accept json
// @Produce json
// @Param id path int true "Mission ID"
// @Success 200 {object} domain.Mission "Mission"
// @Failure 400 {object} domain.Response
// @Failure 404 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /missions/{id} [get]
func (h *MissionHandler) GetMission(c *fiber.Ctx) error {
	const op = "handler.GetMission"
	log := h.log.With(slog.String("operation", op))

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Warn("error while parsing input params", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	mission, err := h.service.MissionByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Warn("mission not found", sl.Err(err))
			return c.Status(fiber.StatusNotFound).JSON(domain.Response{Message: err.Error()})
		}

		log.Warn("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(mission)
}

// @Summary Assign mission to cat
// @Description Assign mission to cat
// @Security ApiKeyAuth
// @Tags Mission
// @Accept json
// @Produce json
// @Param cat_id path int true "Cat ID"
// @Param mission_id path int true "Mission ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 403 {object} domain.Response
// @Failure 404 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /missions/{mission_id}/cats/{cat_id} [patch]
func (h *MissionHandler) AssignMissionToCat(c *fiber.Ctx) error {
	const op = "handler.AssignMissionToCat"
	log := h.log.With(slog.String("operation", op))

	catID, err := c.ParamsInt("cat_id")
	if err != nil {
		log.Warn("error while parsing input params", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	missionID, err := c.ParamsInt("mission_id")
	if err != nil {
		log.Warn("error while parsing input params", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	err = h.service.AssignMissionToCat(c.Context(), catID, missionID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Warn("mission or cat not found", sl.Err(err))
			return c.Status(fiber.StatusNotFound).JSON(domain.Response{Message: err.Error()})
		}

		log.Warn("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{Message: fmt.Sprintf("mission assigned to cat: %d", catID)})
}

// @Summary Complete mission
// @Description Complete mission
// @Security ApiKeyAuth
// @Tags Mission
// @Accept json
// @Produce json
// @Param id path int true "Mission ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 404 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /missions/{id} [patch]
func (h *MissionHandler) CompleteMission(c *fiber.Ctx) error {
	const op = "handler.CompleteMission"
	log := h.log.With(slog.String("operation", op))

	id, err := c.ParamsInt("id")
	if err != nil {
		log.Warn("error while parsing input params", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	err = h.service.CompleteMission(c.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Warn("mission not found", sl.Err(err))
			return c.Status(fiber.StatusNotFound).JSON(domain.Response{Message: err.Error()})
		}

		log.Warn("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{Message: fmt.Sprintf("mission completed: %d", id)})
}
