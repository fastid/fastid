package services

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
)

type Services interface {
	Keys() Keys
}

type services struct {
	cfg          *config.Config
	logger       logger.Logger
	repositories repositories.Repositories
	keys         Keys
}

func New(cfg *config.Config, logger logger.Logger, repositories repositories.Repositories) Services {
	keys := NewKeyService(cfg, logger, repositories)

	srv := services{cfg: cfg, logger: logger, repositories: repositories}
	srv.keys = keys
	return &srv
}

func (s *services) Keys() Keys {
	return s.keys
}
