package repositories

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	log "github.com/sirupsen/logrus"
)

type Repositories interface{}

type repositories struct {
	cfg *config.Config
	log *log.Logger
	db  db.DB
}

func New(cfg *config.Config, log *log.Logger, db db.DB) Repositories {
	return &repositories{
		cfg: cfg,
		log: log,
		db:  db,
	}
}
