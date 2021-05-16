package api

import (
	"net/http"

	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/gorilla/securecookie"
)

const (
	NUMBER_ID    = "numberId"
	SESSION      = "session"
	PATH         = "/"
	EMPTY_STRING = ""
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func SetSession(numberId string, response http.ResponseWriter) {
	value := map[string]string{
		NUMBER_ID: numberId,
	}
	if encoded, err := cookieHandler.Encode(SESSION, value); err == nil {
		cookie := &http.Cookie{
			Name:  SESSION,
			Value: encoded,
			Path:  PATH,
		}
		http.SetCookie(response, cookie)
	}
}

func GetSession(request *http.Request) (numberId string, hasContent bool) {
	if cookie, err := request.Cookie(SESSION); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode(SESSION, cookie.Value, &cookieValue); err == nil {
			numberId = cookieValue[NUMBER_ID]
		}
	}
	return numberId, stringUtils.IsNotBlank(numberId)
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   SESSION,
		Value:  EMPTY_STRING,
		Path:   PATH,
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
