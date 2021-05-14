package security

import (
	"crypto/md5"
	"io"
)

func GenerateKey(particle []string) []byte {
	h := md5.New()
	for _, p := range particle {
		io.WriteString(h, p)
	}
	return h.Sum(nil)
}
