package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"github.com/tursom/GoCollections/lang/atomic"
)

type (
	ConcurrentLinkedQueueNode[T lang.Object] interface {
		Get() (T, exceptions.Exception)
		Remove() exceptions.Exception
		RemoveAndGet() (T, exceptions.Exception)
	}

	ConcurrentLinkedQueue[T lang.Object] struct {
		lang.BaseObject
		ConcurrentLinkedStack[T]
		end *concurrentLinkedStackNode[T]
	}

	concurrentLinkedQueueIterator[T lang.Object] struct {
		node  *concurrentLinkedStackNode[T]
		queue *ConcurrentLinkedQueue[T]
	}
)

func (q *ConcurrentLinkedQueue[T]) String() string {
	return String[T](q)
}

func NewConcurrentLinkedQueue[T lang.Object]() *ConcurrentLinkedQueue[T] {
	return &ConcurrentLinkedQueue[T]{}
}

func (q *ConcurrentLinkedQueue[T]) Iterator() Iterator[T] {
	return q.MutableIterator()
}

func (q *ConcurrentLinkedQueue[T]) Offer(element T) exceptions.Exception {
	_, err := q.offerAndGetNode(element)
	return err
}

func (q *ConcurrentLinkedQueue[T]) OfferAndGetNode(element T) (ConcurrentLinkedQueueNode[T], exceptions.Exception) {
	newNode, err := q.offerAndGetNode(element)
	if err != nil {
		return nil, err
	}
	return &concurrentLinkedQueueIterator[T]{queue: q, node: newNode}, nil
}

func (q *ConcurrentLinkedQueue[T]) offerAndGetNode(element T) (*concurrentLinkedStackNode[T], exceptions.Exception) {
	newNode := &concurrentLinkedStackNode[T]{value: element}
	q.size.Add(1)

	var next **concurrentLinkedStackNode[T]
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

func (q *ConcurrentLinkedQueue[T]) MutableIterator() MutableIterator[T] {
	return &concurrentLinkedQueueIterator[T]{queue: q, node: q.head}
}

func (q *ConcurrentLinkedQueue[T]) Size() int {
	return int(q.size.Load())
}

func (q *ConcurrentLinkedQueue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *ConcurrentLinkedQueue[T]) Contains(element T) bool {
	return Contains[T](q, element)
}

func (q *ConcurrentLinkedQueue[T]) ContainsAll(collection Collection[T]) bool {
	return ContainsAll[T](q, collection)
}

func (q *ConcurrentLinkedQueue[T]) Add(element T) bool {
	exception := q.Push(element)
	exceptions.Print(exception)
	return exception == nil
}

func (q *ConcurrentLinkedQueue[T]) Remove(element T) exceptions.Exception {
	return Remove[T](q, element)
}

func (q *ConcurrentLinkedQueue[T]) AddAll(collection Collection[T]) bool {
	return AddAll[T](q, collection)
}

func (q *ConcurrentLinkedQueue[T]) RemoveAll(collection Collection[T]) bool {
	return RemoveAll[T](q, collection)
}

func (q *ConcurrentLinkedQueue[T]) RetainAll(collection Collection[T]) bool {
	return RetainAll[T](q, collection)
}

func (q *ConcurrentLinkedQueue[T]) Clear() {
	q.head = nil
	q.end = nil
	q.size.Store(0)
}

func (i *concurrentLinkedQueueIterator[T]) HasNext() bool {
	for i.node != nil && i.node.deleted {
		i.node = i.node.next
	}
	return i.node != nil
}

func (i *concurrentLinkedQueueIterator[T]) Next() (T, exceptions.Exception) {
	value, err := i.Get()
	i.node = i.node.next
	for i.node != nil && i.node.deleted {
		i.node = i.node.next
	}
	return value, err
}

func (i *concurrentLinkedQueueIterator[T]) Get() (T, exceptions.Exception) {
	return (*i.node).value, nil
}

func (i *concurrentLinkedQueueIterator[T]) Remove() exceptions.Exception {
	_, err := i.RemoveAndGet()
	return err
}

func (i *concurrentLinkedQueueIterator[T]) RemoveAndGet() (T, exceptions.Exception) {
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
