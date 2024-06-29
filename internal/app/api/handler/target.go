package handler

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/markraiter/spycat/internal/app/service"
	"github.com/markraiter/spycat/internal/domain"
	"github.com/markraiter/spycat/internal/lib/sl"
)

type TargetService interface {
	CompleteTarget(ctx context.Context, id int) error
	AddTargetToMission(ctx context.Context, missionID, targetID int) error
}

type TargetHandler struct {
	log     *slog.Logger
	val     *validator.Validate
	service TargetService
}

// @Summary Complete target
// @Description Complete target
// @Security ApiKeyAuth
// @Tags Target
// @Accept json
// @Produce json
// @Param id path int true "Target ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 403 {object} domain.Response
// @Failure 404 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /targets/{id} [patch]
func (h *TargetHandler) CompleteTarget(c *fiber.Ctx) error {
	const op = "handler.CompleteTarget"
	log := h.log.With(slog.String("operation", op))

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Warn("error while parsing input id", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	err = h.service.CompleteTarget(c.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Warn("target not found", sl.Err(err))
			return c.Status(fiber.StatusNotFound).JSON(domain.Response{Message: err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{Message: fmt.Sprintf("Target %d completed", id)})
}

// @Summary Add target to mission
// @Description Add target to mission
// @Security ApiKeyAuth
// @Tags Target
// @Accept json
// @Produce json
// @Param mission_id path int true "Mission ID"
// @Param target_id path int true "Target ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.Response
// @Failure 403 {object} domain.Response
// @Failure 404 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /missions/{mission_id}/targets/{target_id} [patch]
func (h *TargetHandler) AddTargetToMission(c *fiber.Ctx) error {
	const op = "handler.AddTargetToMission"
	log := h.log.With(slog.String("operation", op))

	missionID, err := strconv.Atoi(c.Params("mission_id"))
	if err != nil {
		log.Warn("error while parsing input mission_id", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	targetID, err := strconv.Atoi(c.Params("target_id"))
	if err != nil {
		log.Warn("error while parsing input target_id", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	err = h.service.AddTargetToMission(c.Context(), missionID, targetID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			log.Warn("mission or target not found", sl.Err(err))
			return c.Status(fiber.StatusNotFound).JSON(domain.Response{Message: err.Error()})
		}
		if errors.Is(err, service.ErrMissionCompleted) {
			log.Warn("mission completed", sl.Err(err))
			return c.Status(fiber.StatusForbidden).JSON(domain.Response{Message: err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{Message: fmt.Sprintf("Target %d added to mission %d", targetID, missionID)})
}
