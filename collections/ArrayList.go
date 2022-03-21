package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type ArrayList[T lang.Object] struct {
	array []T
	used  int
}

func NewArrayList[T lang.Object]() *ArrayList[T] {
	return &ArrayList[T]{
		make([]T, 16),
		0,
	}
}

func NewArrayListByCapacity[T lang.Object](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		make([]T, cap),
		0,
	}
}

func (a ArrayList[T]) String() string {
	return String[T](a)
}

func (a ArrayList[T]) Iterator() Iterator[T] {
	return a.MutableIterator()
}

func (a ArrayList[T]) Size() int {
	return a.used
}

func (a ArrayList[T]) IsEmpty() bool {
	return a.Size() == 0
}

func (a ArrayList[T]) Contains(element T) bool {
	return Contains[T](a, element)
}

func (a ArrayList[T]) ContainsAll(c Collection[T]) bool {
	return ContainsAll[T](a, c)
}

func (a *ArrayList[T]) Add(element T) bool {
	if a.used >= len(a.array) {
		oldArray := a.array
		a.array = make([]T, a.used*2)
		copy(a.array, oldArray)
	}
	a.array[a.used] = element
	a.used++
	return true
}

func (a ArrayList[T]) IndexOf(element T) int {
	for i := 0; i < a.used; i++ {
		if lang.Equals(element, a.array[i]) {
			return int(i)
		}
	}
	return -1
}

func (a *ArrayList[T]) Remove(element T) exceptions.Exception {
	index := a.IndexOf(element)
	if index < 0 {
		return exceptions.NewElementNotFoundException("", nil)
	} else {
		return a.RemoveAt(int(index))
	}
}

func (a *ArrayList[T]) AddAll(c Collection[T]) bool {
	return AddAll[T](a, c)
}

func (a *ArrayList[T]) RemoveAll(c Collection[T]) bool {
	return RemoveAll[T](a, c)
}

func (a *ArrayList[T]) RetainAll(c Collection[T]) bool {
	return RetainAll[T](a, c)
}

func (a *ArrayList[T]) Clear() {
	a.used = 0
}

func (a ArrayList[T]) Get(index int) (T, exceptions.Exception) {
	if index >= a.used {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	} else {
		return a.array[index], nil
	}
}

func (a ArrayList[T]) SubList(from, to int) List[T] {
	return NewSubList[T](a, from, to)
}

func (a *ArrayList[T]) Set(index int, element T) exceptions.Exception {
	if index >= a.used {
		return exceptions.NewIndexOutOfBound("", nil)
	}
	a.array[index] = element
	return nil
}

func (a *ArrayList[T]) AddAtIndex(index int, element T) bool {
	if !a.Add(element) {
		return false
	}

	array := a.array
	for i := a.used - 1; i > index; i++ {
		array[i] = array[i-1]
	}
	array[index] = element
	a.used++
	return true
}

func (a *ArrayList[T]) RemoveAt(index int) exceptions.Exception {
	if index >= a.used {
		return exceptions.NewIndexOutOfBound("", nil)
	}

	array := a.array
	for i := index + 1; i < a.used; i++ {
		array[i-1] = array[i]
	}
	a.used--
	return nil
}

func (a *ArrayList[T]) SubMutableList(from, to int) MutableList[T] {
	panic("implement me")
}

func (a *ArrayList[T]) MutableIterator() MutableIterator[T] {
	return &arrayListIterator[T]{a, 0}
}

type arrayListIterator[T lang.Object] struct {
	arrayList *ArrayList[T]
	index     int
}

func (a *arrayListIterator[T]) HasNext() bool {
	return a.index < a.arrayList.used
}

func (a *arrayListIterator[T]) Next() (T, exceptions.Exception) {
	value, err := a.arrayList.Get(a.index)
	if err != nil {
		return lang.Nil[T](), err
	}
	a.index++
	return value, nil
}

func (a *arrayListIterator[T]) Remove() exceptions.Exception {
	err := a.arrayList.RemoveAt(a.index - 1)
	if err != nil {
		return err
	}
	a.index--
	return nil
}
