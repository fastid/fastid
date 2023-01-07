package repositories

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
	"github.com/google/uuid"
	"sync"
)

type UserData struct {
	Username  interface{} `db:"username"`
	Email     string      `db:"email"`
	Password  string      `db:"password"`
	UserId    uuid.UUID   `db:"user_id"`
	Active    bool        `db:"is_active"`
	SuperUser bool        `db:"is_superuser"`
}

type Users interface {
	Create(ctx context.Context, user *UserData) (err error)
	SetActive(ctx context.Context, userID *uuid.UUID, isActive bool) (err error)
	SetSuperUser(ctx context.Context, userID *uuid.UUID, isSuperUser bool) (err error)
	GetByEmail(ctx context.Context, email string) (userData UserData, err error)
	GetByUsername(ctx context.Context, username string) (userData UserData, err error)
}

type users struct {
	cfg    *config.Config
	logger logger.Logger
	db     db.DB
	mutex  sync.Mutex
}

func NewUsersRepository(cfg *config.Config, logger logger.Logger, db db.DB) Users {
	return &users{cfg: cfg, logger: logger, db: db}
}

func (u *users) Create(ctx context.Context, userData *UserData) (err error) {

	u.mutex.Lock()
	defer u.mutex.Unlock()

	connect, err := u.db.GetConnect().Acquire(ctx)
	if err != nil {
		return err
	}
	defer connect.Release()

	userData.UserId = uuid.New()

	_, err = connect.Exec(
		ctx,
		"INSERT INTO users (user_id, username, email, password, is_active, is_superuser) VALUES ($1, $2, $3, $4, $5, $6)",
		userData.UserId,
		userData.Username,
		userData.Email,
		userData.Password,
		userData.Active,
		userData.SuperUser,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *users) SetActive(ctx context.Context, userID *uuid.UUID, isActive bool) (err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	connect, err := u.db.GetConnect().Acquire(ctx)
	if err != nil {
		return err
	}
	defer connect.Release()

	_, err = connect.Exec(ctx, "UPDATE users SET is_active = $1 WHERE user_id = $2", isActive, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *users) SetSuperUser(ctx context.Context, userID *uuid.UUID, isSuperUser bool) (err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	connect, err := u.db.GetConnect().Acquire(ctx)
	if err != nil {
		return err
	}
	defer connect.Release()

	_, err = connect.Exec(ctx, "UPDATE users SET is_superuser = $1 WHERE user_id = $2", isSuperUser, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *users) GetByEmail(ctx context.Context, email string) (userData UserData, err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	connect, err := u.db.GetConnect().Acquire(ctx)
	if err != nil {
		return UserData{}, nil
	}
	defer connect.Release()

	err = connect.QueryRow(
		ctx,
		"SELECT user_id, username, email, password, is_active, is_superuser FROM users WHERE email = $1",
		email,
	).Scan(
		&userData.UserId,
		&userData.Username,
		&userData.Email,
		&userData.Password,
		&userData.Active,
		&userData.SuperUser,
	)

	if err != nil {
		return UserData{}, err
	}
	return userData, nil
}

func (u *users) GetByUsername(ctx context.Context, username string) (userData UserData, err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	connect, err := u.db.GetConnect().Acquire(ctx)
	if err != nil {
		return UserData{}, nil
	}
	defer connect.Release()

	err = connect.QueryRow(
		ctx,
		"SELECT user_id, username, email, password, is_active, is_superuser FROM users WHERE username = $1",
		username,
	).Scan(
		&userData.UserId,
		&userData.Username,
		&userData.Email,
		&userData.Password,
		&userData.Active,
		&userData.SuperUser,
	)

	if err != nil {
		return UserData{}, err
	}
	return userData, nil
}
