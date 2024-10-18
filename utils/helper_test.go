package utils_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Sotatek-HuyDao/golang-practice-101/utils"
	"github.com/stretchr/testify/assert"
)

type OpenFileTestcases struct {
	filePath string
	err      error
}

func TestOpenFile(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	tests := []struct {
		name string
		info *OpenFileTestcases
	}{
		{
			name: "open success",
			info: &OpenFileTestcases{
				filePath: tmpfile.Name(),
				err:      nil,
			},
		},
		{
			name: "open on existent file",
			info: &OpenFileTestcases{
				filePath: "non-exist-file",
				err:      fmt.Errorf("No such file 'non-exist-file'"),
			},
		},
		{
			name: "open folder",
			info: &OpenFileTestcases{
				filePath: ".",
				err:      fmt.Errorf("Expected file got directory '.'"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			file, err := utils.OpenFile(tc.info.filePath)
			assert.Equal(t, err, tc.info.err)
			defer utils.CloseFile(file)
		})
	}
}
