package migrations

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMigrations(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	cfg.Scheme = "migration_test"

	// Context
	ctx := context.Background()

	database, err := db.New(cfg, ctx)
	require.NoError(t, err)

	err = database.CreateScheme(ctx, cfg.Scheme)
	require.NoError(t, err)

	migration, err := New(cfg, database)
	require.NoError(t, err)

	err = migration.Upgrade()
	require.NoError(t, err)

	err = migration.Downgrade()
	require.NoError(t, err)

	err = migration.Drop()
	require.NoError(t, err)

	err = database.DropScheme(ctx, cfg.Scheme)
	require.NoError(t, err)
}
