package concurrent

import (
	"sync"

	"github.com/petermattis/goid"
)

func GetGoroutineID() int64 {
	return goid.Get()
}

func WaitCond(cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()
	cond.Wait()
}
