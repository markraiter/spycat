package handler

import (
	"log/slog"

	"github.com/go-playground/validator"
	"github.com/markraiter/spycat/internal/config"
)

type IAuthService interface {
	AuthService
}

type Handler struct {
	AuthHandler
}

// New returns new instance of the Handler.
func New(
	log *slog.Logger,
	val *validator.Validate,
	cfg *config.Config,
	a IAuthService,
) *Handler {
	return &Handler{
		AuthHandler: AuthHandler{
			cfg:     cfg,
			log:     log,
			val:     val,
			service: a,
		},
	}
}
