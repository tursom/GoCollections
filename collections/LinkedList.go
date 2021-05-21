package collections

import (
	"github.com/tursom/GoCollections/exceptions"
)

type (
	LinkedList struct {
		head *linkedListNode
		size uint32
	}
	linkedListNode struct {
		prev, next *linkedListNode
		value      interface{}
	}
)

func (l LinkedList) String() string {
	return String(l)
}

func NewLinkedList() *LinkedList {
	tail := &linkedListNode{}
	tail.prev = tail
	tail.next = tail
	return &LinkedList{tail, 0}
}

func (l LinkedList) Size() uint32 {
	return l.size
}

func (l LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l LinkedList) Contains(element interface{}) bool {
	return Contains(&l, element)
}

func (l LinkedList) ContainsAll(c Collection) bool {
	return ContainsAll(&l, c)
}

func (l *LinkedList) Add(element interface{}) bool {
	l.size++
	node := &linkedListNode{l.head.prev, l.head, element}
	node.next.prev = node
	node.prev.next = node
	return true
}

func (l *LinkedList) Remove(element interface{}) exceptions.Exception {
	err := LoopMutable(l, func(e interface{}, iterator MutableIterator) exceptions.Exception {
		if element == e {
			iterator.Remove()
			return exceptions.CollectionLoopFinished
		}
		return nil
	})
	if err != nil {
		return nil
	} else {
		return exceptions.NewElementNotFoundException("", true)
	}
}

func (l *LinkedList) AddAll(c Collection) bool {
	return AddAll(l, c)
}

func (l *LinkedList) RemoveAll(c Collection) bool {
	return RemoveAll(l, c)
}

func (l *LinkedList) RetainAll(c Collection) bool {
	return RetainAll(l, c)
}

func (l *LinkedList) Clear() {
	l.head.next = l.head
	l.size = 0
}

func (l LinkedList) SubList(from, to uint32) List {
	return NewSubList(l, from, to)
}

func (l *LinkedList) Set(index uint32, element interface{}) exceptions.Exception {
	node := l.head
	for node != l.head {
		if index == 0 {
			node.value = element
			return nil
		}
		index--
	}
	return exceptions.NewIndexOutOfBound("", true)
}

func (l *LinkedList) AddAtIndex(index uint32, element interface{}) bool {
	node := l.head
	for node != l.head {
		if index == 0 {
			l.size++
			node.next = &linkedListNode{node, node.next, element}
			node.next.next.prev = node.next
			return true
		}
		index--
	}
	return false
}

func (l *LinkedList) RemoveAt(index uint32) exceptions.Exception {
	node := l.head
	for node != l.head {
		if index == 0 {
			l.size--
			node.remove()
			return nil
		}
		index--
	}
	return exceptions.NewIndexOutOfBound("", true)
}

func (l *LinkedList) SubMutableList(from, to uint32) MutableList {
	return NewMutableSubList(l, from, to)
}

func (l LinkedList) Get(index uint32) (interface{}, exceptions.Exception) {
	node := l.head
	for node != l.head {
		if index == 0 {
			return node.value, nil
		}
		index--
	}

	return nil, exceptions.NewIndexOutOfBound("", true)
}

func (l LinkedList) Iterator() Iterator {
	return l.MutableIterator()
}

func (l *LinkedList) MutableIterator() MutableIterator {
	return &linkedListIterator{l, l.head.next, l.head}
}

type linkedListIterator struct {
	list       *LinkedList
	node, head *linkedListNode
}

func (l *linkedListIterator) HasNext() bool {
	return l.node != l.head
}

func (l *linkedListIterator) Next() (interface{}, exceptions.Exception) {
	if l.node == l.head {
		return nil, exceptions.NewIndexOutOfBound("", true)
	}
	l.node = l.node.next
	return l.node.prev.value, nil
}

func (l *linkedListIterator) Remove() exceptions.Exception {
	if l.node.prev == l.head {
		return exceptions.NewIndexOutOfBound("", true)
	}
	l.node.prev.remove()
	l.list.size--
	return nil
}

func (l *linkedListNode) remove() {
	l.next.prev = l.prev
	l.prev.next = l.next
}
