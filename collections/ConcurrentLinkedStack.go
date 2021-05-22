package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"sync/atomic"
	"unsafe"
)

type ConcurrentLinkedStack struct {
	head *concurrentLinkedStackNode
	p    *unsafe.Pointer
}

func (c ConcurrentLinkedStack) String() string {
	return String(c)
}

type concurrentLinkedStackNode struct {
	value interface{}
	next  *concurrentLinkedStackNode
}

type concurrentLinkedStackIterator struct {
	node *concurrentLinkedStackNode
	prev *concurrentLinkedStackNode
}

func NewConcurrentLinkedStack() *ConcurrentLinkedStack {
	head := &concurrentLinkedStackNode{}
	return &ConcurrentLinkedStack{head, (*unsafe.Pointer)(unsafe.Pointer(&head.next))}
}

func (c ConcurrentLinkedStack) Iterator() Iterator {
	return c.MutableIterator()
}

func (c *ConcurrentLinkedStack) Push(element interface{}) exceptions.Exception {
	newNode := &concurrentLinkedStackNode{element, c.head.next}
	np := unsafe.Pointer(newNode)
	for !atomic.CompareAndSwapPointer(c.p, unsafe.Pointer(newNode.next), np) {
	}
	return nil
}

func (c *ConcurrentLinkedStack) Pop() (interface{}, exceptions.Exception) {
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

func (c *ConcurrentLinkedStack) MutableIterator() MutableIterator {
	return &concurrentLinkedStackIterator{c.head, nil}
}

func (c *concurrentLinkedStackIterator) HasNext() bool {
	return c.node.next != nil
}

func (c *concurrentLinkedStackIterator) Next() (interface{}, exceptions.Exception) {
	c.prev = c.node
	c.node = c.node.next
	if c.node == nil {
		return nil, exceptions.NewIndexOutOfBound("", true)
	}
	return c.node.value, nil
}

func (c *concurrentLinkedStackIterator) Remove() exceptions.Exception {
	if c.node == nil {
		return exceptions.NewIndexOutOfBound("", true)
	}
	next := c.node.next

	p := (*unsafe.Pointer)(unsafe.Pointer(&c.prev.next))
	atomic.CompareAndSwapPointer(p, unsafe.Pointer(c.node), unsafe.Pointer(next))
	return nil
}
