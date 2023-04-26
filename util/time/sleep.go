package time

import "time"

func Sleep(d Duration) {
	time.Sleep(d)
}

type Timer = time.Timer

func NewTimer(d Duration) *Timer {
	return time.NewTimer(d)
}

func After(d Duration) <-chan Time {
	return time.After(d)
}

func AfterFunc(d Duration, f func()) *Timer {
	return time.AfterFunc(d, f)
}
