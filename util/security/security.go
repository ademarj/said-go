package security

import (
	"crypto/md5"
	"io"
)

func GenerateKey(particule []string) []byte {
	h := md5.New()
	for _, p := range particule {
		io.WriteString(h, p)
	}
	return h.Sum(nil)
}
