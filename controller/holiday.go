package controller

import (
	"fmt"

	"github.com/ademarj/said-go/db/dao"
	"github.com/ademarj/said-go/db/model"
	"github.com/ademarj/said-go/http/api"
	"github.com/ademarj/said-go/util/security"
	"github.com/tidwall/gjson"
)

func SearchHolidays(contact model.Contact) model.Holidays {
	success, result := api.Calendarific(contact.IdNumber)

	if success {
		var holidaysFromApi []model.Holiday
		result.ForEach(func(key, value gjson.Result) bool {
			name := gjson.Get(value.String(), "name")
			description := gjson.Get(value.String(), "description")
			date := gjson.Get(value.String(), "date.iso")

			holiday := model.Holiday{
				Id:          fmt.Sprintf("%x", security.GenerateKey([]string{contact.IdNumber, ":", contact.DateOfBirthday, ":", name.String()})),
				Name:        name.String(),
				Description: description.String(),
				Date:        date.String(),
				ContactId:   contact.IdNumber,
			}

			holidaysFromApi = append(holidaysFromApi, holiday)

			return true // keep iterating
		})

		holidaysFromContact, _ := dao.GetHolidaysFrom(contact.IdNumber)

		filtered := filter(holidaysFromContact, holidaysFromApi)

		dao.SaveHoliday(filtered)

		//return append(filtered, holidaysFromContact...)

		holidaysMerged := append(filtered, holidaysFromContact...)
		if len(holidaysMerged) <= 0 {
			return model.Holidays{}
		}

		var holidays []model.Holiday
		columnGrid := 0
		rowGrid := 1
		for _, h := range holidaysMerged {
			columnGrid += 1
			h.GridColumn = columnGrid
			h.GridRow = rowGrid
			holidays = append(holidays, h)
			if columnGrid == 4 {
				columnGrid = 0
				rowGrid += 1
			}
		}
		return holidays
	}

	return model.Holidays{}
}

func filter(holidaysFromContact model.Holidays, holidaysFromApi model.Holidays) model.Holidays {
	if len(holidaysFromApi) > 0 {
		if len(holidaysFromContact) > 0 {
			var filtered model.Holidays
			for _, fromService := range holidaysFromApi {
				contains := false
				for _, holiday := range holidaysFromContact {
					if fromService.Id == holiday.Id {
						contains = true
						break
					}
				}
				if !contains {
					filtered = append(filtered, fromService)
				}
			}
			return filtered
		}
		return holidaysFromApi
	}
	return model.Holidays{}
}
