package repositories

import (
	"context"
	"database/sql"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/pkg/crypto"
	log "github.com/sirupsen/logrus"
)

type Keys interface {
	GetKey(ctx context.Context) (keysSchema KeysSchema, err error)
	CreateKey(ctx context.Context) (keysSchema KeysSchema, err error)
}

type keys struct {
	cfg *config.Config
	log *log.Logger
	db  db.DB
}

func NewKeysRepository(cfg *config.Config, log *log.Logger, db db.DB) Keys {
	return &keys{cfg: cfg, log: log, db: db}
}

func (k *keys) GetKey(ctx context.Context) (keysSchema KeysSchema, err error) {
	connect, err := k.db.GetConnect().Acquire(ctx)
	if err != nil {
		return keysSchema, err
	}
	defer connect.Release()

	query := `SELECT unpacking_key, signature_key  FROM keys LIMIT 1`
	k.log.Traceln("SQL select keys: " + query)

	err = connect.QueryRow(ctx, query).Scan(
		&keysSchema.UnpackingKey,
		&keysSchema.SignatureKey,
	)

	if err == sql.ErrNoRows {
		return keysSchema, err
	}

	if err != nil {
		return keysSchema, err
	}
	return keysSchema, nil
}

func (k *keys) CreateKey(ctx context.Context) (keysSchema KeysSchema, err error) {
	connect, err := k.db.GetConnect().Acquire(ctx)
	if err != nil {
		return keysSchema, err
	}
	defer connect.Release()

	cipher, err := crypto.GenerateCipher()
	if err != nil {
		return keysSchema, err
	}

	cipherPrivate, err := crypto.GenerateCipher()
	if err != nil {
		return keysSchema, err
	}

	cr := crypto.New(cipher)
	encrypt, err := cr.Encrypt(cipherPrivate)
	if err != nil {
		return KeysSchema{}, err
	}

	_, err = connect.Exec(ctx, "INSERT INTO keys (unpacking_key, signature_key) VALUES ($1, $2)", encrypt, encrypt)
	if err != nil {
		return KeysSchema{}, err
	}
	return KeysSchema{}, nil
}
