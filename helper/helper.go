package helper

import (
	"strconv"
	"time"

	"github.com/agrison/go-commons-lang/stringUtils"
)

const (
	layOut = "2006-01-02"
)

func IntValue(source string) int {
	result, _ := strconv.Atoi(source)
	return result
}

func CurrentYear() int {
	return IntValue(stringUtils.Left(time.Now().Format(layOut), 4))
}

func Year(date string) string {
	return stringUtils.Left(date, 4)
}

func Day(date string) int {
	return IntValue(stringUtils.Right(date, 2))
}

func Month(date string) string {
	dateStamp, _ := time.Parse(layOut, date)
	return dateStamp.Month().String()
}

func LastDayOfMonth(date string) int {
	now, _ := time.Parse(layOut, date)
	year, month, _ := now.Date()
	return Day(time.Date(year, month+1, 0, 0, 0, 0, 0, now.Location()).Format(layOut))
}

func DaysOfMonth(count int) []int {
	days := []int{}
	for i := 0; i < count; i++ {
		days = append(days, i+1)
	}
	return days
}
