package db

import (
	"context"
	"fmt"
	"github.com/fastid/fastid/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type DB interface {
	GetConnect() *pgxpool.Pool
	GetAcquire(ctx context.Context) (*pgxpool.Conn, error)
	Dsn() string
	CreateScheme(ctx context.Context, schema string) error
	DropScheme(ctx context.Context, schema string) error
}

type database struct {
	cfg *config.Config
	db  *pgxpool.Pool
	dsn string
}

func New(cfg *config.Config, ctx context.Context) (DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s application_name=%s sslmode=%s search_path=%s",
		cfg.DATABASE.Host,
		cfg.DATABASE.Port,
		cfg.DATABASE.User,
		cfg.DATABASE.Password,
		cfg.DATABASE.DBName,
		cfg.DATABASE.ApplicationName,
		cfg.DATABASE.SslMode,
		cfg.DATABASE.Scheme,
	)

	configPoll, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(ctx, configPoll)
	if err != nil {
		return nil, err
	}
	return &database{db: db, cfg: cfg, dsn: dsn}, nil
}

func (d *database) GetConnect() *pgxpool.Pool {
	return d.db
}

func (d *database) GetAcquire(ctx context.Context) (*pgxpool.Conn, error) {
	return d.db.Acquire(ctx)
}

func (d *database) Dsn() string {
	return d.dsn
}

func (d *database) CreateScheme(ctx context.Context, schema string) error {
	_, err := d.db.Exec(ctx, "CREATE SCHEMA IF NOT EXISTS "+schema)
	if err != nil {
		return err
	}
	return nil
}

func (d *database) DropScheme(ctx context.Context, schema string) error {
	_, err := d.db.Exec(ctx, "DROP SCHEMA IF EXISTS "+schema+" CASCADE")
	if err != nil {
		return err
	}
	return nil
}
