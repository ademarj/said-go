package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ademarj/said-go/helper"
	jmodel "github.com/ademarj/said-go/http/json"
	"github.com/ademarj/said-go/said"
	"github.com/ademarj/said-go/util/security"
	"github.com/tidwall/gjson"
)

const (
	ALL_HOLIDAYS_CURRENT_YEAR = false
	URL_SERVICE               = "URL_CALENDARIFIC_REST_API"
	API_KEY                   = "API_KEY_CALENDARIFIC_REST_API"
	COUNTRY                   = "ZA"
	SUCCESS                   = "success"
	ERROR_RESULT              = "{success: false}"
	URL_API_COUNTRY           = "%s&api_key=%s&country=%s"
	URL_CURRENT_YEAR          = "%s&year=%d"
	URL_BIRTHDAY              = "%s&year=%d&month=%d&day=%d"
	NODE_HOLIDAYS             = "response.holidays"
)

func Calendarific(numberId string) (success bool, jsonResult gjson.Result) {
	resp, err := http.Get(makeUrlRequest(numberId))
	if err != nil {
		log.Printf("ERRO MAKE REQUEST %s", err.Error())
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
	if ALL_HOLIDAYS_CURRENT_YEAR {
		return fmt.Sprintf(URL_CURRENT_YEAR, url, helper.CurrentYear())
	}
	return fmt.Sprintf(URL_BIRTHDAY, url, helper.CurrentYear(), said.Month(numberId), said.Day(numberId))
}
