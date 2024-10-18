package line_count_test

import (
	"testing"

	"github.com/Sotatek-HuyDao/golang-practice-101/internal/line_count"
	"github.com/stretchr/testify/assert"
)

func TestLineCount(t *testing.T) {
	svc := line_count.NewService()

	count, _ := svc.LineCount("../../test/myfile.txt")
	assert.Equal(t, 4, count)
}
