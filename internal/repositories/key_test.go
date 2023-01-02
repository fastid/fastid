package repositories

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKey(t *testing.T) {
	teardownSuite, repo, ctx, err := setupSuite(t)
	defer teardownSuite(t)
	require.NoError(t, err)

	//key, err := repo.Keys().GetKey(ctx)
	//if err != nil {
	//	return
	//}
	//fmt.Println(key)

	privateKey, err := repo.Keys().CreateKey(ctx)
	if err != nil {
		return
	}
	require.NotEmpty(t, privateKey)

	repo.Keys().GetKey(ctx, privateKey)

}
