package version_test

import (
	"os"
	"testing"

	"github.com/Sotatek-HuyDao/golang-practice-101/internal/version"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	os.Setenv("VERSION", "0.0.1")
	svc := version.NewService()

	version, _ := svc.Version()
	assert.Equal(t, "futil v0.0.1", version)
}
