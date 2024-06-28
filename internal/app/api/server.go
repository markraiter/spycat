package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/markraiter/spycat/internal/app/api/handler"
	"github.com/markraiter/spycat/internal/config"
	"github.com/markraiter/spycat/internal/domain"
)

type Server struct {
	HTTPServer *fiber.App
}

// New returns new instance of the Server.
func New(cfg *config.Config, handler *handler.Handler) *Server {
	server := new(Server)

	fconfig := fiber.Config{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var localError *fiber.Error
			if errors.As(err, &localError) {
				code = localError.Code
			}

			c.Status(code)

			if err := c.JSON(domain.Response{Message: localError.Message}); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}

			return nil
		},
	}
	server.HTTPServer = fiber.New(fconfig)
	server.HTTPServer.Use(recover.New())
	server.HTTPServer.Use(logger.New())
	server.HTTPServer.Use(cors.New(corsConfig()))
	server.initRoutes(server.HTTPServer, handler, cfg)

	return server
}

func (s *Server) Shutdown(ctx context.Context) error {
	const op = "api.Server.Shutdown"

	return fmt.Errorf("%s: %w", op, s.HTTPServer.ShutdownWithContext(ctx))
}

func corsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Access-Control-Allow-Credentials, Authorization",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
		AllowCredentials: false,
	}
}
