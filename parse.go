package dtf

import (
	"errors"
	"regexp"
	"time"
)

var regexZ *regexp.Regexp = regexp.MustCompile("Z$")

// Parse generate time.Time from W3C-DTF string
func Parse(timeStr string) (time.Time, error) {
	switch true {
	case IsYear(timeStr):
		return ParseYear(timeStr)
	case IsYearAndMonth(timeStr):
		return ParseYearAndMonth(timeStr)
	case IsCompleteDate(timeStr):
		return ParseCompleteDate(timeStr)
	case IsCompleteDateWithMinutes(timeStr):
		return ParseCompleteDateWithMinutes(timeStr)
	case IsCompleteDateWithSeconds(timeStr):
		return ParseCompleteDateWithSeconds(timeStr)
	case IsCompleteDateWithFractionOfSecond(timeStr):
		return ParseCompleteDateWithFractionOfSecond(timeStr)
	default:
		var parsedTime time.Time
		return parsedTime, errors.New("provided string is not W3C-DTF format")
	}
}

// ParseYear generate time.Time from 'YYYY' string
func ParseYear(timeStr string) (time.Time, error) {
	return time.Parse("2006", timeStr)
}

func ParseYearAndMonth(timeStr string) (time.Time, error) {
	return time.Parse("2006-01", timeStr)
}

func ParseCompleteDate(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02", timeStr)
}

func ParseCompleteDateWithMinutes(timeStr string) (time.Time, error) {
	if IsUTC(timeStr) {
		timeStr = regexZ.ReplaceAllString(timeStr, "+00:00")
	}
	return time.Parse("2006-01-02T15:04-07:00", timeStr)
}

func ParseCompleteDateWithSeconds(timeStr string) (time.Time, error) {
	if IsUTC(timeStr) {
		timeStr = regexZ.ReplaceAllString(timeStr, "+00:00")
	}
	return time.Parse("2006-01-02T15:04:05-07:00", timeStr)
}

func ParseCompleteDateWithFractionOfSecond(timeStr string) (time.Time, error) {
	if IsUTC(timeStr) {
		timeStr = regexZ.ReplaceAllString(timeStr, "+00:00")
	}
	return time.Parse("2006-01-02T15:04:05-07:00", timeStr)
}
