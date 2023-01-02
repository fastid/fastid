package services

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
	"sync"
)

type Server interface {
	//UnlockDatabase(ctx context.Context, key string) bool
	//CreateSignatureKey(ctx context.Context) (signatureKey string, err error)
	//GetSignatureKey(ctx context.Context) (signatureKey string, err error)
}

type server struct {
	mutex        sync.Mutex
	cfg          *config.Config
	logger       logger.Logger
	repositories repositories.Repositories
	key          string
}

func NewServerService(cfg *config.Config, logger logger.Logger, repositories repositories.Repositories) Server {
	return &server{cfg: cfg, logger: logger, repositories: repositories}
}

//func (s *server) GetSignatureKey(ctx context.Context) (signatureKey string, err error) {
//	signatureKey, err = s.repositories.Keys().GetKey(ctx)
//	if err != nil {
//		return "", err
//	}
//	return signatureKey, nil
//}
//
//func (s *server) CreateSignatureKey(ctx context.Context) (signatureKey string, err error) {
//	s.mutex.Lock()
//	defer s.mutex.Unlock()
//
//	signatureKey, err = s.GetSignatureKey(ctx)
//	if err != nil {
//		return "", err
//	}
//
//	if signatureKey != "" {
//		return "", err
//	}
//
//	signatureKey, err = s.repositories.Keys().CreateKey(ctx)
//	if err != nil {
//		return "", err
//	}
//	return signatureKey, err
//}
//
//func (s *server) UnlockDatabase(ctx context.Context, key string) bool {
//
//	if s.key != "" {
//		s.mutex.Lock()
//		defer s.mutex.Unlock()
//
//		defer s.logger.Infof(ctx, "Skip unlock database (key:%s)", masker.Address(key))
//		return false
//	}
//
//	key, err := s.repositories.Keys().GetKey(ctx)
//	fmt.Println(key)
//
//	if err != nil {
//		return false
//	}
//
//	//fmt.Println(createKey)
//	//s.key = key
//	s.logger.Infof(ctx, "Unlock database (key:%s)", masker.Address(key))
//
//	return true
//}
