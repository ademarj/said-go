package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/ademarj/said-go/db/dao"
	"github.com/ademarj/said-go/db/model"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
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
	URL_SERVICE    = "https://calendarific.com/api/v2/holidays?&api_key=468db849dfcf900f0f47eca41cc6abf0bc5f55d2&country=ZA&year=2019"
)

func indexPage(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles(index_page)
	t.Execute(response, nil)
}

func holidaysInfoPage(response http.ResponseWriter, request *http.Request) {
	redirectPage := root_path
	saNumberId := request.FormValue(req_said)

	if saNumberId != " " {
		contact, _ := model.CreateContactFrom(saNumberId)
		objContact, _ := dao.FindBy(contact.IdNumber)
		ableToCallout := false
		if contact.IdNumber == objContact.IdNumber {
			contact.Counter = contact.Counter + objContact.Counter
			updated, _ := dao.UpdateContact(contact)
			if updated {
				ableToCallout = updated
			}
		} else {
			inserted, _ := dao.SaveNewContact(contact)
			if inserted {
				ableToCallout = inserted
			}
		}

		fmt.Printf("ableToCallout %t\n", ableToCallout)

		var responseAPI model.ResponseRest
		resp, err := http.Get(URL_SERVICE)
		if err != nil {
			panic(err)
		}

		responseBody, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(responseBody, &responseAPI)

		if responseAPI.Meta.Code == 200 {

			result := gjson.Get(string(responseBody), "response.holidays")
			var holidays []model.Holiday
			result.ForEach(func(key, value gjson.Result) bool {
				name := gjson.Get(value.String(), "name")
				holidays = append(holidays, model.Holiday{Name: name.String()})
				println(name.String())
				return true // keep iterating
			})

			rep := model.Response{Holidays: holidays}
			fmt.Println(holidays)

			t, _ := template.ParseFiles(holiday_page)
			t.Execute(response, rep)
			return
		}

		t, _ := template.ParseFiles(holiday_page)
		t.Execute(response, contact)
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
