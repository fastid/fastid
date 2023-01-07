package random

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRnd(t *testing.T) {
	stringRandom := String(10)
	require.NotEmpty(t, stringRandom)
	require.Equal(t, len(stringRandom), 10)

	stringNumber := Int(10, 20)
	require.NotEmpty(t, stringNumber)
}
