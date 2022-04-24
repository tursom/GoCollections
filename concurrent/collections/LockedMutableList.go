package collections

import (
	"sync"

	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/concurrent"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	LockedMutableList[T lang.Object] struct {
		list collections.MutableList[T]
		lock concurrent.RWLock
	}
	lockedMutableListIterator[T lang.Object] struct {
		iterator collections.MutableIterator[T]
		lock     concurrent.RWLock
	}
)

func MutableListWithLock[T lang.Object](list collections.MutableList[T]) collections.MutableList[T] {
	return &LockedMutableList[T]{
		list: list,
		lock: &sync.RWMutex{},
	}
}

func (l *LockedMutableList[T]) String() string {
	return collections.String[T](l)
}

func (l *LockedMutableList[T]) Iterator() collections.Iterator[T] {
	return l.MutableIterator()
}

func (l *LockedMutableList[T]) Size() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.list.Size()
}

func (l *LockedMutableList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *LockedMutableList[T]) Contains(element T) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.list.Contains(element)
}

func (l *LockedMutableList[T]) ContainsAll(c collections.Collection[T]) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.list.ContainsAll(c)
}

func (l *LockedMutableList[T]) MutableIterator() collections.MutableIterator[T] {
	return &lockedMutableListIterator[T]{l.list.MutableIterator(), l.lock}
}

func (l *LockedMutableList[T]) Add(element T) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.list.Add(element)
}

func (l *LockedMutableList[T]) Remove(element T) exceptions.Exception {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.list.Remove(element)
}

func (l *LockedMutableList[T]) AddAll(c collections.Collection[T]) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.list.AddAll(c)
}

func (l *LockedMutableList[T]) RemoveAll(c collections.Collection[T]) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.list.RemoveAll(c)
}

func (l *LockedMutableList[T]) RetainAll(c collections.Collection[T]) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.list.RetainAll(c)
}

func (l *LockedMutableList[T]) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.list.Clear()
}

func (l *LockedMutableList[T]) Get(index int) (T, exceptions.Exception) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.list.Get(index)
}

func (l *LockedMutableList[T]) SubList(from, to int) collections.List[T] {
	return l.SubMutableList(from, to)
}

func (l *LockedMutableList[T]) Set(index int, element T) exceptions.Exception {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.list.Set(index, element)
}

func (l *LockedMutableList[T]) AddAtIndex(index int, element T) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.list.AddAtIndex(index, element)
}

func (l *LockedMutableList[T]) RemoveAt(index int) exceptions.Exception {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.list.RemoveAt(index)
}

func (l *LockedMutableList[T]) SubMutableList(from, to int) collections.MutableList[T] {
	return &LockedMutableList[T]{l.list.SubMutableList(from, to), l.lock}
}

func (l *lockedMutableListIterator[T]) HasNext() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.iterator.HasNext()
}

func (l *lockedMutableListIterator[T]) Next() (T, exceptions.Exception) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.iterator.Next()
}

func (l *lockedMutableListIterator[T]) Remove() exceptions.Exception {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.iterator.Remove()
}

func (l *LockedMutableList[T]) ListIterator() collections.ListIterator[T] {
	//TODO implement me
	panic("implement me")
}

func (l *LockedMutableList[T]) MutableListIterator() collections.MutableListIterator[T] {
	//TODO implement me
	panic("implement me")
}
