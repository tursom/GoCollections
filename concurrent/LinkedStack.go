package concurrent

import (
	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"github.com/tursom/GoCollections/lang/atomic"
)

type (
	LinkedStack[T lang.Object] struct {
		lang.BaseObject
		size    atomic.Int32
		deleted atomic.Int32
		head    *linkedStackNode[T]
	}

	linkedStackNode[T any] struct {
		deleted bool
		next    *linkedStackNode[T]
		value   T
	}

	linkedStackIterator[T lang.Object] struct {
		node  *linkedStackNode[T]
		stack *LinkedStack[T]
	}
)

func (s *LinkedStack[T]) String() string {
	return collections.String[T](s)
}

func NewLinkedStack[T lang.Object]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

func (s *LinkedStack[T]) Iterator() collections.Iterator[T] {
	return s.MutableIterator()
}

func (s *LinkedStack[T]) Push(element T) exceptions.Exception {
	s.push(element)
	return nil
}

func (s *LinkedStack[T]) PushAndGetNode(element T) collections.StackNode[T] {
	return &linkedStackIterator[T]{
		node:  s.push(element),
		stack: s,
	}
}

func (s *LinkedStack[T]) push(element T) *linkedStackNode[T] {
	newNode := &linkedStackNode[T]{value: element, next: s.head}
	for !atomic.CompareAndSwapPointer(&s.head, newNode.next, newNode) {
		newNode.next = s.head
	}
	s.size.Add(1)
	return newNode
}

func (s *LinkedStack[T]) Pop() (T, exceptions.Exception) {
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

func (s *LinkedStack[T]) CleanDeleted() {
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

func (s *LinkedStack[T]) Size() int {
	return int(s.size.Load())
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *LinkedStack[T]) Contains(element T) bool {
	return collections.Contains[T](s, element)
}

func (s *LinkedStack[T]) ContainsAll(c collections.Collection[T]) bool {
	return collections.ContainsAll[T](s, c)
}

func (s *LinkedStack[T]) Add(element T) bool {
	exception := s.Push(element)
	exceptions.Print(exception)
	return exception == nil
}

func (s *LinkedStack[T]) Remove(element T) exceptions.Exception {
	return collections.Remove[T](s, element)
}

func (s *LinkedStack[T]) AddAll(c collections.Collection[T]) bool {
	return collections.AddAll[T](s, c)
}

func (s *LinkedStack[T]) RemoveAll(c collections.Collection[T]) bool {
	return collections.RemoveAll[T](s, c)
}

func (s *LinkedStack[T]) RetainAll(c collections.Collection[T]) bool {
	return collections.RetainAll[T](s, c)
}

func (s *LinkedStack[T]) Clear() {
	s.head = nil
	s.size.Store(0)
}

func (s *LinkedStack[T]) MutableIterator() collections.MutableIterator[T] {
	return &linkedStackIterator[T]{s.head, nil}
}

func (i *linkedStackIterator[T]) HasNext() bool {
	return i.node.next != nil
}

func (i *linkedStackIterator[T]) Next() (T, exceptions.Exception) {
	if i.node == nil {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	}
	value := i.node.value
	i.node = i.node.next
	return value, nil
}

func (i *linkedStackIterator[T]) Remove() exceptions.Exception {
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

func (i *linkedStackIterator[T]) Get() (T, exceptions.Exception) {
	return i.node.value, nil
}

func (i *linkedStackIterator[T]) Set(value T) exceptions.Exception {
	i.node.value = value
	return nil
}

func (i *linkedStackIterator[T]) RemoveAndGet() (T, exceptions.Exception) {
	i.node.deleted = true
	return i.node.value, nil
}
