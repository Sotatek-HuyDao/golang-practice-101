package version

import (
	"fmt"

	"github.com/Sotatek-HuyDao/golang-practice-101/utils"
	"github.com/sirupsen/logrus"
)

var version = "0.0.1"

type Service interface {
	Version() (string, error)
}

type ServiceImpl struct {
	version string
	logger  *logrus.Logger
}

func NewService() *ServiceImpl {
	svc := &ServiceImpl{
		version: version,
		logger:  utils.Logger(),
	}
	return svc
}

func (s *ServiceImpl) Version() (string, error) {
	return fmt.Sprintf("futil v%s", s.version), nil
}
