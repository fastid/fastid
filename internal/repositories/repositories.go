package repositories

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	log "github.com/sirupsen/logrus"
)

type Repositories interface {
	Keys() Keys
}

type repositories struct {
	cfg  *config.Config
	log  *log.Logger
	db   db.DB
	keys Keys
}

func New(cfg *config.Config, log *log.Logger, db db.DB) Repositories {
	keys := NewKeysRepository(cfg, log, db)

	repo := &repositories{
		cfg: cfg,
		log: log,
		db:  db,
	}

	repo.keys = keys
	return repo
}

func (r *repositories) Keys() Keys {
	return r.keys
}
