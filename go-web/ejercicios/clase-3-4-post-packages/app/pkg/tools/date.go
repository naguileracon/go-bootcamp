package tools

import (
	"fmt"
	"regexp"
	"strconv"
)

type DateError struct {
	Date string
	Msg  string
}

func (e *DateError) Error() string {
	return fmt.Sprintf("%s: %s", e.Date, e.Msg)
}

// CheckDate Check if a date follows DD/MM/YYYY format and day, month, year contain valid values
func CheckDate(date string) (err error) {
	dateRegex := regexp.MustCompile(`^(\d{2})/(\d{2})/(\d{4})$`)

	matches := dateRegex.FindStringSubmatch(date)
	if len(matches) != 4 {
		err = &DateError{
			Date: date,
			Msg:  "date format does not follow DD/MM/YYYY format",
		}
		return
	}

	day, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	year, _ := strconv.Atoi(matches[3])

	if day < 1 || day > 31 || month < 1 || month > 12 || year < 1000 {
		err = &DateError{
			Date: date,
			Msg:  "date does not contain valid values",
		}
		return
	}

	return
}
