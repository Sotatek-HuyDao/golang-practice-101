package version

import (
	"fmt"

	config "github.com/Sotatek-HuyDao/golang-practice-101/internal/configs"
	"github.com/Sotatek-HuyDao/golang-practice-101/utils"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Version() (string, error)
}

type ServiceImpl struct {
	version string
	logger  *logrus.Logger
}

func NewService() *ServiceImpl {
	cfg := config.MustParseConfig()
	svc := &ServiceImpl{
		version: cfg.Version,
		logger:  utils.Logger(),
	}
	return svc
}

func (s *ServiceImpl) Version() (string, error) {
	return fmt.Sprintf("futil v%s", s.version), nil
}
