package collections

type SubList struct {
	list     List
	from, to uint32
}

func NewSubList(list List, from, to uint32) *SubList {
	return &SubList{list, from, to}
}

func (s *SubList) Iterator() Iterator {
	iterator := s.list.Iterator()
	for i := 0; i < int(s.from); i++ {
		iterator.Next()
	}
	return iterator
}

func (s *SubList) Size() uint32 {
	return s.to - s.from
}

func (s *SubList) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SubList) Contains(element interface{}) bool {
	return Contains(s, element)
}

func (s *SubList) ContainsAll(c Collection) bool {
	return ContainsAll(s, c)
}

func (s *SubList) Get(index uint32) (interface{}, error) {
	return s.list.Get(index + s.from)
}

func (s *SubList) SubList(from, to uint32) List {
	return NewSubList(s, from, to)
}
