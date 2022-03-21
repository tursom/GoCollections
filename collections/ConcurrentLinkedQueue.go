package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"sync/atomic"
	"unsafe"
)

type ConcurrentLinkedQueue[T lang.Object] struct {
	head *concurrentLinkedQueueNode[T]
}

func (c ConcurrentLinkedQueue[T]) String() string {
	return String[T](c)
}

type concurrentLinkedQueueNode[T any] struct {
	value T
	prev  *concurrentLinkedQueueNode[T]
	next  *concurrentLinkedQueueNode[T]
}

type concurrentLinkedQueueIterator[T any] struct {
	head *concurrentLinkedQueueNode[T]
	node *concurrentLinkedQueueNode[T]
}

func NewConcurrentLinkedQueue[T lang.Object]() *ConcurrentLinkedQueue[T] {
	head := &concurrentLinkedQueueNode[T]{}
	head.prev = head
	head.next = head
	return &ConcurrentLinkedQueue[T]{head}
}

func (c ConcurrentLinkedQueue[T]) Iterator() Iterator[T] {
	return c.MutableIterator()
}

func (c *ConcurrentLinkedQueue[T]) Push(element T) exceptions.Exception {
	newNode := &concurrentLinkedQueueNode[T]{element, c.head.prev, c.head}
	p := (*unsafe.Pointer)(unsafe.Pointer(&c.head.prev))
	for !atomic.CompareAndSwapPointer(p, unsafe.Pointer(&*newNode.prev), unsafe.Pointer(newNode)) {
		newNode.prev = c.head.prev
	}
	atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&newNode.prev.next)),
		unsafe.Pointer(&*c.head),
		unsafe.Pointer(newNode),
	)
	return nil
}

func (c *ConcurrentLinkedQueue[T]) Offer() (T, exceptions.Exception) {
	next := c.head.next
	if next == c.head {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	}

	p := (*unsafe.Pointer)(unsafe.Pointer(&next.next.prev))

	if !next.removeNode(p) {
		next = c.head.next
		p = (*unsafe.Pointer)(unsafe.Pointer(&next.prev))
		if next == nil {
			return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
		}
	}

	return next.value, nil
}

func (node *concurrentLinkedQueueNode[T]) removeNode(p *unsafe.Pointer) bool {
	if p == nil {
		p = (*unsafe.Pointer)(unsafe.Pointer(&node.next.prev))
	}
	if !atomic.CompareAndSwapPointer(p, unsafe.Pointer(node), unsafe.Pointer(&*node.prev)) {
		return false
	}
	atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&node.prev.next)),
		unsafe.Pointer(node),
		unsafe.Pointer(&*node.next),
	)
	return true
}

func (c *ConcurrentLinkedQueue[T]) MutableIterator() MutableIterator[T] {
	return &concurrentLinkedQueueIterator[T]{c.head, c.head}
}

func (c *concurrentLinkedQueueIterator[T]) HasNext() bool {
	return c.node.next != c.head
}

func (c *concurrentLinkedQueueIterator[T]) Next() (T, exceptions.Exception) {
	c.node = c.node.next
	if c.node == c.head {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	}
	return c.node.value, nil
}

func (c *concurrentLinkedQueueIterator[T]) Remove() exceptions.Exception {
	if c.node == c.head {
		return exceptions.NewIndexOutOfBound("", nil)
	}
	c.node.removeNode(nil)
	c.node = c.node.prev
	return nil
}

func (c *ConcurrentLinkedQueue[T]) Size() int {
	size, err := Size[T](c)
	exceptions.Print(err)
	return size
}

func (c *ConcurrentLinkedQueue[T]) IsEmpty() bool {
	return c.head.next == c.head
}

func (c *ConcurrentLinkedQueue[T]) Contains(element T) bool {
	return Contains[T](c, element)
}

func (c *ConcurrentLinkedQueue[T]) ContainsAll(collection Collection[T]) bool {
	return ContainsAll[T](c, collection)
}

func (c *ConcurrentLinkedQueue[T]) Add(element T) bool {
	exception := c.Push(element)
	exceptions.Print(exception)
	return exception == nil
}

func (c *ConcurrentLinkedQueue[T]) Remove(element T) exceptions.Exception {
	return Remove[T](c, element)
}

func (c *ConcurrentLinkedQueue[T]) AddAll(collection Collection[T]) bool {
	return AddAll[T](c, collection)
}

func (c *ConcurrentLinkedQueue[T]) RemoveAll(collection Collection[T]) bool {
	return RemoveAll[T](c, collection)
}

func (c *ConcurrentLinkedQueue[T]) RetainAll(collection Collection[T]) bool {
	return RetainAll[T](c, collection)
}

func (c *ConcurrentLinkedQueue[T]) Clear() {
	head := &concurrentLinkedQueueNode[T]{}
	head.prev = head
	head.next = head
	c.head = head
}
