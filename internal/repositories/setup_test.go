package repositories

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/migrations"
	"testing"
)

func setupSuite(tb testing.TB) (func(tb testing.TB), Repositories, context.Context, error) {

	cfg, err := config.New()
	if err != nil {
		return nil, nil, nil, err
	}

	log := logger.New(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	cfg.Scheme = "migration_test"

	// Context
	ctx := context.Background()

	database, err := db.New(cfg, ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	err = database.CreateScheme(ctx, cfg.Scheme)
	if err != nil {
		return nil, nil, nil, err
	}

	migration, err := migrations.New(cfg, database)
	if err != nil {
		return nil, nil, nil, err
	}

	err = migration.Upgrade()
	if err != nil {
		return nil, nil, nil, err
	}

	repo := New(cfg, log, database)

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

	return teardown, repo, ctx, nil
}
