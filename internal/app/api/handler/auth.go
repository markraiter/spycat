package handler

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/markraiter/spycat/internal/app/service"
	"github.com/markraiter/spycat/internal/config"
	"github.com/markraiter/spycat/internal/domain"
	"github.com/markraiter/spycat/internal/lib/sl"
)

type AuthService interface {
	Register(ctx context.Context, user *domain.UserRequest) (int, error)
	Login(ctx context.Context, cfg config.Auth, email, password string) (string, error)
}

type AuthHandler struct {
	cfg     *config.Config
	log     *slog.Logger
	val     *validator.Validate
	service AuthService
}

// @Summary Register user
// @Description Register user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Register_request body domain.UserRequest true "User data"
// @Success 201 {integer} int "User ID"
// @Failure 400 {object} domain.Response
// @Failure 403 {object} domain.Response
// @Failure 406 {object} domain.Response
// @Failure 500 {object} domain.Response
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	const op = "handler.Register"
	log := h.log.With(slog.String("operation", op))

	var rr domain.UserRequest
	if err := c.BodyParser(&rr); err != nil {
		log.Warn("error while parsing input body", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	if err := h.val.Struct(rr); err != nil {
		log.Warn("validation error", sl.Err(err))
		return c.Status(fiber.StatusNotAcceptable).JSON(domain.Response{Message: err.Error()})
	}

	id, err := h.service.Register(c.Context(), &rr)
	if err != nil {
		if errors.Is(err, service.ErrAlreadyExists) {
			log.Warn("user already exists", sl.Err(err))
			return c.Status(fiber.StatusForbidden).JSON(domain.Response{Message: err.Error()})
		}
		log.Warn("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}

// @Summary Login
// @Tags Auth
// @Description	Logs user in
// @ID login
// @Accept json
// @Produce json
// @Param input	body domain.LoginRequest true "credentials"
// @Success	200	{string} string "Token"
// @Failure	400	{object} domain.Response
// @Failure	406	{object} domain.Response
// @Failure	500	{object} domain.Response
// @Router /auth/login [post].
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	const op = "handler.Login"
	log := h.log.With(slog.String("op", op))

	var loginReq domain.LoginRequest
	if err := c.BodyParser(&loginReq); err != nil {
		log.Warn("error while parsing input body", sl.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{Message: err.Error()})
	}

	if err := h.val.Struct(loginReq); err != nil {
		log.Warn("validation error", sl.Err(err))
		return c.Status(fiber.StatusNotAcceptable).JSON(domain.Response{Message: err.Error()})
	}

	token, err := h.service.Login(c.UserContext(), h.cfg.Auth, loginReq.Email, loginReq.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			log.Warn("invalid credentials", sl.Err(err))
			return c.Status(fiber.StatusForbidden).JSON(domain.Response{Message: err.Error()})
		}
		log.Error("internal error", sl.Err(err))
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(token)
}

// @Summary Logout
// @Description	Logs user out
// @Security ApiKeyAuth
// @Tags Auth
// @Produce json
// @Success 200	{object} domain.Response
// @Router /auth/logout [post].
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Request().Header.Del("Authorization")
	c.Response().Header.Del("Authorization")

	return c.Status(fiber.StatusOK).JSON(domain.Response{Message: "you are logged out"})
}
