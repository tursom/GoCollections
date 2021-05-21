package collections

import "github.com/tursom/GoCollections/exceptions"

type ArrayList struct {
	array []interface{}
	used  uint32
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		make([]interface{}, 16),
		0,
	}
}

func (a ArrayList) Iterator() Iterator {
	return a.MutableIterator()
}

func (a ArrayList) Size() uint32 {
	return a.used
}

func (a ArrayList) IsEmpty() bool {
	return a.Size() == 0
}

func (a ArrayList) Contains(element interface{}) bool {
	return Contains(a, element)
}

func (a ArrayList) ContainsAll(c Collection) bool {
	return ContainsAll(a, c)
}

func (a *ArrayList) MutableIterator() MutableIterator {
	panic("implement me")
}

func (a *ArrayList) Add(element interface{}) bool {
	if a.used >= uint32(len(a.array)) {
		oldArray := a.array
		a.array = make([]interface{}, a.used*2)
		copy(a.array, oldArray)
	}
	a.array[a.used] = element
	a.used++
	return true
}

func (a ArrayList) IndexOf(element interface{}) int {
	var i uint32 = 0
	for ; i < a.used; i++ {
		if a.array[i] == element {
			return int(i)
		}
	}
	return -1
}

func (a *ArrayList) RemoveIndex(index uint32) error {
	if index >= a.used {
		return exceptions.NewIndexOutOfBound("", true)
	}

	array := a.array
	for i := index + 1; i < a.used; i++ {
		array[index-1] = array[index]
	}
	return nil
}

func (a *ArrayList) Remove(element interface{}) error {
	index := a.IndexOf(element)
	if index < 0 {
		return exceptions.NewElementNotFoundException("", true)
	} else {
		return a.RemoveIndex(uint32(index))
	}
}

func (a *ArrayList) AddAll(c Collection) bool {
	return AddAll(a, c)
}

func (a *ArrayList) RemoveAll(c Collection) bool {
	return RemoveAll(a, c)
}

func (a *ArrayList) RetainAll(c Collection) bool {
	return RetainAll(a, c)
}

func (a *ArrayList) Clear() {
	a.used = 0
}

func (a ArrayList) Get(index uint32) (interface{}, error) {
	if index >= a.used {
		return nil, exceptions.NewIndexOutOfBound("", true)
	} else {
		return a.array[index], nil
	}
}

func (a ArrayList) SubList(from, to uint32) List {
	return NewSubList(a, from, to)
}

func (a *ArrayList) Set(index uint32, element interface{}) error {
	if index >= a.used {
		return exceptions.NewIndexOutOfBound("", true)
	}
	a.array[index] = element
	return nil
}

func (a *ArrayList) AddAtIndex(index uint32, element interface{}) bool {
	panic("implement me")
}

func (a *ArrayList) RemoveAt(index uint32) bool {
	panic("implement me")
}

func (a *ArrayList) SubMutableList(from, to uint32) MutableList {
	panic("implement me")
}
