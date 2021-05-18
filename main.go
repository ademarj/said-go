package main

import (
	"net/http"

	"github.com/ademarj/said-go/src/handler"
	"github.com/gorilla/mux"
)

const (
	INDEX     = "/"
	HOLIDAY   = "/holiday"
	RESOURCE  = "/web/src/"
	HTTP_PORT = ":9000"
	POST      = "POST"
)

var router = InitRouter()

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	resourceDIR := RESOURCE
	router.
		PathPrefix(resourceDIR).
		Handler(http.StripPrefix(resourceDIR, http.FileServer(http.Dir("."+resourceDIR))))
	return router
}

func main() {
	router.HandleFunc(INDEX, handler.Index)
	router.HandleFunc(HOLIDAY, handler.HolidayPage).Methods(POST)
	http.Handle(INDEX, router)
	http.ListenAndServe(HTTP_PORT, nil)
}
