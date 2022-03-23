package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type SubList[T lang.Object] struct {
	lang.BaseObject
	list     List[T]
	from, to int
}

func NewSubList[T lang.Object](list List[T], from, to int) *SubList[T] {
	return &SubList[T]{lang.NewBaseObject(), list, from, to}
}

func (s *SubList[T]) Iterator() Iterator[T] {
	iterator := s.list.Iterator()
	for i := 0; i < int(s.from); i++ {
		_, err := iterator.Next()
		if err != nil {
			return nil
		}
	}
	return iterator
}

func (s *SubList[T]) Size() int {
	return s.to - s.from
}

func (s *SubList[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SubList[T]) Contains(element T) bool {
	return Contains[T](s, element)
}

func (s *SubList[T]) ContainsAll(c Collection[T]) bool {
	return ContainsAll[T](s, c)
}

func (s *SubList[T]) Get(index int) (T, exceptions.Exception) {
	return s.list.Get(index + s.from)
}

func (s *SubList[T]) SubList(from, to int) List[T] {
	return NewSubList[T](s, from, to)
}
