package security

import (
	"crypto/md5"
	"io"
	"os"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func GenerateKey(particle []string) []byte {
	h := md5.New()
	for _, p := range particle {
		io.WriteString(h, p)
	}
	return h.Sum(nil)
}

func GetSecret(key string) (string, bool) {
	var m = map[string]string{
		"URL_CALENDARIFIC_REST_API":     "URL_CALENDARIFIC_REST_API",
		"API_KEY_CALENDARIFIC_REST_API": "API_KEY_CALENDARIFIC_REST_API",
		"URL_DATABASE_SAID":             "URL_DATABASE_SAID",
		"PORT_DATA_BASE":                "PORT_DATA_BASE",
		"USER_DATABASE_SAID":            "USER_DATABASE_SAID",
		"PASSWORD_DATABASE_SAID":        "PASSWORD_DATABASE_SAID",
		"NAME_DATABASE":                 "NAME_DATABASE",
	}
	result, contain := m[key]
	if contain {
		value := os.Getenv(result)
		if stringUtils.IsBlank(value) {
			return value, false
		}
		return value, true
	}
	return "", false
}
