package http

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/handlers"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/migrations"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/internal/services"
	"github.com/fastid/fastid/internal/swagger"
	"github.com/fastid/fastid/internal/validators"
	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	internalLog "log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func HTTP() {

	// Configs
	cfg, err := config.New()
	if err != nil {
		internalLog.Fatalln(err.Error())
	}

	// Context
	ctx := context.Background()

	// Logger
	log := logger.New(cfg)
	log.Info(ctx, "Starting the server")

	// Echo
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	groupApi := e.Group("/api")
	groupApiV1 := groupApi.Group("/v1")

	// Prometheus
	prom := prometheus.NewPrometheus("fastid", nil)
	prom.Use(e)

	// Middleware
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		RequestIDHandler: func(e echo.Context, requestID string) {
			ctx := context.WithValue(e.Request().Context(), logger.KeyContext("requestID"), requestID)
			req := e.Request().WithContext(ctx)
			e.SetRequest(req)
		},
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogRemoteIP:  true,
		LogRequestID: true,
		LogMethod:    true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log := log.GetLogger().WithFields(logrus.Fields{
				"method":       values.Method,
				"uri":          values.URI,
				"status":       values.Status,
				"ip":           values.RemoteIP,
				"x-request-id": values.RequestID,
			})

			log.Infof("%s %s", values.Method, values.URI)
			return nil
		}},
	))

	// CORS
	if len(cfg.HTTP.CORS) > 0 {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: cfg.HTTP.CORS,
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
	}

	// DB
	database, err := db.New(cfg, ctx)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	// Migrations
	migration, err := migrations.New(cfg, database)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	if err = migration.Upgrade(); err != nil {
		log.Infof(ctx, "migration %s", err.Error())
	}

	// Repository
	repos := repositories.New(cfg, log, database)

	// Service
	srv := services.New(cfg, log, repos)

	// Validator
	validator := validators.New()
	validator.Register(e)

	// Handlers
	handler := handlers.New(cfg, log, srv)
	handler.Register(groupApiV1)

	// Swagger
	sw := swagger.New(cfg, log)
	sw.Register(groupApi)

	// Http server
	go func() {
		if err := e.Start(cfg.HTTP.Listen); err != nil && err != http.ErrServerClosed {
			log.Fatalf(ctx, "Shutting down the server %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctxShutdown); err != nil {
		log.Fatal(ctx, err.Error())
	}
}
