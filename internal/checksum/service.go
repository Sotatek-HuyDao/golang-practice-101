package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/Sotatek-HuyDao/golang-practice-101/utils"
	"github.com/sirupsen/logrus"
)

type Service interface {
	MD5Checksum(filePath string) (string, error)
	SHA1Checksum(filePath string) (string, error)
	SHA256Checksum(filePath string) (string, error)
}

type ServiceImpl struct {
	logger *logrus.Logger
}

func NewService() *ServiceImpl {
	svc := &ServiceImpl{
		logger: utils.Logger(),
	}
	return svc
}

func (s *ServiceImpl) MD5Checksum(filePath string) (string, error) {
	data, err := readFile(filePath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(data)), nil
}

func (s *ServiceImpl) SHA1Checksum(filePath string) (string, error) {
	data, err := readFile(filePath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha1.Sum(data)), nil
}

func (s *ServiceImpl) SHA256Checksum(filePath string) (string, error) {
	data, err := readFile(filePath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(data)), nil
}

func readFile(filePath string) ([]byte, error) {
	file, err := utils.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer utils.CloseFile(file)
	return io.ReadAll(file)
}
