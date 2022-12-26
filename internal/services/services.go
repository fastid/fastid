package services

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type Services interface {
	Keys() Keys
}

type services[LOG log.Logger | log.Entry] struct {
	cfg          *config.Config
	log          *LOG
	repositories repositories.Repositories
	keys         Keys
}

func New(cfg *config.Config, log *log.Logger, repositories repositories.Repositories) Services {
	keys := NewKeyService(cfg, log, repositories)

	srv := services[logrus.Logger]{cfg: cfg, log: log, repositories: repositories}
	srv.keys = keys
	return &srv
}

func (s *services[T]) Keys() Keys {
	return s.keys
}
