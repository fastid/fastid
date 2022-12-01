package logger

import (
	"github.com/fastid/fastid/internal/config"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func New(cfg *config.Config) *log.Logger {
	var logger = log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	if strings.ToLower(cfg.LOGGER.Level) == "debug" {
		logger.SetLevel(log.DebugLevel)
	} else if strings.ToLower(cfg.LOGGER.Level) == "info" {
		logger.SetLevel(log.InfoLevel)
	} else if strings.ToLower(cfg.LOGGER.Level) == "warn" {
		logger.SetLevel(log.WarnLevel)
	} else if strings.ToLower(cfg.LOGGER.Level) == "trace" {
		logger.SetLevel(log.TraceLevel)
	}
	return logger
}
