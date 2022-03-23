package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type MutableSubList[T lang.Object] struct {
	lang.BaseObject
	list     MutableList[T]
	from, to int
}

func NewMutableSubList[T lang.Object](list MutableList[T], from, to int) *MutableSubList[T] {
	return &MutableSubList[T]{lang.NewBaseObject(), list, from, to}
}

func (s *MutableSubList[T]) Iterator() Iterator[T] {
	iterator := s.list.Iterator()
	for i := 0; i < s.from; i++ {
		_, err := iterator.Next()
		if err != nil {
			return nil
		}
	}
	return iterator
}

func (s *MutableSubList[T]) Size() int {
	return s.to - s.from
}

func (s *MutableSubList[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *MutableSubList[T]) Contains(element T) bool {
	return Contains[T](s, element)
}

func (s *MutableSubList[T]) ContainsAll(c Collection[T]) bool {
	return ContainsAll[T](s, c)
}

func (s *MutableSubList[T]) Get(index int) (T, exceptions.Exception) {
	return s.list.Get(index + s.from)
}

func (s *MutableSubList[T]) SubList(from, to int) List[T] {
	return NewSubList[T](s, from, to)
}

func (s *MutableSubList[T]) MutableIterator() MutableIterator[T] {
	return nil
}

func (s *MutableSubList[T]) Add(_ T) bool {
	return false
}

func (s *MutableSubList[T]) Remove(element T) exceptions.Exception {
	return nil
}

func (s *MutableSubList[T]) AddAll(_ Collection[T]) bool {
	return false
}

func (s *MutableSubList[T]) RemoveAll(_ Collection[T]) bool {
	return false
}

func (s *MutableSubList[T]) RetainAll(_ Collection[T]) bool {
	return false
}

func (s *MutableSubList[T]) Clear() {
}

func (s *MutableSubList[T]) Set(index int, element T) exceptions.Exception {
	if index >= s.to-s.from {
		return exceptions.NewIndexOutOfBound("", nil)
	}
	return s.list.Set(index+s.from, element)
}

func (s *MutableSubList[T]) AddAtIndex(_ int, _ T) bool {
	return false
}

func (s *MutableSubList[T]) RemoveAt(index int) exceptions.Exception {
	return exceptions.NewOperationNotSupportedException("", nil)
}

func (s *MutableSubList[T]) SubMutableList(from, to int) MutableList[T] {
	return NewMutableSubList[T](s.list, s.from+from, to)
}
