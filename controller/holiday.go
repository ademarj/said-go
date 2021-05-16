package controller

import (
	"fmt"

	"github.com/ademarj/said-go/db/dao"
	"github.com/ademarj/said-go/db/model"
	"github.com/ademarj/said-go/helper"
	"github.com/ademarj/said-go/http/api"
	"github.com/ademarj/said-go/said"
	"github.com/ademarj/said-go/util/security"
	"github.com/ademarj/said-go/view"
	"github.com/tidwall/gjson"
)

func BuildViewHoliday(saNumberId string) (view.Holidays, bool) {
	contact, success := said.CreateContactFrom(saNumberId)

	if success {
		contactFromDB, _ := dao.FindBy(contact.IdNumber)
		if contact.IdNumber == contactFromDB.IdNumber {
			contact.Counter = contact.Counter + contactFromDB.Counter
			dao.UpdateContact(contact)
		} else {
			dao.SaveNewContact(contact)
		}
		return searchHolidays(contact), true
	}

	return view.Holidays{}, false
}

func searchHolidays(contact model.Contact) view.Holidays {
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

		return prepareView(append(filtered, holidaysFromContact...))
	}

	holidaysFromContact, _ := dao.GetHolidaysFrom(contact.IdNumber)
	return prepareView(holidaysFromContact)
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

func prepareView(holidaysMerged []model.Holiday) view.Holidays {
	if len(holidaysMerged) <= 0 {
		return view.Holidays{}
	}
	var holidays []view.Holiday
	columnGrid := 0
	rowGrid := 1
	for _, h := range holidaysMerged {
		columnGrid += 1

		holidays = append(holidays, view.Holiday{
			Id:          h.Id,
			Name:        h.Name,
			Description: h.Description,
			Date:        h.Date,
			ContactId:   h.ContactId,
			GridColumn:  columnGrid,
			GridRow:     rowGrid,
			Calendar: view.Calendar{
				Year:           helper.Year(h.Date),
				Month:          helper.Month(h.Date),
				Day:            helper.Day(h.Date),
				LastDayOfMonth: helper.LastDayOfMonth(h.Date),
				Days:           helper.DaysOfMonth(helper.LastDayOfMonth(h.Date)),
			},
		})
		if columnGrid == 4 {
			columnGrid = 0
			rowGrid += 1
		}
	}
	return view.Holidays{Holidays: holidays}
}
