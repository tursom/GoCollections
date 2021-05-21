package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"sync/atomic"
	"unsafe"
)

type ConcurrentLinkedQueue struct {
	head *concurrentLinkedQueueNode
	tail *concurrentLinkedQueueNode
}

func (c ConcurrentLinkedQueue) String() string {
	return String(c)
}

type concurrentLinkedQueueNode struct {
	value interface{}
	next  *concurrentLinkedQueueNode
}

type concurrentLinkedQueueIterator struct {
	prev *concurrentLinkedQueueNode
	head *concurrentLinkedQueueNode
}

func NewConcurrentLinkedQueue() *ConcurrentLinkedQueue {
	head := &concurrentLinkedQueueNode{}
	return &ConcurrentLinkedQueue{head, head}
}

func (c ConcurrentLinkedQueue) Iterator() Iterator {
	return c.MutableIterator()
}

func (c *ConcurrentLinkedQueue) Push(element interface{}) exceptions.Exception {
	newNode := &concurrentLinkedQueueNode{element, nil}
	p := (*unsafe.Pointer)(unsafe.Pointer(&c.tail.next))
	np := unsafe.Pointer(newNode)
	for !atomic.CompareAndSwapPointer(p, nil, np) {
		p = (*unsafe.Pointer)(unsafe.Pointer(&c.tail.next))
	}
	c.tail = newNode
	return nil
}

func (c *ConcurrentLinkedQueue) Offer() (interface{}, exceptions.Exception) {
	next := c.head.next
	if next == nil {
		return nil, exceptions.NewIndexOutOfBound("", true)
	}

	p := (*unsafe.Pointer)(unsafe.Pointer(&c.head.next))

	if !atomic.CompareAndSwapPointer(p, unsafe.Pointer(next), unsafe.Pointer(next.next)) {
		next = c.head.next
		if next == nil {
			return nil, exceptions.NewIndexOutOfBound("", true)
		}
	}

	return next.value, nil
}

func (c *ConcurrentLinkedQueue) MutableIterator() MutableIterator {
	return &concurrentLinkedQueueIterator{nil, c.head}
}

func (c *concurrentLinkedQueueIterator) HasNext() bool {
	return c.head.next != nil
}

func (c *concurrentLinkedQueueIterator) Next() (interface{}, exceptions.Exception) {
	c.prev = c.head
	c.head = c.head.next
	if c.head == nil {
		return nil, exceptions.NewIndexOutOfBound("", true)
	}
	return c.head.value, nil
}

func (c *concurrentLinkedQueueIterator) Remove() exceptions.Exception {
	next := c.head.next
	if next == nil {
		return exceptions.NewIndexOutOfBound("", true)
	}

	p := (*unsafe.Pointer)(unsafe.Pointer(&c.prev.next))
	atomic.CompareAndSwapPointer(p, unsafe.Pointer(c.head), unsafe.Pointer(next))
	return nil
}
