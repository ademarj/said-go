package main

import (
	"html/template"
	"net/http"

	"github.com/ademarj/said-go/controller"
	"github.com/gorilla/mux"
)

const (
	INDEX_PAGE        = "public/index.html"
	HOLIDAY_PAGE      = "public/holiday.html"
	HOLIDAYS_PAGE     = "public/holidays.html"
	ROOT_PATH         = "/"
	RESOURCE_DIR      = "/static/"
	REQUEST_NUMBER_ID = "southAfricaNumberId"
	HOLIDAY_ACTION    = "/holiday"
	HTTP_PORT         = ":9000"
	POST              = "POST"
)

func indexPage(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles(INDEX_PAGE)
	t.Execute(response, nil)
}

func holidayPage(response http.ResponseWriter, request *http.Request) {
	numberId := request.FormValue(REQUEST_NUMBER_ID)
	view, _ := controller.BuildViewHoliday(numberId)
	renderPage := HOLIDAY_PAGE
	if len(view.Holidays) > 1 {
		renderPage = HOLIDAYS_PAGE
	}
	t, _ := template.ParseFiles(renderPage)
	t.Execute(response, view)
}

var router = InitRouter()

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	resourceDIR := RESOURCE_DIR
	router.
		PathPrefix(resourceDIR).
		Handler(http.StripPrefix(resourceDIR, http.FileServer(http.Dir("."+resourceDIR))))
	return router
}

func main() {
	router.HandleFunc(ROOT_PATH, indexPage)
	router.HandleFunc(HOLIDAY_ACTION, holidayPage).Methods(POST)
	http.Handle(ROOT_PATH, router)
	http.ListenAndServe(HTTP_PORT, nil)
}
