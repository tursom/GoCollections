package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"github.com/tursom/GoCollections/lang/atomic"
)

type (
	ConcurrentLinkedStack[T lang.Object] struct {
		lang.BaseObject
		size    atomic.Int32
		deleted atomic.Int32
		head    *concurrentLinkedStackNode[T]
	}

	concurrentLinkedStackNode[T any] struct {
		deleted bool
		next    *concurrentLinkedStackNode[T]
		value   T
	}

	concurrentLinkedStackIterator[T lang.Object] struct {
		node  *concurrentLinkedStackNode[T]
		stack *ConcurrentLinkedStack[T]
	}
)

func (s *ConcurrentLinkedStack[T]) String() string {
	return String[T](s)
}

func NewConcurrentLinkedStack[T lang.Object]() *ConcurrentLinkedStack[T] {
	return &ConcurrentLinkedStack[T]{}
}

func (s *ConcurrentLinkedStack[T]) Iterator() Iterator[T] {
	return s.MutableIterator()
}

func (s *ConcurrentLinkedStack[T]) Push(element T) exceptions.Exception {
	newNode := &concurrentLinkedStackNode[T]{value: element, next: s.head}
	for !atomic.CompareAndSwapPointer(&s.head, newNode.next, newNode) {
		newNode.next = s.head
	}
	s.size.Add(1)
	return nil
}

func (s *ConcurrentLinkedStack[T]) Pop() (T, exceptions.Exception) {
	node := s.head
	for node != nil && (!atomic.CompareAndSwapPointer(&s.head, node, node.next) || node.deleted) {
		node = s.head
	}
	if node == nil {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	}

	s.size.Add(-1)
	return node.value, nil
}

func (s *ConcurrentLinkedStack[T]) CleanDeleted() {
	if s.deleted.Load() <= s.size.Load() {
		return
	}
	s.deleted.Store(0)
	node := &s.head
	for node != nil && *node != nil {
		if (*node).deleted {
			atomic.CompareAndSwapPointer(node, *node, (*node).next)
		} else {
			node = &(*node).next
		}
	}
}

func (s *ConcurrentLinkedStack[T]) Size() int {
	return int(s.size.Load())
}

func (s *ConcurrentLinkedStack[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *ConcurrentLinkedStack[T]) Contains(element T) bool {
	return Contains[T](s, element)
}

func (s *ConcurrentLinkedStack[T]) ContainsAll(c Collection[T]) bool {
	return ContainsAll[T](s, c)
}

func (s *ConcurrentLinkedStack[T]) Add(element T) bool {
	exception := s.Push(element)
	exceptions.Print(exception)
	return exception == nil
}

func (s *ConcurrentLinkedStack[T]) Remove(element T) exceptions.Exception {
	return Remove[T](s, element)
}

func (s *ConcurrentLinkedStack[T]) AddAll(c Collection[T]) bool {
	return AddAll[T](s, c)
}

func (s *ConcurrentLinkedStack[T]) RemoveAll(c Collection[T]) bool {
	return RemoveAll[T](s, c)
}

func (s *ConcurrentLinkedStack[T]) RetainAll(c Collection[T]) bool {
	return RetainAll[T](s, c)
}

func (s *ConcurrentLinkedStack[T]) Clear() {
	s.head = nil
	s.size.Store(0)
}

func (s *ConcurrentLinkedStack[T]) MutableIterator() MutableIterator[T] {
	return &concurrentLinkedStackIterator[T]{s.head, nil}
}

func (i *concurrentLinkedStackIterator[T]) HasNext() bool {
	return i.node.next != nil
}

func (i *concurrentLinkedStackIterator[T]) Next() (T, exceptions.Exception) {
	if i.node == nil {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	}
	value := i.node.value
	i.node = i.node.next
	return value, nil
}

func (i *concurrentLinkedStackIterator[T]) Remove() exceptions.Exception {
	if i.node == nil {
		return exceptions.NewIndexOutOfBound("", nil)
	}
	i.node.deleted = true
	i.node = i.node.next
	i.stack.size.Add(-1)
	i.stack.deleted.Add(1)
	i.stack.CleanDeleted()
	return nil
}
