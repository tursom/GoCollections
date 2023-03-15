package lock

import (
	"github.com/petermattis/goid"

	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang/atomic"
)

type (
	CLH struct {
		tail  atomic.Reference[clhNode]
		state atomic.Reference[clhNode]
	}

	clhNode struct {
		prev   *clhNode
		locked atomic.Bool
		gid    int64
	}
)

func (c *CLH) Lock() {
	node := &clhNode{
		prev:   c.tail.Load(),
		locked: 1, // true
		gid:    goid.Get(),
	}

	for !c.tail.CompareAndSwap(node.prev, node) {
		node.prev = c.tail.Load()
	}

	if node.prev == nil {
		return
	}

	for node.prev.locked.Load() {
	}

	c.state.Store(node)
}

func (c *CLH) Unlock() {
	node := c.state.Load()
	if node.gid != goid.Get() {
		panic(exceptions.NewIllegalAccessException("unlock with wrong goroutine", nil))
	}

	node.locked.Store(false)
}
