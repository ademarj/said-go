package model

import (
	"fmt"
	"strconv"

	"github.com/agrison/go-commons-lang/stringUtils"
)

type Contact struct {
	IdNumber       string `db:"id_number"`
	DateOfBirthday string `db:"date_of_birth"`
	Gender         string `db:"gender"`
	SaCitizen      bool   `db:"sa_citizen"`
	Counter        string `db:"counter"`
}

type Contacts []Contact

func CreateContactFrom(idNumber string) (Contact, error) {

	fmt.Printf("ID Number %s\n", idNumber)

	yymmdd := stringUtils.Left(idNumber, 6)
	fmt.Printf("yymmdd %s\n", yymmdd)

	gssscaz := stringUtils.Right(idNumber, 7)
	fmt.Printf("gssscaz %s\n", gssscaz)

	citizenCode := stringUtils.Left(stringUtils.Right(gssscaz, 3), 1)

	genderCode := stringUtils.Left(gssscaz, 1)
	fmt.Printf("gender %s\n", genderFrom(genderCode))

	contact := Contact{
		IdNumber:  idNumber,
		Gender:    genderFrom(genderCode),
		SaCitizen: isSACitizen(citizenCode),
	}

	return contact, nil

}

func genderFrom(genderCode string) string {
	code, _ := strconv.Atoi(genderCode)
	if code > 4 {
		return "M"
	}

	return "F"
}

func isSACitizen(citizenCode string) bool {
	code, _ := strconv.Atoi(citizenCode)
	return code == 0
}
