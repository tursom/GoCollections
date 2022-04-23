package collections

import (
	"sync"

	"github.com/tursom/GoCollections/concurrent"
	"github.com/tursom/GoCollections/lang/atomic"
)

type (
	MessageQueue[T any] struct {
		lock sync.Mutex
		cond *sync.Cond
		end  *messageQueueNode[T]
	}

	messageQueueNode[T any] struct {
		value T
		next  *messageQueueNode[T]
	}
)

func (q *MessageQueue[T]) getEnd() *messageQueueNode[T] {
	if q.end == nil {
		q.lock.Lock()
		defer q.lock.Unlock()
		if q.end == nil {
			q.end = &messageQueueNode[T]{}
		}
	}
	return q.end
}

func (q *MessageQueue[T]) getCond() *sync.Cond {
	if q.cond == nil {
		q.lock.Lock()
		defer q.lock.Unlock()
		q.cond = sync.NewCond(&q.lock)
	}
	return q.cond
}

func (q *MessageQueue[T]) Subscribe() (in <-chan T, canceler func()) {
	end := q.getEnd()
	ch := make(chan T)
	canceled := false
	go func() {
		cond := q.getCond()
		node := end
		for !canceled {
			for node.next != nil {
				if canceled {
					return
				}
				node = node.next
				ch <- node.value
			}
			concurrent.WaitCond(cond)
		}
	}()
	return ch, func() {
		canceled = true
	}
}

func (q *MessageQueue[T]) Send(msg T) {
	node := &messageQueueNode[T]{
		value: msg,
	}
	p := &q.getEnd().next
	for !atomic.CompareAndSwapPointer(p, nil, node) {
		for *p != nil {
			p = &q.end.next
		}
	}
	q.end = node
	q.getCond().Broadcast()
}
