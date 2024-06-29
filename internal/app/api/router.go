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

	}
}
