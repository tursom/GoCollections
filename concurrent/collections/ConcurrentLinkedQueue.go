package collections

import (
	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"github.com/tursom/GoCollections/lang/atomic"
)

type (
	ConcurrentLinkedQueue[T lang.Object] struct {
		lang.BaseObject
		ConcurrentLinkedStack[T]
		end *linkedStackNode[T]
	}

	linkedQueueIterator[T lang.Object] struct {
		node  *linkedStackNode[T]
		queue *ConcurrentLinkedQueue[T]
	}
)

func (q *ConcurrentLinkedQueue[T]) String() string {
	return collections.String[T](q)
}

func NewLinkedQueue[T lang.Object]() *ConcurrentLinkedQueue[T] {
	return &ConcurrentLinkedQueue[T]{}
}

func (q *ConcurrentLinkedQueue[T]) Iterator() collections.Iterator[T] {
	return q.MutableIterator()
}

func (q *ConcurrentLinkedQueue[T]) Offer(element T) exceptions.Exception {
	_, err := q.offerAndGetNode(element)
	return err
}

func (q *ConcurrentLinkedQueue[T]) OfferAndGetNode(element T) (collections.QueueNode[T], exceptions.Exception) {
	newNode, err := q.offerAndGetNode(element)
	if err != nil {
		return nil, err
	}
	return &linkedQueueIterator[T]{queue: q, node: newNode}, nil
}

func (q *ConcurrentLinkedQueue[T]) offerAndGetNode(element T) (*linkedStackNode[T], exceptions.Exception) {
	newNode := &linkedStackNode[T]{value: element}
	q.size.Add(1)

	var next **linkedStackNode[T]
	ref := q.end
	switch {
	case ref == nil:
		next = &q.head
	default:
		next = &ref.next
	}
	for !atomic.CompareAndSwapPointer(next, nil, newNode) {
		if ref == nil || ref.next == nil {
			next = &q.head
			ref = q.head
		} else {
			for ref.next != nil {
				ref = ref.next
			}
			next = &ref.next
		}
	}
	q.end = newNode
	return newNode, nil
}

func (q *ConcurrentLinkedQueue[T]) Poll() (T, exceptions.Exception) {
	return q.Pop()
}

func (q *ConcurrentLinkedQueue[T]) MutableIterator() collections.MutableIterator[T] {
	return &linkedQueueIterator[T]{queue: q, node: q.head}
}

func (q *ConcurrentLinkedQueue[T]) Size() int {
	return int(q.size.Load())
}

func (q *ConcurrentLinkedQueue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *ConcurrentLinkedQueue[T]) Contains(element T) bool {
	return collections.Contains[T](q, element)
}

func (q *ConcurrentLinkedQueue[T]) ContainsAll(collection collections.Collection[T]) bool {
	return collections.ContainsAll[T](q, collection)
}

func (q *ConcurrentLinkedQueue[T]) Add(element T) bool {
	exception := q.Push(element)
	exceptions.Print(exception)
	return exception == nil
}

func (q *ConcurrentLinkedQueue[T]) Remove(element T) exceptions.Exception {
	return collections.Remove[T](q, element)
}

func (q *ConcurrentLinkedQueue[T]) AddAll(collection collections.Collection[T]) bool {
	return collections.AddAll[T](q, collection)
}

func (q *ConcurrentLinkedQueue[T]) RemoveAll(collection collections.Collection[T]) bool {
	return collections.RemoveAll[T](q, collection)
}

func (q *ConcurrentLinkedQueue[T]) RetainAll(collection collections.Collection[T]) bool {
	return collections.RetainAll[T](q, collection)
}

func (q *ConcurrentLinkedQueue[T]) Clear() {
	q.head = nil
	q.end = nil
	q.size.Store(0)
}

func (i *linkedQueueIterator[T]) HasNext() bool {
	for i.node != nil && i.node.deleted {
		i.node = i.node.next
	}
	return i.node != nil
}

func (i *linkedQueueIterator[T]) Next() (T, exceptions.Exception) {
	value, err := i.Get()
	i.node = i.node.next
	for i.node != nil && i.node.deleted {
		i.node = i.node.next
	}
	return value, err
}

func (i *linkedQueueIterator[T]) Get() (T, exceptions.Exception) {
	return i.node.value, nil
}

func (i *linkedQueueIterator[T]) Set(value T) exceptions.Exception {
	i.node.value = value
	return nil
}

func (i *linkedQueueIterator[T]) Remove() exceptions.Exception {
	_, err := i.RemoveAndGet()
	return err
}

func (i *linkedQueueIterator[T]) RemoveAndGet() (T, exceptions.Exception) {
	if i.node == nil {
		return lang.Nil[T](), nil
	}
	load := i.node
	load.deleted = true
	i.queue.size.Add(-1)
	i.queue.deleted.Add(1)
	i.queue.CleanDeleted()
	i.node = load.next
	return load.value, nil
}
