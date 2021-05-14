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
		var holidays []model.Holiday
		result.ForEach(func(key, value gjson.Result) bool {
			name := gjson.Get(value.String(), "name")
			description := gjson.Get(value.String(), "description")
			date := gjson.Get(value.String(), "date.iso")

			fmt.Println(date.String())
			holiday := model.Holiday{
				Id:          fmt.Sprintf("%x", security.GenerateKey([]string{contact.IdNumber, ":", contact.DateOfBirthday, ":", name.String()})),
				Name:        name.String(),
				Description: description.String(),
				Date:        date.String(),
				ContactId:   contact.IdNumber,
			}
			fmt.Println(holiday)
			holidays = append(holidays, holiday)

			return true // keep iterating
		})

		if len(holidays) > 0 {
			r, err := dao.SaveHoliday(holidays[0])
			if !r {
				fmt.Println(err.Error())
			}
		}

		return holidays
	}

	return model.Holidays{}
}
