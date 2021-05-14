package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ademarj/said-go/helper"
	jmodel "github.com/ademarj/said-go/http/json"
	"github.com/ademarj/said-go/util/said"
	"github.com/tidwall/gjson"
)

const (
	URL_SERVICE = "https://calendarific.com/api/v2/holidays?&api_key=468db849dfcf900f0f47eca41cc6abf0bc5f55d2&country=ZA"
)

func Calendarific(numberId string) (success bool, jsonResult gjson.Result) {
	reqURL := fmt.Sprintf("%s&year=%d&month=%d&day=%d", URL_SERVICE, helper.CurrentYear(), said.Month(numberId), said.Day(numberId))

	fmt.Println(reqURL)

	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}

	var responseAPI jmodel.ResponseRest
	responseBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseBody, &responseAPI)

	if responseAPI.Meta.Code == 200 {

		fmt.Println(string(responseBody))

		result := gjson.Get(string(responseBody), "response.holidays")

		return true, result

	}

	return false, gjson.Get("{success: false}", "success")
}
