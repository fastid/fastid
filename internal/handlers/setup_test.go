package handlers

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/migrations"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
	"testing"
)

func setupSuite(tb testing.TB) (func(tb testing.TB), context.Context, *echo.Echo, Handlers, error) {

	// Context
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		return nil, ctx, nil, nil, err
	}

	// Change scheme
	cfg.Scheme = "migration_test"

	// Logger
	log := logger.New(cfg)

	// DB
	database, err := db.New(cfg, ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if err = database.DropScheme(ctx, cfg.Scheme); err != nil {
		return nil, nil, nil, nil, err
	}

	if err = database.CreateScheme(ctx, cfg.Scheme); err != nil {
		return nil, nil, nil, nil, err
	}

	migration, err := migrations.New(cfg, database)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if err = migration.Upgrade(); err != nil {
		return nil, nil, nil, nil, err
	}

	// Storage
	repos := repositories.New(cfg, log, database)

	// Service
	srv := services.New(cfg, log, repos)

	// Handlers
	handler := New(cfg, log, srv)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	group := e.Group("/api/v1")
	handler.Register(group)

	teardown := func(tb testing.TB) {
		err := migration.Drop()
		if err != nil {
			return
		}

		err = database.DropScheme(ctx, cfg.Scheme)
		if err != nil {
			return
		}
	}
	return teardown, ctx, e, handler, err
}
