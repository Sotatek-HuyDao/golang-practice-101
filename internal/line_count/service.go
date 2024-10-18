package line_count

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"

	"github.com/Sotatek-HuyDao/golang-practice-101/utils"
	"github.com/sirupsen/logrus"
)

type Service interface {
	LineCount(filePath string) (int, error)
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

func (s *ServiceImpl) LineCount(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {

		return 0, fmt.Errorf("No such file '%s'", filePath)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	if fileInfo.IsDir() {
		return 0, fmt.Errorf("Expected file got directory '%s'", filePath)
	}

	if isBinary(file) {
		return 0, fmt.Errorf("Cannot do linecount for binary file '%s'\n", filePath)

	}
	file.Seek(0, io.SeekStart)

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
func isBinary(file *os.File) bool {
	buf := make([]byte, 512)
	n, err := file.Read(buf)
	if err != nil {
		return false
	}

	// Check for non-printable characters
	for i := 0; i < n; i++ {
		if !unicode.IsPrint(rune(buf[i])) && !unicode.IsSpace(rune(buf[i])) {
			return true
		}
	}
	return false
}
