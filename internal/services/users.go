package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/pbkdf2"
	"strings"
	"sync"
)

type UserData struct {
	Username  interface{}
	Email     string
	Password  string
	UserId    uuid.UUID
	Active    bool
	SuperUser bool
}

type Users interface {
	Create(ctx context.Context, user *UserData) (err error)
	SetActive(ctx context.Context, userID *uuid.UUID, isActive bool) (err error)
	SetSuperUser(ctx context.Context, userID *uuid.UUID, isSuperUser bool) (err error)
	HashPassword(password string) string
	GetByEmail(ctx context.Context, email string) (userData UserData, err error)
	GetByUsername(ctx context.Context, username string) (userData UserData, err error)
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
		Password: u.HashPassword(userData.Password),
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

// HashPassword ...
func (u *users) HashPassword(password string) string {
	password = strings.TrimRight(password, " ")
	password = strings.TrimLeft(password, " ")
	passwordByte := pbkdf2.Key([]byte(password), []byte(u.cfg.APP.Salt), 390000, sha256.Size, sha256.New)
	return hex.EncodeToString(passwordByte)
}

// GetByEmail - Get a user by email
func (u *users) GetByEmail(ctx context.Context, email string) (userData UserData, err error) {
	email = strings.TrimRight(email, " ")
	email = strings.TrimLeft(email, " ")
	email = strings.ToLower(email)

	getByEmail, err := u.repositories.Users().GetByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return UserData{}, nil
		}
		return UserData{}, err
	}

	u.logger.Infof(ctx, "Get a user by email %s", email)

	return UserData{
		UserId:    getByEmail.UserId,
		Username:  getByEmail.Username,
		Email:     getByEmail.Email,
		Password:  getByEmail.Password,
		Active:    getByEmail.Active,
		SuperUser: getByEmail.SuperUser,
	}, nil
}

func (u *users) GetByUsername(ctx context.Context, username string) (userData UserData, err error) {
	username = strings.TrimRight(username, " ")
	username = strings.TrimLeft(username, " ")
	username = strings.ToLower(username)

	getByUsername, err := u.repositories.Users().GetByUsername(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return UserData{}, nil
		}
		return UserData{}, err
	}

	u.logger.Infof(ctx, "Get a user by username %s", username)

	return UserData{
		UserId:    getByUsername.UserId,
		Username:  getByUsername.Username,
		Email:     getByUsername.Email,
		Password:  getByUsername.Password,
		Active:    getByUsername.Active,
		SuperUser: getByUsername.SuperUser,
	}, nil

}
