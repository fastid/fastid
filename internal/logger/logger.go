package logger

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type RequestID string

type Logger interface {
	withField(ctx context.Context) *logrus.Entry
	GetLogger() *logrus.Logger

	Debug(ctx context.Context, msg string)
	Info(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
	Trace(ctx context.Context, msg string)
	Fatal(ctx context.Context, msg string)
	Debugf(ctx context.Context, msg string, args ...interface{})
	Infof(ctx context.Context, msg string, args ...interface{})
	Warnf(ctx context.Context, msg string, args ...interface{})
	Tracef(ctx context.Context, msg string, args ...interface{})
	Fatalf(ctx context.Context, msg string, args ...interface{})
}

type logger struct {
	logger *logrus.Logger
}

func New(cfg *config.Config) Logger {
	l := logrus.New()

	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(os.Stdout)
	l.SetReportCaller(false)

	if strings.ToLower(cfg.LOGGER.Level) == "debug" {
		l.SetLevel(logrus.DebugLevel)
	} else if strings.ToLower(cfg.LOGGER.Level) == "info" {
		l.SetLevel(logrus.InfoLevel)
	} else if strings.ToLower(cfg.LOGGER.Level) == "warn" {
		l.SetLevel(logrus.WarnLevel)
	} else if strings.ToLower(cfg.LOGGER.Level) == "trace" {
		l.SetLevel(logrus.TraceLevel)
	} else if strings.ToLower(cfg.LOGGER.Level) == "fatal" {
		l.SetLevel(logrus.FatalLevel)
	}

	return &logger{logger: l}
}

func (l *logger) withField(ctx context.Context) *logrus.Entry {
	if ctx.Value(RequestID("requestID")) != nil {
		return l.logger.WithField("x-request-id", ctx.Value("requestID").(string))
	} else {
		return l.logger.WithField("x-request-id", nil)
	}
}

func (l *logger) GetLogger() *logrus.Logger {
	return l.logger
}

func (l *logger) Debug(ctx context.Context, msg string) {
	l.withField(ctx).Debug(msg)
}

func (l *logger) Info(ctx context.Context, msg string) {
	l.withField(ctx).Info(msg)
}

func (l *logger) Warn(ctx context.Context, msg string) {
	l.withField(ctx).Warn(msg)
}

func (l *logger) Trace(ctx context.Context, msg string) {
	l.withField(ctx).Trace(msg)
}

func (l *logger) Fatal(ctx context.Context, msg string) {
	l.withField(ctx).Fatal(msg)
}

func (l *logger) Debugf(ctx context.Context, msg string, args ...interface{}) {
	l.withField(ctx).Debugf(msg, args...)
}

func (l *logger) Infof(ctx context.Context, msg string, args ...interface{}) {
	l.withField(ctx).Infof(msg, args...)
}

func (l *logger) Warnf(ctx context.Context, msg string, args ...interface{}) {
	l.withField(ctx).Warnf(msg, args...)
}

func (l *logger) Tracef(ctx context.Context, msg string, args ...interface{}) {
	l.withField(ctx).Warnf(msg, args...)
}

func (l *logger) Fatalf(ctx context.Context, msg string, args ...interface{}) {
	l.withField(ctx).Fatalf(msg, args...)
}
