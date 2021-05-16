package said

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ademarj/said-go/db/model"
	"github.com/ademarj/said-go/helper"
	"github.com/agrison/go-commons-lang/stringUtils"
)

const (
	COUNTER_INIT   = 1
	NUMBER_ID_SIZE = 13
	layoutISO      = "2006-01-02"
)

func Year(numberId string) int {
	return helper.IntValue(stringUtils.Left(Birthday(numberId), 4))
}

func Month(numberId string) int {
	return helper.IntValue(MM(numberId))
}

func Day(numberId string) int {
	return helper.IntValue(DD(numberId))
}

func Gender(numberId string) string {
	if G(numberId) > 4 {
		return "M"
	}
	return "F"
}

func SouthAfricanCitizen(numberId string) bool {
	return C(numberId) == 0
}

func YY(numberId string) string {
	return stringUtils.Left(YYMMDD(numberId), 2)
}

func MM(numberId string) string {
	return stringUtils.Left(stringUtils.Right(YYMMDD(numberId), 4), 2)
}

func DD(numberId string) string {
	return stringUtils.Right(YYMMDD(numberId), 2)
}

func YYMMDD(numberId string) string {
	return stringUtils.Left(numberId, 6)
}

func Birthday(numberId string) string {
	return helper.GetFullYearFrom(YY(numberId), MM(numberId), DD(numberId))
}

func G(numberId string) int {
	code, _ := strconv.Atoi(stringUtils.Left(stringUtils.Right(numberId, 7), 1))
	return code
}

func C(numberId string) int {
	return helper.IntValue(stringUtils.Left(stringUtils.Right(numberId, 3), 1))
}

func A(numberId string) int {
	return helper.IntValue(stringUtils.Left(stringUtils.Right(numberId, 2), 1))
}

func CreateContactFrom(numberId string) (model.Contact, bool) {
	if !stringUtils.IsNumeric(numberId) {
		return model.Contact{}, false
	}

	if len(numberId) != NUMBER_ID_SIZE {
		return model.Contact{}, false
	}

	_, err := time.Parse(layoutISO, Birthday(numberId))
	if err != nil {
		return model.Contact{}, false
	}

	if C(numberId) > 1 {
		return model.Contact{}, false
	}

	sum := 0
	reversed := stringUtils.Reverse(numberId)

	for index := 0; index < NUMBER_ID_SIZE; index++ {
		if index%2 == 0 {
			sum += helper.IntValue(string(reversed[index]))
		} else {
			result := helper.IntValue(string(reversed[index])) * 2
			if result > A(numberId) {
				resultString := fmt.Sprint(result)
				sumCharByChar := 0
				for i := 0; i < len(resultString); i++ {
					sumCharByChar += helper.IntValue(string(resultString[i]))
				}
				sum += sumCharByChar
			} else {
				sum += result
			}
		}
	}

	if sum%10 != 0 {
		return model.Contact{}, false
	}

	contact := model.Contact{
		IdNumber:        numberId,
		Gender:          Gender(numberId),
		SaCitizen:       SouthAfricanCitizen(numberId),
		DateOfBirthday:  Birthday(numberId),
		Counter:         COUNTER_INIT,
		DayOfBirthday:   Day(numberId),
		MonthOfBirthday: Month(numberId),
	}

	return contact, true
}
