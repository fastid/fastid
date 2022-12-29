package services

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/pkg/crypto"
	"github.com/ggwhite/go-masker"
)

type Keys interface {
	GenerateKey(ctx context.Context) (cipher string, err error)
	Key(ctx context.Context) (err error)
}

type keys struct {
	cfg          *config.Config
	logger       logger.Logger
	repositories repositories.Repositories
	requestID    interface{}
}

func NewKeyService(cfg *config.Config, logger logger.Logger, repositories repositories.Repositories) Keys {
	return &keys{cfg: cfg, logger: logger, repositories: repositories}
}

// GenerateKey - Generates a key for encryption
func (k *keys) GenerateKey(ctx context.Context) (cipher string, err error) {

	cipher, err = crypto.GenerateCipher()
	if err != nil {
		return "", err
	}

	k.logger.Infof(ctx, "Generate key %s", masker.Address(cipher))
	return cipher, nil
}

// Key - Generates a key for encryption
func (k *keys) Key(ctx context.Context) (err error) {
	_, err = k.repositories.Keys().CreateKey(ctx)
	if err != nil {
		return err
	}

	//key, err := k.repositories.Keys().GetKey(ctx)
	//if err == sql.ErrNoRows {
	//
	//}
	//
	//if err != nil {
	//	k.log.Errorln(err.Error())
	//	return err
	//}
	//
	//fmt.Println(key)
	return err
}
