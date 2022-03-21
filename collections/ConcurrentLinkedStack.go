package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"sync/atomic"
	"unsafe"
)

type ConcurrentLinkedStack[T any] struct {
	head *concurrentLinkedStackNode[T]
	p    *unsafe.Pointer
}

func (c ConcurrentLinkedStack[T]) String() string {
	return String[T](c)
}

type concurrentLinkedStackNode[T any] struct {
	value T
	next  *concurrentLinkedStackNode[T]
}

type concurrentLinkedStackIterator[T any] struct {
	node *concurrentLinkedStackNode[T]
	prev *concurrentLinkedStackNode[T]
}

func NewConcurrentLinkedStack[T any]() *ConcurrentLinkedStack[T] {
	head := &concurrentLinkedStackNode[T]{}
	return &ConcurrentLinkedStack[T]{head, (*unsafe.Pointer)(unsafe.Pointer(&head.next))}
}

func (c ConcurrentLinkedStack[T]) Iterator() Iterator[T] {
	return c.MutableIterator()
}

func (c *ConcurrentLinkedStack[T]) Push(element T) exceptions.Exception {
	newNode := &concurrentLinkedStackNode[T]{element, c.head.next}
	np := unsafe.Pointer(newNode)
	for !atomic.CompareAndSwapPointer(c.p, unsafe.Pointer(&*newNode.next), np) {
	}
	return nil
}

func (c *ConcurrentLinkedStack[T]) Pop() (T, exceptions.Exception) {
	next := c.head.next
	if next == nil {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	}

	p := (*unsafe.Pointer)(unsafe.Pointer(&c.head.next))

	if !atomic.CompareAndSwapPointer(p, unsafe.Pointer(&*next), unsafe.Pointer(&*next.next)) {
		next = c.head.next
		if next == nil {
			return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
		}
	}

	return next.value, nil
}

func (c *ConcurrentLinkedStack[T]) MutableIterator() MutableIterator[T] {
	return &concurrentLinkedStackIterator[T]{c.head, nil}
}

func (c *concurrentLinkedStackIterator[T]) HasNext() bool {
	return c.node.next != nil
}

func (c *concurrentLinkedStackIterator[T]) Next() (T, exceptions.Exception) {
	c.prev = c.node
	c.node = c.node.next
	if c.node == nil {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	}
	return c.node.value, nil
}

func (c *concurrentLinkedStackIterator[T]) Remove() exceptions.Exception {
	if c.node == nil {
		return exceptions.NewIndexOutOfBound("", nil)
	}
	next := c.node.next

	p := (*unsafe.Pointer)(unsafe.Pointer(&c.prev.next))
	atomic.CompareAndSwapPointer(p, unsafe.Pointer(&*c.node), unsafe.Pointer(&*next))
	return nil
}
