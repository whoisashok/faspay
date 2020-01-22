package currentdate

import (
	"time"

	helperPayment "payment_http/payment_services/helper"
)

func CurrentUTCDateTime() string {
	currentUTC := time.Now().UTC()
	result := currentUTC.Format(helperPayment.DateTimeFormat)
	return result
}

func CurrentUTCDateTimeMilSec() string {
	currentUTC := time.Now().UTC()
	result := currentUTC.Format(helperPayment.DateTimeMilSecFormat)
	return result
}

func CurrentUTCDateTimeMacSec() string {
	currentUTC := time.Now().UTC()
	result := currentUTC.Format(helperPayment.DateTimeMacSecFormat)
	return result
}

func CurrentUTCDateTimeNanSec() string {
	currentUTC := time.Now().UTC()
	result := currentUTC.Format(helperPayment.DateTimeNanSecFormat)
	return result
}

func ParseDateFromString(timeRaw string) time.Time {
	const timeLayout = "2006-01-02"
	t, _ := time.Parse(timeLayout, timeRaw)
	return t
}

func ParseDate(date time.Time) time.Time {
	const timeLayout = "2006-01-02"
	timeRaw := date.Format(timeLayout)
	t, _ := time.Parse(timeLayout, timeRaw)
	return t
}

func ParseTimeFrom(from time.Time) time.Time {
	const timeLayout = "2006-01-02 00:00:00"
	timeRaw := from.Format(timeLayout)
	t, _ := time.Parse(timeLayout, timeRaw)
	return t
}

func ParseTimeTo(to time.Time) time.Time {
	const timeLayout = "2006-01-02 15:04:05"
	timeRaw := to.Format(timeLayout)
	t, _ := time.Parse(timeLayout, timeRaw)
	return t
}
