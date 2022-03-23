package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type ArrayList[T lang.Object] struct {
	lang.BaseObject
	array []T
}

func NewArrayList[T lang.Object]() *ArrayList[T] {
	return NewArrayListByCapacity[T](16)
}

func NewArrayListByCapacity[T lang.Object](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		BaseObject: lang.NewBaseObject(),
		array:      make([]T, 0, cap),
	}
}

func (a ArrayList[T]) String() string {
	return String[T](a)
}

func (a ArrayList[T]) Iterator() Iterator[T] {
	return a.MutableIterator()
}

func (a ArrayList[T]) Size() int {
	return lang.Len(a.array)
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
	a.array = lang.Append(a.array, element)
	return true
}

func (a ArrayList[T]) IndexOf(element T) int {
	for i := 0; i < a.Size(); i++ {
		if lang.Equals(element, a.array[i]) {
			return i
		}
	}
	return -1
}

func (a *ArrayList[T]) Remove(element T) exceptions.Exception {
	index := a.IndexOf(element)
	if index < 0 {
		return exceptions.NewElementNotFoundException("", nil)
	} else {
		return a.RemoveAt(index)
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
	a.array = []T{}
}

func (a ArrayList[T]) Get(index int) (T, exceptions.Exception) {
	if index >= a.Size() {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	} else {
		return a.array[index], nil
	}
}

func (a ArrayList[T]) SubList(from, to int) List[T] {
	return NewSubList[T](a, from, to)
}

func (a *ArrayList[T]) Set(index int, element T) exceptions.Exception {
	if index >= a.Size() {
		return exceptions.NewIndexOutOfBound("", nil)
	}
	a.array[index] = element
	return nil
}

func (a *ArrayList[T]) AddAtIndex(index int, element T) bool {
	if index >= a.Size() {
		return false
	}

	array := a.array
	a.array = lang.Append[T](array[:index], element)
	a.array = lang.Append[T](a.array, array[index:]...)
	return true
}

func (a *ArrayList[T]) RemoveAt(index int) exceptions.Exception {
	if index >= a.Size() {
		return exceptions.NewIndexOutOfBound("", nil)
	}

	a.array = lang.Append[T](a.array[:index], a.array[index+1:]...)
	return nil
}

func (a *ArrayList[T]) SubMutableList(from, to int) MutableList[T] {
	return NewMutableSubList[T](a, from, to)
}

func (a *ArrayList[T]) MutableIterator() MutableIterator[T] {
	return &arrayListIterator[T]{a, 0}
}

func (a *ArrayList[T]) RemoveLast() (T, exceptions.Exception) {
	v, _ := a.Get(a.Size() - 1)
	return v, a.RemoveAt(a.Size() - 1)
}

type arrayListIterator[T lang.Object] struct {
	arrayList *ArrayList[T]
	index     int
}

func (a *arrayListIterator[T]) HasNext() bool {
	return a.index < a.arrayList.Size()
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
