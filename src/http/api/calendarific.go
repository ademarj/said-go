package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ademarj/said-go/src/helper"
	jmodel "github.com/ademarj/said-go/src/http/json"
	"github.com/ademarj/said-go/src/said"
	"github.com/ademarj/said-go/src/util/security"
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/tidwall/gjson"
)

const (
	/*
		[ HOLIDAY_TYPE ]
			national 	- Returns public, federal and bank holidays
			local 		- Returns local, regional and state holidays
			religious 	- Return religious holidays:
							buddhism, christian, hinduism, muslim, etc
			observance 	- Observance, Seasons, Times
	*/

	HOLIDAY_TYPE              = "national" // "" for all holidays type
	ALL_HOLIDAYS_CURRENT_YEAR = false
	URL_SERVICE               = "URL_CALENDARIFIC_REST_API"
	API_KEY                   = "API_KEY_CALENDARIFIC_REST_API"
	COUNTRY                   = "ZA"
	SUCCESS                   = "success"
	ERROR_RESULT              = "{success: false}"
	URL_API_COUNTRY           = "%s&api_key=%s&country=%s"
	URL_CURRENT_YEAR          = "%s&year=%d"
	URL_HOLIDAY_TYPE          = "%s&type=%s"
	URL_BIRTHDAY              = "%s&year=%d&month=%d&day=%d"
	NODE_HOLIDAYS             = "response.holidays"
)

func Calendarific(numberId string) (success bool, jsonResult gjson.Result) {
	resp, err := http.Get(makeUrlRequest(numberId))
	if err != nil {
		return false, gjson.Get(ERROR_RESULT, SUCCESS)
	}
	var response jmodel.ResponseRest
	responseBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseBody, &response)

	if response.Meta.Code == 200 {
		return true, gjson.Get(string(responseBody), NODE_HOLIDAYS)
	}

	return false, gjson.Get(ERROR_RESULT, SUCCESS)
}

func makeUrlRequest(numberId string) string {
	url, _ := security.GetSecret(URL_SERVICE)
	apikey, _ := security.GetSecret(API_KEY)
	url = fmt.Sprintf(URL_API_COUNTRY, url, apikey, COUNTRY)

	if stringUtils.IsNotBlank(HOLIDAY_TYPE) {
		url = fmt.Sprintf(URL_HOLIDAY_TYPE, url, HOLIDAY_TYPE)
	}

	if ALL_HOLIDAYS_CURRENT_YEAR {
		return fmt.Sprintf(URL_CURRENT_YEAR, url, helper.CurrentYear())
	}
	return fmt.Sprintf(URL_BIRTHDAY, url, helper.CurrentYear(), said.Month(numberId), said.Day(numberId))
}
