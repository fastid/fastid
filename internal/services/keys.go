package services

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/pkg/crypto"
	"github.com/ggwhite/go-masker"
	log "github.com/sirupsen/logrus"
)

type Keys interface {
	GenerateKey(ctx context.Context) (cipher string, err error)
	Key(ctx context.Context) (err error)

	//SetLogger(logger *log.Entry)
}

type keys struct {
	log          *log.Logger
	cfg          *config.Config
	repositories repositories.Repositories
	requestID    interface{}
}

func NewKeyService(cfg *config.Config, logger *log.Logger, repositories repositories.Repositories) Keys {
	return &keys{cfg: cfg, log: logger, repositories: repositories}
}

// GenerateKey - Generates a key for encryption
func (k *keys) GenerateKey(ctx context.Context) (cipher string, err error) {

	cipher, err = crypto.GenerateCipher()
	if err != nil {
		return "", err
	}

	if ctx.Value("requestID") != nil {
		logger := k.log.WithField("x-request-id", ctx.Value("requestID").(string))
		logger.Infof("Generate key %s", masker.Address(cipher))
	} else {
		k.log.Infof("Generate key %s", masker.Address(cipher))
	}

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
