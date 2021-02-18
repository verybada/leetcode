package reformatDate

import (
	"fmt"
	"strconv"
	"strings"
)

func reformatDate(date string) string {
	parts := strings.Split(date, " ")
	day := parts[0]
	month := parts[1]
	year := parts[2]
	return fmt.Sprintf("%s-%02d-%02d", year, monthToInt(month), dayToInt(day))
}

func dayToInt(day string) int {
	s := ""
	if len(day) == 3 {
		s = string(day[0])
	} else {
		s = string(day[0:2])
	}

	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return value
}

func monthToInt(month string) int {
	switch month {
	case "Jan":
		return 1
	case "Feb":
		return 2
	case "Mar":
		return 3
	case "Apr":
		return 4
	case "May":
		return 5
	case "Jun":
		return 6
	case "Jul":
		return 7
	case "Aug":
		return 8
	case "Sep":
		return 9
	case "Oct":
		return 10
	case "Nov":
		return 11
	case "Dec":
		return 12
	}
	return 0
}
