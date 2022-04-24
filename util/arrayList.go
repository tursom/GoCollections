package util

import (
	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type arrayList[T lang.Object] struct {
	lang.BaseObject
	array []T
}

func (a *arrayList[T]) ListIterator() collections.ListIterator[T] {
	return &arrayListIterator[T]{a.array, 0}
}

func (a *arrayList[T]) Iterator() collections.Iterator[T] {
	return a.ListIterator()
}

func (a *arrayList[T]) Size() int {
	return lang.Len(a.array)
}

func (a *arrayList[T]) IsEmpty() bool {
	return a.Size() == 0
}

func (a *arrayList[T]) Contains(element T) bool {
	return collections.Contains[T](a, element)
}

func (a *arrayList[T]) ContainsAll(c collections.Collection[T]) bool {
	return collections.ContainsAll[T](a, c)
}

func (a *arrayList[T]) Get(index int) (T, exceptions.Exception) {
	return CheckedGet(a.array, index)
}

func (a *arrayList[T]) SubList(from, to int) collections.List[T] {
	return &arrayList[T]{array: a.array[from:to]}
}

type arrayListIterator[T lang.Object] struct {
	array []T
	index int
}

func (a *arrayListIterator[T]) HasPrevious() bool {
	return a.index > 0
}

func (a *arrayListIterator[T]) Previous() (T, exceptions.Exception) {
	a.index--
	return a.array[a.index], nil
}

func (a *arrayListIterator[T]) NextIndex() int {
	return a.index
}

func (a *arrayListIterator[T]) PreviousIndex() int {
	return a.index - 1
}

func (a *arrayListIterator[T]) HasNext() bool {
	return a.index < lang.Len(a.array)
}

func (a *arrayListIterator[T]) Next() (r T, err exceptions.Exception) {
	defer func() {
		r := recover()
		if r != nil {
			err = exceptions.NewNPE("", exceptions.Cfg().SetCause(r))
		}
	}()
	i := a.index
	a.index++
	r = a.array[i]
	return
}
