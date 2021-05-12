package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	code_320       = 320
	index_page     = "public/index.html"
	holiday_page   = "public/holiday.html"
	root_path      = "/"
	resource_dir   = "/static/"
	req_said       = "southAfricaNumberId"
	holiday_action = "/holiday"
	http_port      = ":9000"
	method_post    = "POST"
)

func indexPage(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles(index_page)
	t.Execute(response, nil)
}

type SouthObject struct {
	SouthAfricaNumberId string
}

func holidaysInfoPage(response http.ResponseWriter, request *http.Request) {
	redirectPage := root_path
	saNumberId := request.FormValue(req_said)
	fmt.Println(saNumberId)

	if saNumberId != " " {
		p := SouthObject{SouthAfricaNumberId: saNumberId}
		t, _ := template.ParseFiles(holiday_page)
		t.Execute(response, p)
	} else {
		http.Redirect(response, request, redirectPage, code_320)
	}
}

var router = InitRouter()

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	resourceDIR := resource_dir

	router.
		PathPrefix(resourceDIR).
		Handler(http.StripPrefix(resourceDIR, http.FileServer(http.Dir("."+resourceDIR))))

	return router
}

func main() {
	router.HandleFunc(root_path, indexPage)
	router.HandleFunc(holiday_action, holidaysInfoPage).Methods(method_post)

	http.Handle(root_path, router)
	http.ListenAndServe(http_port, nil)
}
