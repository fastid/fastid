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

	//Create(ctx context.Context, username interface{}, email string, password string)
	//UnlockDatabase(ctx context.Context, key string) bool
	//CreateSignatureKey(ctx context.Context) (signatureKey string, err error)
	//GetSignatureKey(ctx context.Context) (signatureKey string, err error)
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
