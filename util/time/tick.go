package time

import (
	"time"
	unsafe2 "unsafe"

	"github.com/tursom/GoCollections/unsafe"
)

type Ticker struct {
	time.Ticker
}

func NewTicker(d Duration) *Ticker {
	return &Ticker{*time.NewTicker(d)}
}

func (t *Ticker) Stop() {
	t.Ticker.Stop()
	close(*unsafe.ForceCast[chan time.Time](unsafe2.Pointer(&t.C)))
}

func Tick(d Duration) <-chan Time {
	return time.Tick(d)
}
