package helper

import (
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
