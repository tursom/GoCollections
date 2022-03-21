package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	LinkedList[T lang.Object] struct {
		head *linkedListNode[T]
		size int
	}
	linkedListNode[T lang.Object] struct {
		prev, next *linkedListNode[T]
		value      T
	}
)

func (l LinkedList[T]) String() string {
	return String[T](l)
}

func NewLinkedList[T lang.Object]() *LinkedList[T] {
	tail := &linkedListNode[T]{}
	tail.prev = tail
	tail.next = tail
	return &LinkedList[T]{tail, 0}
}

func (l LinkedList[T]) Size() int {
	return l.size
}

func (l LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l LinkedList[T]) Contains(element T) bool {
	return Contains[T](&l, element)
}

func (l LinkedList[T]) ContainsAll(c Collection[T]) bool {
	return ContainsAll[T](&l, c)
}

func (l *LinkedList[T]) Add(element T) bool {
	l.size++
	node := &linkedListNode[T]{l.head.prev, l.head, element}
	node.next.prev = node
	node.prev.next = node
	return true
}

func (l *LinkedList[T]) Remove(element T) exceptions.Exception {
	err := LoopMutable[T](l, func(e T, iterator MutableIterator[T]) exceptions.Exception {
		if lang.Equals(element, e) {
			iterator.Remove()
			return exceptions.CollectionLoopFinished
		}
		return nil
	})
	if err != nil {
		return nil
	} else {
		return exceptions.NewElementNotFoundException("", nil)
	}
}

func (l *LinkedList[T]) AddAll(c Collection[T]) bool {
	return AddAll[T](l, c)
}

func (l *LinkedList[T]) RemoveAll(c Collection[T]) bool {
	return RemoveAll[T](l, c)
}

func (l *LinkedList[T]) RetainAll(c Collection[T]) bool {
	return RetainAll[T](l, c)
}

func (l *LinkedList[T]) Clear() {
	l.head.next = l.head
	l.size = 0
}

func (l LinkedList[T]) SubList(from, to int) List[T] {
	return NewSubList[T](l, from, to)
}

func (l *LinkedList[T]) Set(index int, element T) exceptions.Exception {
	node := l.head
	for node != l.head {
		if index == 0 {
			node.value = element
			return nil
		}
		index--
	}
	return exceptions.NewIndexOutOfBound("", nil)
}

func (l *LinkedList[T]) AddAtIndex(index int, element T) bool {
	node := l.head
	for node != l.head {
		if index == 0 {
			l.size++
			node.next = &linkedListNode[T]{node, node.next, element}
			node.next.next.prev = node.next
			return true
		}
		index--
	}
	return false
}

func (l *LinkedList[T]) RemoveAt(index int) exceptions.Exception {
	node := l.head
	for node != l.head {
		if index == 0 {
			l.size--
			node.remove()
			return nil
		}
		index--
	}
	return exceptions.NewIndexOutOfBound("", nil)
}

func (l *LinkedList[T]) SubMutableList(from, to int) MutableList[T] {
	return NewMutableSubList[T](l, from, to)
}

func (l LinkedList[T]) Get(index int) (T, exceptions.Exception) {
	node := l.head
	for node != l.head {
		if index == 0 {
			return node.value, nil
		}
		index--
	}

	return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
}

func (l LinkedList[T]) Iterator() Iterator[T] {
	return l.MutableIterator()
}

func (l *LinkedList[T]) MutableIterator() MutableIterator[T] {
	return &linkedListIterator[T]{l, l.head.next, l.head}
}

type linkedListIterator[T lang.Object] struct {
	list       *LinkedList[T]
	node, head *linkedListNode[T]
}

func (l *linkedListIterator[T]) HasNext() bool {
	return l.node != l.head
}

func (l *linkedListIterator[T]) Next() (T, exceptions.Exception) {
	if l.node == l.head {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	}
	l.node = l.node.next
	return l.node.prev.value, nil
}

func (l *linkedListIterator[T]) Remove() exceptions.Exception {
	if l.node.prev == l.head {
		return exceptions.NewIndexOutOfBound("", nil)
	}
	l.node.prev.remove()
	l.list.size--
	return nil
}

func (l *linkedListNode[T]) remove() {
	l.next.prev = l.prev
	l.prev.next = l.next
}
