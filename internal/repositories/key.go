package repositories

import (
	"context"
	"fmt"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/pkg/crypto"
	"github.com/jackc/pgx/v5"
	"sync"
)

type Keys interface {
	GetKey(ctx context.Context, privateKey string) (signatureKey string, err error)
	CreateKey(ctx context.Context) (privateKey string, err error)
}

type keys struct {
	cfg    *config.Config
	logger logger.Logger
	db     db.DB
	mutex  sync.Mutex
}

func NewKeysRepository(cfg *config.Config, logger logger.Logger, db db.DB) Keys {
	return &keys{cfg: cfg, logger: logger, db: db}
}

func (k *keys) GetKey(ctx context.Context, privateKey string) (signatureKey string, err error) {
	connect, err := k.db.GetConnect().Acquire(ctx)
	if err != nil {
		return "", err
	}
	defer connect.Release()

	query := `SELECT signature_key FROM key LIMIT 1`
	k.logger.Tracef(ctx, "SQL select keys: %s", query)

	var signatureKeyByte []byte
	err = connect.QueryRow(ctx, query).Scan(&signatureKeyByte)

	crpt := crypto.New(privateKey)
	decrypt, err := crpt.Decrypt(signatureKeyByte)
	if err != nil {
		return "", err
	}
	fmt.Println(decrypt)

	if err == pgx.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return signatureKey, nil
}

func (k *keys) CreateKey(ctx context.Context) (privateKey string, err error) {
	k.mutex.Lock()
	defer k.mutex.Unlock()

	connect, err := k.db.GetConnect().Acquire(ctx)
	if err != nil {
		return "", err
	}
	defer connect.Release()

	// Private key
	privateKey, err = crypto.GenerateCipher()
	if err != nil {
		return "", err
	}

	// Generating a key to encrypt the database
	cipher, err := crypto.GenerateCipher()
	if err != nil {
		return "", err
	}

	// Encrypting the key
	cr := crypto.New(privateKey)
	signatureKey, err := cr.Encrypt(cipher)
	if err != nil {
		return "", err
	}

	_, err = connect.Exec(ctx, "INSERT INTO key (signature_key) VALUES ($1)", signatureKey)
	if err != nil {
		return "", err
	}
	return privateKey, nil
}
