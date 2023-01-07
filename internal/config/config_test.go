package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := New()
	require.NoError(t, err)
	assert.Equal(t, cfg.HTTP.Listen, ":8000")
}
