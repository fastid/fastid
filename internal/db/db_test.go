package db

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDB(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	ctx := context.Background()

	pool, err := New(cfg, ctx)
	require.NoError(t, err)

	db := pool.GetConnect()

	var version string
	err = db.QueryRow(ctx, "SELECT VERSION() as version;").Scan(&version)
	require.NoError(t, err)
	require.NotEmpty(t, version)
}
