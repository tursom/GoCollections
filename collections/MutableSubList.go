package collections

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
		iterator.Next()
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

func (s *MutableSubList) Add(element interface{}) bool {
	return false
}

func (s *MutableSubList) Remove(element interface{}) bool {
	return false
}

func (s *MutableSubList) AddAll(c Collection) bool {
	return false
}

func (s *MutableSubList) RemoveAll(c Collection) bool {
	return false
}

func (s *MutableSubList) RetainAll(c Collection) bool {
	return false
}

func (s *MutableSubList) Clear() {
}

func (s *MutableSubList) Set(index uint32, element interface{}) bool {
	if index >= s.to-s.from {
		return false
	}
	return s.list.Set(index+s.from, element)
}

func (s *MutableSubList) AddAtIndex(index uint32, element interface{}) bool {
	return false
}

func (s *MutableSubList) RemoveAt(index uint32) bool {
	return false
}

func (s *MutableSubList) SubMutableList(from, to uint32) MutableList {
	return NewMutableSubList(s.list, s.from+from, to)
}
