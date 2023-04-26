package time

import (
	"time"
)

const (
	Layout      = time.Layout
	ANSIC       = time.ANSIC
	UnixDate    = time.UnixDate
	RubyDate    = time.RubyDate
	RFC822      = time.RFC822
	RFC822Z     = time.RFC822Z
	RFC850      = time.RFC850
	RFC1123     = time.RFC1123
	RFC1123Z    = time.RFC1123Z
	RFC3339     = time.RFC3339
	RFC3339Nano = time.RFC3339Nano
	Kitchen     = time.Kitchen
	// Handy time stamps.
	Stamp      = time.Stamp
	StampMilli = time.StampMilli
	StampMicro = time.StampMicro
	StampNano  = time.StampNano
	DateTime   = time.DateTime
	DateOnly   = time.DateOnly
	TimeOnly   = time.TimeOnly
)

type ParseError = time.ParseError

func Parse(layout, value string) (Time, error) {
	return time.Parse(layout, value)
}

func ParseInLocation(layout, value string, loc *Location) (Time, error) {
	return time.ParseInLocation(layout, value, loc)
}

func ParseDuration(s string) (Duration, error) {
	return time.ParseDuration(s)
}
