package concurrent

import (
	"sync"
	"time"
)

type (
	Cond interface {
		Signal()
		Broadcast()
		// Wait wait within unlocked environment
		Wait()
		// LockedWait wait within locked environment
		LockedWait()
		WaitOrTimeout(timeout time.Duration) bool
	}

	CondImpl sync.Cond
)

func NewCond(l sync.Locker) Cond {
	return (*CondImpl)(sync.NewCond(l))
}

func (c *CondImpl) Signal() {
	(*sync.Cond)(c).Signal()
}

func (c *CondImpl) Broadcast() {
	(*sync.Cond)(c).Broadcast()
}

func (c *CondImpl) Wait() {
	c.L.Lock()
	defer c.L.Unlock()
	(*sync.Cond)(c).Wait()
}

func (c *CondImpl) LockedWait() {
	(*sync.Cond)(c).Wait()
}

func (c *CondImpl) WaitOrTimeout(timeout time.Duration) bool {
	done := make(chan struct{})
	go func() {
		c.L.Lock()
		defer c.L.Unlock()

		c.Wait()
		close(done)
	}()

	select {
	case <-time.After(timeout):
		return false
	case <-done:
		return true
	}
}
