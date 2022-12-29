package repositories

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
)

type Repositories interface {
	Keys() Keys
}

type repositories struct {
	cfg    *config.Config
	logger logger.Logger
	db     db.DB
	keys   Keys
}

func New(cfg *config.Config, logger logger.Logger, db db.DB) Repositories {
	keys := NewKeysRepository(cfg, logger, db)

	repo := &repositories{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}

	repo.keys = keys
	return repo
}

func (r *repositories) Keys() Keys {
	return r.keys
}
