package handler

import (
	"log/slog"

	"github.com/go-playground/validator"
	"github.com/markraiter/spycat/internal/config"
)

type IAuthService interface {
}

type Handler struct {
}

// New returns new instance of the Handler.
func New(
	log *slog.Logger,
	val *validator.Validate,
	cfg *config.Config,
	a IAuthService,
) *Handler {
	return &Handler{}
}
