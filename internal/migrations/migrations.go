package migrations

import (
	"database/sql"
	"embed"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed sql/*.sql
var schemaFs embed.FS

type Migration interface {
	Upgrade() error
	Downgrade() error
	Drop() error
}

type migration struct {
	cfg    *config.Config
	driver *database.Driver
}

func New(cfg *config.Config, db db.DB) (Migration, error) {
	sqlDb, err := sql.Open("postgres", db.Dsn())
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(sqlDb, &postgres.Config{})
	return &migration{cfg: cfg, driver: &driver}, nil
}

func (m *migration) Upgrade() error {

	dirs, err := iofs.New(schemaFs, "sql")
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithInstance("iofs", dirs, m.cfg.DBName, *m.driver)
	if err != nil {
		return err
	}

	if err := mig.Up(); err != nil {
		return err
	}
	return nil
}

func (m *migration) Downgrade() error {
	dirs, err := iofs.New(schemaFs, "sql")
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithInstance("iofs", dirs, m.cfg.DBName, *m.driver)
	if err != nil {
		return err
	}

	if err := mig.Down(); err != nil {
		return err
	}
	return nil
}

func (m *migration) Drop() error {
	dirs, err := iofs.New(schemaFs, "sql")
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithInstance("iofs", dirs, m.cfg.DBName, *m.driver)
	if err != nil {
		return err
	}

	if err := mig.Drop(); err != nil {
		return err
	}
	return nil
}
