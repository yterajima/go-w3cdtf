package dtf

import (
	"errors"
	"strings"
	"time"
)

// Parse generate time.Time from W3C-DTF string
func Parse(timeStr string) (time.Time, error) {
	err := errors.New("provided string is not W3C-DTF format")

	if !strings.Contains(timeStr, "T") {
		switch true {
		case IsYear(timeStr):
			return ParseYear(timeStr)
		case IsYearAndMonth(timeStr):
			return ParseYearAndMonth(timeStr)
		case IsCompleteDate(timeStr):
			return ParseCompleteDate(timeStr)
		default:
			return time.Time{}, err
		}
	} else {
		switch true {
		case IsCompleteDateWithMinutes(timeStr):
			return ParseCompleteDateWithMinutes(timeStr)
		case IsCompleteDateWithSeconds(timeStr):
			return ParseCompleteDateWithSeconds(timeStr)
		case IsCompleteDateWithFractionOfSecond(timeStr):
			return ParseCompleteDateWithFractionOfSecond(timeStr)
		default:
			return time.Time{}, err
		}
	}
}

// ParseYear generate time.Time from 'YYYY' string
func ParseYear(timeStr string) (time.Time, error) {
	return time.Parse("2006", timeStr)
}

// ParseYearAndMonth generate time.Time from 'YYYY-MM' string
func ParseYearAndMonth(timeStr string) (time.Time, error) {
	return time.Parse("2006-01", timeStr)
}

// ParseCompleteDate generate time.Time from 'YYYY-MM-DD' string
func ParseCompleteDate(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02", timeStr)
}

// ParseCompleteDateWithMinutes generate time.Time from 'YYYY-MM-DDThh:ii+00:00'
func ParseCompleteDateWithMinutes(timeStr string) (time.Time, error) {
	if IsUTC(timeStr) {
		timeStr = rUTC.ReplaceAllString(timeStr, "+00:00")
	}
	return time.Parse("2006-01-02T15:04-07:00", timeStr)
}

// ParseCompleteDateWithSeconds generate time.Time from 'YYYY-MM-DDThh:ii:ss+00:00'
func ParseCompleteDateWithSeconds(timeStr string) (time.Time, error) {
	if IsUTC(timeStr) {
		timeStr = rUTC.ReplaceAllString(timeStr, "+00:00")
	}
	return time.Parse("2006-01-02T15:04:05-07:00", timeStr)
}

// ParseCompleteDateWithFractionOfSecond generate time.Time from 'YYYY-MM-DDThh:ii:ss.s+00:00'
func ParseCompleteDateWithFractionOfSecond(timeStr string) (time.Time, error) {
	return ParseCompleteDateWithSeconds(timeStr)
}
