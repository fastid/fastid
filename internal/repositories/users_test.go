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

	t.Run("Create user", func(t *testing.T) {
		userData := UserData{
			Email:     "user@exmaple.com",
			Password:  "password",
			Username:  "user",
			Active:    true,
			SuperUser: true,
		}
		fmt.Println(userData.UserId)

		err = repo.Users().Create(ctx, &userData)
		require.NoError(t, err)
		require.NotEmpty(t, userData.UserId)
		require.Equal(t, userData.Active, true)
		require.Equal(t, userData.SuperUser, true)
	})

	t.Run("Get by email", func(t *testing.T) {
		userData := UserData{
			Email:    "user@exmaple.com",
			Password: "password",
			Username: "user",
		}
		err = repo.Users().Create(ctx, &userData)
		require.NoError(t, err)

		user, err := repo.Users().GetByEmail(ctx, "user@exmaple.com")
		if err != nil {
			return
		}
		require.Equal(t, user.Username, "user")
		require.Equal(t, user.Email, "user@exmaple.com")
		fmt.Println(user.Password)
		require.NotEmpty(t, user.Password)
	})

	//err = repo.Users().Create(ctx, &userData)
	//if err != nil {
	//	require.NoError(t, err)
	//}
	//
	//// .GetByEmail(ctx, "user@exmaple.com")
	//userDataEmail, err := repo.Users().GetByEmail(ctx, "user@exmaple.com")
	//if err != nil {
	//	require.NoError(t, err)
	//}
	//
	//fmt.Println(userDataEmail)
	//fmt.Println(userData.UserId)

}
