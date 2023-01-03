package services

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/google/uuid"
	"sync"
)

type UserData struct {
	Username interface{}
	Email    string
	Password string
	UserId   uuid.UUID
}

type Users interface {
	Create(ctx context.Context, user *UserData) (err error)
	SetActive(ctx context.Context, userID *uuid.UUID, isActive bool) (err error)
	SetSuperUser(ctx context.Context, userID *uuid.UUID, isSuperUser bool) (err error)
}

type users struct {
	mutex        sync.Mutex
	cfg          *config.Config
	logger       logger.Logger
	repositories repositories.Repositories
}

func NewUsersService(cfg *config.Config, logger logger.Logger, repositories repositories.Repositories) Users {
	return &users{cfg: cfg, logger: logger, repositories: repositories}
}

// Create a new user
func (u *users) Create(ctx context.Context, userData *UserData) (err error) {

	repoUserData := repositories.UserData{
		Username: userData.Username,
		Email:    userData.Email,
		Password: userData.Password,
	}

	err = u.repositories.Users().Create(ctx, &repoUserData)
	if err != nil {
		return err
	}

	userData.UserId = repoUserData.UserId
	u.logger.Infof(ctx, "Create a new user %s", userData.UserId)
	return nil
}

// SetActive - Sets the value user activity
func (u *users) SetActive(ctx context.Context, userID *uuid.UUID, isActive bool) (err error) {
	err = u.repositories.Users().SetActive(ctx, userID, isActive)
	if err != nil {
		return err
	}

	u.logger.Infof(ctx, "Set flag is_active %t for user_id %s", isActive, *userID)
	return nil
}

// SetSuperUser - Sets the superuser value
func (u *users) SetSuperUser(ctx context.Context, userID *uuid.UUID, isSuperUser bool) (err error) {
	err = u.repositories.Users().SetSuperUser(ctx, userID, isSuperUser)
	if err != nil {
		return err
	}
	u.logger.Infof(ctx, "Set flag is_superuser %t for user_id %s", isSuperUser, *userID)
	return nil
}
