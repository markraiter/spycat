package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/markraiter/spycat/docs"
	"github.com/markraiter/spycat/internal/app/api"
	"github.com/markraiter/spycat/internal/app/api/handler"
	"github.com/markraiter/spycat/internal/app/service"
	"github.com/markraiter/spycat/internal/app/storage/postgres"
	"github.com/markraiter/spycat/internal/config"
	"github.com/markraiter/spycat/internal/domain"
)

// @title SpyCat API
// @version	1.0
// @description	Docs for SpyCat API
// @contact.name Mark Raiter
// @contact.email raitermark@proton.me
// @host localhost:8888
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.MustLoad()

	fmt.Printf("config: %+v\n", cfg)

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	validate := validator.New()
	validate.RegisterValidation("number", domain.ValidateContainsNumber, false)   // nolint: errcheck
	validate.RegisterValidation("upper", domain.ValidateContainsUpper, false)     // nolint: errcheck
	validate.RegisterValidation("lower", domain.ValidateContainsLower, false)     // nolint: errcheck
	validate.RegisterValidation("special", domain.ValidateContainsSpecial, false) // nolint: errcheck

	log.Info("Starting application...")
	log.Info("port: " + cfg.Server.Port)

	storage := postgres.New(cfg.Postgres)

	service := service.New(
		storage,
	)

	handler := handler.New(
		log,
		validate,
		cfg,
		service,
	)

	server := api.New(cfg, handler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := server.HTTPServer.Listen(cfg.Server.Port); err != nil {
			log.Error("HTTPServer.Listen", "error", err)
		}
	}()

	<-stop

	if err := server.HTTPServer.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Error("ShutdownWithTimeout", "error", err)
	}

	if err := server.HTTPServer.Shutdown(); err != nil {
		log.Error("Shutdown", "error", err)
	}

	log.Info("server stopped")
}
