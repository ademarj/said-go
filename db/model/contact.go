package model

import (
	"github.com/ademarj/said-go/util/said"
)

const (
	ONE = 1
)

type Contact struct {
	IdNumber        string `db:"id_number"`
	DateOfBirthday  string `db:"date_of_birth"`
	Gender          string `db:"gender"`
	SaCitizen       bool   `db:"sa_citizen"`
	Counter         int32  `db:"counter"`
	DayOfBirthday   int
	MonthOfBirthday int
}

type Contacts []Contact

func CreateContactFrom(numberId string) (Contact, error) {

	contact := Contact{
		IdNumber:        numberId,
		Gender:          said.Gender(numberId),
		SaCitizen:       said.SouthAfricanCitizen(numberId),
		DateOfBirthday:  said.Birthday(numberId),
		Counter:         ONE,
		DayOfBirthday:   said.Day(numberId),
		MonthOfBirthday: said.Month(numberId),
	}

	return contact, nil
}
