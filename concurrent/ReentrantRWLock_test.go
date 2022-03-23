package concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestReentrantRWLock_RLock(t *testing.T) {
	lock := NewReentrantRWLock()
	lock.Lock()
	defer lock.Unlock()

	go func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("get lock")
	}()
	time.Sleep(time.Second)
	lock.Lock()
	defer lock.Unlock()
	lock.RLock()
	defer lock.RUnlock()
	lock.RLock()
	defer lock.RUnlock()
	fmt.Println("release lock")
}
