package handler

import (
	"log/slog"

	"github.com/go-playground/validator"
	"github.com/markraiter/spycat/internal/config"
)

type IService interface {
	AuthService
	CatService
	MissionService
	TargetService
}

type Handler struct {
	AuthHandler
	CatHandler
	MissionHandler
	TargetHandler
}

// New returns new instance of the Handler.
func New(
	log *slog.Logger,
	val *validator.Validate,
	cfg *config.Config,
	i IService,
) *Handler {
	return &Handler{
		AuthHandler: AuthHandler{
			cfg:     cfg,
			log:     log,
			val:     val,
			service: i,
		},
		CatHandler: CatHandler{
			log:     log,
			val:     val,
			service: i,
		},
		MissionHandler: MissionHandler{
			log:     log,
			val:     val,
			service: i,
		},
		TargetHandler: TargetHandler{
			log:     log,
			val:     val,
			service: i,
		},
	}
}
