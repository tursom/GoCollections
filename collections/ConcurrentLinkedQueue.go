package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"sync/atomic"
	"unsafe"
)

type ConcurrentLinkedQueue struct {
	head *concurrentLinkedQueueNode
}

func (c ConcurrentLinkedQueue) String() string {
	return String(c)
}

type concurrentLinkedQueueNode struct {
	value interface{}
	prev  *concurrentLinkedQueueNode
	next  *concurrentLinkedQueueNode
}

type concurrentLinkedQueueIterator struct {
	head *concurrentLinkedQueueNode
	node *concurrentLinkedQueueNode
}

func NewConcurrentLinkedQueue() *ConcurrentLinkedQueue {
	head := &concurrentLinkedQueueNode{}
	head.prev = head
	head.next = head
	return &ConcurrentLinkedQueue{head}
}

func (c ConcurrentLinkedQueue) Iterator() Iterator {
	return c.MutableIterator()
}

func (c *ConcurrentLinkedQueue) Push(element interface{}) exceptions.Exception {
	newNode := &concurrentLinkedQueueNode{element, c.head.prev, c.head}
	p := (*unsafe.Pointer)(unsafe.Pointer(&c.head.prev))
	for !atomic.CompareAndSwapPointer(p, unsafe.Pointer(newNode.prev), unsafe.Pointer(newNode)) {
		newNode.prev = c.head.prev
	}
	atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&newNode.prev.next)),
		unsafe.Pointer(c.head),
		unsafe.Pointer(newNode),
	)
	return nil
}

func (c *ConcurrentLinkedQueue) Offer() (interface{}, exceptions.Exception) {
	next := c.head.next
	if next == c.head {
		return nil, exceptions.NewIndexOutOfBound("", true)
	}

	p := (*unsafe.Pointer)(unsafe.Pointer(&next.next.prev))

	if !next.removeNode(p) {
		next = c.head.next
		p = (*unsafe.Pointer)(unsafe.Pointer(&next.prev))
		if next == nil {
			return nil, exceptions.NewIndexOutOfBound("", true)
		}
	}

	return next.value, nil
}

func (node *concurrentLinkedQueueNode) removeNode(p *unsafe.Pointer) bool {
	if p == nil {
		p = (*unsafe.Pointer)(unsafe.Pointer(&node.next.prev))
	}
	if !atomic.CompareAndSwapPointer(p, unsafe.Pointer(node), unsafe.Pointer(node.prev)) {
		return false
	}
	atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&node.prev.next)),
		unsafe.Pointer(node),
		unsafe.Pointer(node.next),
	)
	return true
}

func (c *ConcurrentLinkedQueue) MutableIterator() MutableIterator {
	return &concurrentLinkedQueueIterator{c.head, c.head}
}

func (c *concurrentLinkedQueueIterator) HasNext() bool {
	return c.node.next != c.head
}

func (c *concurrentLinkedQueueIterator) Next() (interface{}, exceptions.Exception) {
	c.node = c.node.next
	if c.node == c.head {
		return nil, exceptions.NewIndexOutOfBound("", true)
	}
	return c.node.value, nil
}

func (c *concurrentLinkedQueueIterator) Remove() exceptions.Exception {
	if c.node == c.head {
		return exceptions.NewIndexOutOfBound("", true)
	}
	c.node.removeNode(nil)
	c.node = c.node.prev
	return nil
}

func (c *ConcurrentLinkedQueue) Size() uint32 {
	size, err := Size(c)
	exceptions.Print(err)
	return size
}

func (c *ConcurrentLinkedQueue) IsEmpty() bool {
	return c.head.next == c.head
}

func (c *ConcurrentLinkedQueue) Contains(element interface{}) bool {
	return Contains(c, element)
}

func (c *ConcurrentLinkedQueue) ContainsAll(collection Collection) bool {
	return ContainsAll(c, collection)
}

func (c *ConcurrentLinkedQueue) Add(element interface{}) bool {
	exception := c.Push(element)
	exceptions.Print(exception)
	return exception == nil
}

func (c *ConcurrentLinkedQueue) Remove(element interface{}) exceptions.Exception {
	return Remove(c, element)
}

func (c *ConcurrentLinkedQueue) AddAll(collection Collection) bool {
	return AddAll(c, collection)
}

func (c *ConcurrentLinkedQueue) RemoveAll(collection Collection) bool {
	return RemoveAll(c, collection)
}

func (c *ConcurrentLinkedQueue) RetainAll(collection Collection) bool {
	return RetainAll(c, collection)
}

func (c *ConcurrentLinkedQueue) Clear() {
	head := &concurrentLinkedQueueNode{}
	head.prev = head
	head.next = head
	c.head = head
}
