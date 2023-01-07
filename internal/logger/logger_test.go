package logger

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogger(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	logger := New(cfg)

	ctx := context.TODO()
	ctxRequestID := context.WithValue(ctx, KeyContext("requestID"), "Test-RequestID")
	logger.Debug(ctxRequestID, "Test")
}
