package said

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ademarj/said-go/helper"
	"github.com/agrison/go-commons-lang/stringUtils"
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
	caz := stringUtils.Right(numberId, 3)
	return helper.IntValue(stringUtils.Left(caz, 1)) == 0
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
	currentTime := time.Now()
	currentYY := stringUtils.Right(currentTime.Format("01-02-2006"), 2)
	yy := YY(numberId)

	if helper.IntValue(yy) >= helper.IntValue(currentYY) {
		return fmt.Sprintf("19%s-%s-%s", yy, MM(numberId), DD(numberId))
	}
	return fmt.Sprintf("20%s-%s-%s", currentYY, MM(numberId), DD(numberId))
}

func G(numberId string) int {
	gssscaz := stringUtils.Right(numberId, 7)
	genderCode := stringUtils.Left(gssscaz, 1)
	code, _ := strconv.Atoi(genderCode)
	return code
}
