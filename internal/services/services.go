package services

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
)

type Services interface {
	Server() Server
	Users() Users
}

type services struct {
	cfg          *config.Config
	logger       logger.Logger
	repositories repositories.Repositories
	server       Server
	users        Users
}

func New(cfg *config.Config, logger logger.Logger, repositories repositories.Repositories) Services {
	server := NewServerService(cfg, logger, repositories)
	users := NewUsersService(cfg, logger, repositories)

	srv := services{cfg: cfg, logger: logger, repositories: repositories}
	srv.server = server
	srv.users = users
	return &srv
}

func (s *services) Server() Server {
	return s.server
}

func (s *services) Users() Users {
	return s.users
}
