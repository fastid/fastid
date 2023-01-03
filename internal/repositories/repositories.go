package repositories

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
)

type Repositories interface {
	Keys() Keys
	Users() Users
}

type repositories struct {
	cfg    *config.Config
	logger logger.Logger
	db     db.DB
	keys   Keys
	users  Users
}

func New(cfg *config.Config, logger logger.Logger, db db.DB) Repositories {
	keys := NewKeysRepository(cfg, logger, db)
	users := NewUsersRepository(cfg, logger, db)

	repo := &repositories{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}

	repo.keys = keys
	repo.users = users
	return repo
}

func (r *repositories) Keys() Keys {
	return r.keys
}

func (r *repositories) Users() Users {
	return r.users
}
