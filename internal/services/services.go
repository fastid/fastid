package services

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	log "github.com/sirupsen/logrus"
)

type Services interface{}

type services struct {
	cfg *config.Config
	log *log.Logger
}

func New(cfg *config.Config, log *log.Logger, repositories repositories.Repositories) Services {
	srv := services{cfg: cfg, log: log}
	return &srv
}
