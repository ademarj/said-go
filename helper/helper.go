package helper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func IntValue(source string) int {
	result, _ := strconv.Atoi(source)
	return result
}

func CurrentYear() int {
	currentTime := time.Now()
	currentYear := stringUtils.Right(currentTime.Format("01-02-2006"), 4)
	return IntValue(currentYear)
}

func Year(date string) string {
	return stringUtils.Left(date, 4)
}

func Day(date string) int {
	return IntValue(stringUtils.Left(date, 2))
}

func Month(date string) string {
	layOut := "2006-01-02" // yyyy-dd-MM
	dateStamp, _ := time.Parse(layOut, date)
	return dateStamp.Month().String()
}

func LastDayOfMonth(date string) int {
	layOut := "2006-01-02" // yyyy-dd-MM
	now, _ := time.Parse(layOut, date)
	year, month, _ := now.Date()
	endOfThisMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, now.Location())
	convertedDateString := endOfThisMonth.Format("2-Jan-2006")
	lastDay := IntValue(stringUtils.Left(convertedDateString, 2))
	fmt.Printf("Last Day : %d\n", lastDay)
	return lastDay
}
