package services

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	log "github.com/sirupsen/logrus"
)

type Services interface {
	Keys() Keys
}

type services struct {
	cfg          *config.Config
	log          *log.Logger
	repositories repositories.Repositories
	keys         Keys
}

func New(cfg *config.Config, log *log.Logger, repositories repositories.Repositories) Services {
	keys := NewKeyService(cfg, log, repositories)

	srv := services{cfg: cfg, log: log, repositories: repositories}
	srv.keys = keys
	return &srv
}

func (s *services) Keys() Keys {
	return s.keys
}
