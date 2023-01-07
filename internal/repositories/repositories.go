package repositories

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
)

type Repositories interface {
	Users() Users
}

type repositories struct {
	cfg    *config.Config
	logger logger.Logger
	db     db.DB
	users  Users
}

func New(cfg *config.Config, logger logger.Logger, db db.DB) Repositories {
	users := NewUsersRepository(cfg, logger, db)

	repo := &repositories{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}

	repo.users = users
	return repo
}

func (r *repositories) Users() Users {
	return r.users
}
