package time

import (
	"time"
)

type Location = time.Location

var UTC = time.UTC

var Local = time.Local

func FixedZone(name string, offset int) *Location {
	return time.FixedZone(name, offset)
}

func LoadLocation(name string) (*Location, error) {
	return time.LoadLocation(name)
}
