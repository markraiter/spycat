package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/swagger"
	"github.com/markraiter/spycat/internal/app/api/handler"
	"github.com/markraiter/spycat/internal/app/api/middleware"
	"github.com/markraiter/spycat/internal/config"
)

// initRoutes configures the routes for the app.
func (s Server) initRoutes(app *fiber.App, handler *handler.Handler, cfg *config.Config) {
	basicAuth := middleware.NewUserIdentity(cfg.Auth)

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api/v1")
	{
		authentication := api.Group("/auth")
		{
			authentication.Post("/register", timeout.NewWithContext(handler.Register, cfg.Server.WriteTimeout))
			authentication.Post("/login", timeout.NewWithContext(handler.Login, cfg.Server.WriteTimeout))
			authentication.Post("/logout", basicAuth, timeout.NewWithContext(handler.Logout, cfg.Server.WriteTimeout))
		}

		cats := api.Group("/cats")
		{
			cats.Post("/", basicAuth, timeout.NewWithContext(handler.CreateCat, cfg.Server.WriteTimeout))
			cats.Get("/", basicAuth, timeout.NewWithContext(handler.GetCats, cfg.Server.ReadTimeout))
			cats.Get("/:id", basicAuth, timeout.NewWithContext(handler.GetCat, cfg.Server.ReadTimeout))
			cats.Put("/:id", basicAuth, timeout.NewWithContext(handler.UpdateCat, cfg.Server.WriteTimeout))
			cats.Delete("/:id", basicAuth, timeout.NewWithContext(handler.DeleteCat, cfg.Server.WriteTimeout))
		}

		missions := api.Group("/missions")
		{
			missions.Post("/", basicAuth, timeout.NewWithContext(handler.CreateMission, cfg.Server.WriteTimeout))
			missions.Get("/", basicAuth, timeout.NewWithContext(handler.GetMissions, cfg.Server.ReadTimeout))
			missions.Get("/:id", basicAuth, timeout.NewWithContext(handler.GetMission, cfg.Server.ReadTimeout))
		}

	}
}
