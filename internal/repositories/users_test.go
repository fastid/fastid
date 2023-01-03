package repositories

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUser(t *testing.T) {
	teardownSuite, repo, ctx, err := setupSuite(t)
	defer teardownSuite(t)
	require.NoError(t, err)

	userData := UserData{Email: "user@exmaple.com", Password: "password", Username: "user"}

	err = repo.Users().Create(ctx, &userData)
	if err != nil {
		require.NoError(t, err)
	}

	fmt.Println(userData.UserId)

}
