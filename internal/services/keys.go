package services

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/pkg/crypto"
	log "github.com/sirupsen/logrus"
)

type Keys interface {
	GenerateKey() (cipher string, err error)
	Key(ctx context.Context) (err error)
	//SetLogger(logger *log.Entry)
}

type keys struct {
	logger       *log.Logger
	cfg          *config.Config
	repositories repositories.Repositories
}

func NewKeyService(cfg *config.Config, logger *log.Logger, repositories repositories.Repositories) Keys {
	return &keys{cfg: cfg, logger: logger, repositories: repositories}
}

// GenerateKey - Generates a key for encryption
func (k *keys) GenerateKey() (cipher string, err error) {
	cipher, err = crypto.GenerateCipher()
	if err != nil {
		return "", err
	}
	k.logger.Infof("Generate key %s", cipher)
	return cipher, nil
}

// Key - Generates a key for encryption
func (k *keys) Key(ctx context.Context) (err error) {
	k.repositories.Keys().CreateKey(ctx)

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
