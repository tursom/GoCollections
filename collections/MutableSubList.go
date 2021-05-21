package collections

import "github.com/tursom/GoCollections/exceptions"

type MutableSubList struct {
	list     MutableList
	from, to uint32
}

func NewMutableSubList(list MutableList, from, to uint32) *MutableSubList {
	return &MutableSubList{list, from, to}
}

func (s *MutableSubList) Iterator() Iterator {
	iterator := s.list.Iterator()
	for i := 0; i < int(s.from); i++ {
		_, err := iterator.Next()
		if err != nil {
			return nil
		}
	}
	return iterator
}

func (s *MutableSubList) Size() uint32 {
	return s.to - s.from
}

func (s *MutableSubList) IsEmpty() bool {
	return s.Size() == 0
}

func (s *MutableSubList) Contains(element interface{}) bool {
	return Contains(s, element)
}

func (s *MutableSubList) ContainsAll(c Collection) bool {
	return ContainsAll(s, c)
}

func (s *MutableSubList) Get(index uint32) (interface{}, error) {
	return s.list.Get(index + s.from)
}

func (s *MutableSubList) SubList(from, to uint32) List {
	return NewSubList(s, from, to)
}

func (s *MutableSubList) MutableIterator() MutableIterator {
	return nil
}

func (s *MutableSubList) Add(_ interface{}) bool {
	return false
}

func (s *MutableSubList) Remove(element interface{}) error {
	return nil
}

func (s *MutableSubList) AddAll(_ Collection) bool {
	return false
}

func (s *MutableSubList) RemoveAll(_ Collection) bool {
	return false
}

func (s *MutableSubList) RetainAll(_ Collection) bool {
	return false
}

func (s *MutableSubList) Clear() {
}

func (s *MutableSubList) Set(index uint32, element interface{}) error {
	if index >= s.to-s.from {
		return exceptions.NewIndexOutOfBound("", true)
	}
	return s.list.Set(index+s.from, element)
}

func (s *MutableSubList) AddAtIndex(_ uint32, _ interface{}) bool {
	return false
}

func (s *MutableSubList) RemoveAt(index uint32) error {
	return exceptions.NewOperationNotSupportedException("", true)
}

func (s *MutableSubList) SubMutableList(from, to uint32) MutableList {
	return NewMutableSubList(s.list, s.from+from, to)
}
