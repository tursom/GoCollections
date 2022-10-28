package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	ArrayList[T lang.Object] struct {
		lang.BaseObject
		array []T
	}
	arrayListIterator[T lang.Object] struct {
		arrayList *ArrayList[T]
		index     int
	}
)

func NewArrayList[T lang.Object]() *ArrayList[T] {
	return NewArrayListByCapacity[T](16)
}

func NewArrayListByCapacity[T lang.Object](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		BaseObject: lang.NewBaseObject(),
		array:      make([]T, 0, cap),
	}
}

// NewArrayListFrom create a new ArrayList from list by index from [from] until [to]
func NewArrayListFrom[T lang.Object](list List[T], from, to int) *ArrayList[T] {
	newList := NewArrayListByCapacity[T](to - from)
	iterator, err := SkipIterator[T](list.ListIterator(), to-from)
	if err != nil {
		panic(err)
	}

	for i := 0; i < to-from; i++ {
		next, err := iterator.Next()
		if err != nil {
			panic(err)
		}

		// newList wont throw any exception in this place
		_ = newList.Add(next)
	}

	return newList
}

func (a *ArrayList[T]) String() string {
	return String[T](a)
}

func (a *ArrayList[T]) Iterator() Iterator[T] {
	return a.MutableIterator()
}

func (a *ArrayList[T]) ListIterator() ListIterator[T] {
	return a.MutableListIterator()
}

func (a *ArrayList[T]) MutableListIterator() MutableListIterator[T] {
	return &arrayListIterator[T]{a, 0}
}

func (a *ArrayList[T]) Size() int {
	return lang.Len(a.array)
}

func (a *ArrayList[T]) IsEmpty() bool {
	return a.Size() == 0
}

func (a *ArrayList[T]) Contains(element T) bool {
	return Contains[T](a, element)
}

func (a *ArrayList[T]) ContainsAll(c Collection[T]) bool {
	return ContainsAll[T](a, c)
}

func (a *ArrayList[T]) Add(element T) bool {
	a.array = lang.Append(a.array, element)
	return true
}

func (a *ArrayList[T]) IndexOf(element T) int {
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

func (a *ArrayList[T]) Get(index int) (T, exceptions.Exception) {
	if index >= a.Size() {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	} else {
		return a.array[index], nil
	}
}

func (a *ArrayList[T]) SubList(from, to int) List[T] {
	return a.SubMutableList(from, to)
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
	return &ArrayList[T]{
		array: a.array[from:to],
	}
}

func (a *ArrayList[T]) MutableIterator() MutableIterator[T] {
	return &arrayListIterator[T]{a, 0}
}

func (a *ArrayList[T]) RemoveLast() (T, exceptions.Exception) {
	v, _ := a.Get(a.Size() - 1)
	return v, a.RemoveAt(a.Size() - 1)
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

func (a *arrayListIterator[T]) HasPrevious() bool {
	return a.index > 0
}

func (a *arrayListIterator[T]) Previous() (T, exceptions.Exception) {
	if a.index <= 0 || a.index >= len(a.arrayList.array) {
		return lang.Nil[T](), exceptions.NewIndexOutOfBound("", nil)
	}
	a.index--
	return a.arrayList.array[a.index], a.arrayList.RemoveAt(a.index)
}

func (a *arrayListIterator[T]) NextIndex() int {
	return a.index
}

func (a *arrayListIterator[T]) PreviousIndex() int {
	return a.index - 1
}

func (a *arrayListIterator[T]) Set(value T) exceptions.Exception {
	if a.index <= 0 {
		return exceptions.NewIndexOutOfBound("", nil)
	}
	a.arrayList.array[a.index-1] = value
	return nil
}

func (a *arrayListIterator[T]) Add(value T) exceptions.Exception {
	a.arrayList.AddAtIndex(a.index, value)
	a.index++
	return nil
}
