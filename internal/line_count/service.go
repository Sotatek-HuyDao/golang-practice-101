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
	reader := os.Stdin
	if filePath != "-" {
		file, err := utils.OpenFile(filePath)
		if err != nil {
			return 0, err
		}
		defer utils.CloseFile(file)
		if isBinary(file) {
			return 0, fmt.Errorf("Cannot do linecount for binary file '%s'\n", filePath)
		}
		file.Seek(0, io.SeekStart)
		reader = file
	}

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := reader.Read(buf)
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
