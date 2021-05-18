package handler

import (
	"net/http"
	"text/template"

	"github.com/ademarj/said-go/src/controller"
)

const (
	INDEX_PAGE        = "web/page/index.html"
	HOLIDAY_PAGE      = "web/page/holiday.html"
	HOLIDAYS_PAGE     = "web/page/holidays.html"
	REQUEST_NUMBER_ID = "southAfricaNumberId"
)

func Index(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles(INDEX_PAGE)
	t.Execute(response, nil)
}

func HolidayPage(response http.ResponseWriter, request *http.Request) {
	numberId := request.FormValue(REQUEST_NUMBER_ID)
	view, _ := controller.BuildViewHoliday(numberId)
	renderPage := HOLIDAY_PAGE
	if len(view.Holidays) > 1 {
		renderPage = HOLIDAYS_PAGE
	}
	t, _ := template.ParseFiles(renderPage)
	t.Execute(response, view)
}
