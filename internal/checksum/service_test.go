package checksum_test

import (
	"testing"

	"github.com/Sotatek-HuyDao/golang-practice-101/internal/checksum"
	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	svc := checksum.NewService()

	md5Checksum, _ := svc.MD5Checksum("../../test/myfile.txt")
	assert.Equal(t, "a8c5d553ed101646036a811772ffbdd8", md5Checksum)

	sha1Checksum, _ := svc.SHA1Checksum("../../test/myfile.txt")
	assert.Equal(t, "a656582ca3143a5f48718f4a15e7df018d286521", sha1Checksum)

	sha256Checksum, _ := svc.SHA256Checksum("../../test/myfile.txt")
	assert.Equal(t, "495a3496cfd90e68a53b5e3ff4f9833b431fe996298f5a28228240ee2a25c09d", sha256Checksum)
}
