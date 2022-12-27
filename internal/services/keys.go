package services

import (
	"context"
	"fmt"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/pkg/crypto"
	"github.com/ggwhite/go-masker"
	log "github.com/sirupsen/logrus"
)

type Keys interface {
	RequestID(requestID interface{}) Keys
	ResetRequestID() Keys
	GenerateKey() (cipher string, err error)
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

func (k *keys) RequestID(requestID interface{}) Keys {
	if requestID != nil {
		k.requestID = fmt.Sprintf("%s", requestID)
	}
	return k
}

func (k *keys) ResetRequestID() Keys {
	k.requestID = nil
	return k
}

// GenerateKey - Generates a key for encryption
func (k *keys) GenerateKey() (cipher string, err error) {

	cipher, err = crypto.GenerateCipher()
	if err != nil {
		return "", err
	}

	logger := k.log.WithField("x-request-id", k.requestID)
	logger.Infof("Generate key %s", masker.Address(cipher))

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
