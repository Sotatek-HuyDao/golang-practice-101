package utils

import (
	"fmt"
	"os"
)

func CloseFile(file *os.File) error {
	return file.Close()
}

func OpenFile(filePath string) (*os.File, error) {
	// read from stdin if filePath equal -
	if filePath == "-" {
		return os.Stdin, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("No such file '%s'", filePath)
	}

	// check if filePath is a directory
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("Expected file got directory '%s'", filePath)
	}

	return file, nil
}
