package time

import (
	"time"
)

type Time = time.Time

// A Month specifies a month of the year (January = 1, ...).
type Month = time.Month

const (
	January   = time.January
	February  = time.February
	March     = time.March
	April     = time.April
	May       = time.May
	June      = time.June
	July      = time.July
	August    = time.August
	September = time.September
	October   = time.October
	November  = time.November
	December  = time.December
)

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday = time.Weekday

const (
	Sunday    = time.Sunday
	Monday    = time.Monday
	Tuesday   = time.Tuesday
	Wednesday = time.Wednesday
	Thursday  = time.Thursday
	Friday    = time.Friday
	Saturday  = time.Saturday
)

type Duration = time.Duration

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
)

// Now returns the current local time.
func Now() Time {
	return time.Now()
}

func Unix(sec int64, nsec int64) Time {
	return time.Unix(sec, nsec)
}

func UnixMilli(msec int64) Time {
	return time.UnixMilli(msec)
}

func UnixMicro(usec int64) Time {
	return time.UnixMicro(usec)
}

func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
