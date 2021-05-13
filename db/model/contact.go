package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/agrison/go-commons-lang/stringUtils"
)

const (
	ONE = 1
)

type Contact struct {
	IdNumber       string `db:"id_number"`
	DateOfBirthday string `db:"date_of_birth"`
	Gender         string `db:"gender"`
	SaCitizen      bool   `db:"sa_citizen"`
	Counter        int32  `db:"counter"`
}

type Contacts []Contact

func CreateContactFrom(idNumber string) (Contact, error) {
	fmt.Printf("ID Number %s\n", idNumber)
	yymmdd := stringUtils.Left(idNumber, 6)
	gssscaz := stringUtils.Right(idNumber, 7)
	citizenCode := stringUtils.Left(stringUtils.Right(gssscaz, 3), 1)
	genderCode := stringUtils.Left(gssscaz, 1)

	contact := Contact{
		IdNumber:       idNumber,
		Gender:         genderFrom(genderCode),
		SaCitizen:      isSACitizen(citizenCode),
		DateOfBirthday: formatDate(yymmdd),
		Counter:        ONE,
	}

	return contact, nil
}

func formatDate(yymmdd string) string {
	yyS := stringUtils.Left(yymmdd, 2)
	yy, _ := strconv.Atoi(yyS)
	dd := stringUtils.Right(yymmdd, 2)
	mm := stringUtils.Left(stringUtils.Right(yymmdd, 4), 2)
	currentTime := time.Now()
	currentYearS := stringUtils.Right(currentTime.Format("01-02-2006"), 2)
	currentYear, _ := strconv.Atoi(currentYearS)
	if yy >= currentYear {
		return "19" + yyS + "-" + mm + "-" + dd
	}

	return "20" + currentYearS + "-" + mm + "-" + dd
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
