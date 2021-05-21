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

func (a ArrayList) String() string {
	return String(a)
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

func (a *ArrayList) Remove(element interface{}) exceptions.Exception {
	index := a.IndexOf(element)
	if index < 0 {
		return exceptions.NewElementNotFoundException("", true)
	} else {
		return a.RemoveAt(uint32(index))
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

func (a ArrayList) Get(index uint32) (interface{}, exceptions.Exception) {
	if index >= a.used {
		return nil, exceptions.NewIndexOutOfBound("", true)
	} else {
		return a.array[index], nil
	}
}

func (a ArrayList) SubList(from, to uint32) List {
	return NewSubList(a, from, to)
}

func (a *ArrayList) Set(index uint32, element interface{}) exceptions.Exception {
	if index >= a.used {
		return exceptions.NewIndexOutOfBound("", true)
	}
	a.array[index] = element
	return nil
}

func (a *ArrayList) AddAtIndex(index uint32, element interface{}) bool {
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

func (a *ArrayList) RemoveAt(index uint32) exceptions.Exception {
	if index >= a.used {
		return exceptions.NewIndexOutOfBound("", true)
	}

	array := a.array
	for i := index + 1; i < a.used; i++ {
		array[i-1] = array[i]
	}
	a.used--
	return nil
}

func (a *ArrayList) SubMutableList(from, to uint32) MutableList {
	panic("implement me")
}

func (a *ArrayList) MutableIterator() MutableIterator {
	return &arrayListIterator{a, 0}
}

type arrayListIterator struct {
	arrayList *ArrayList
	index     uint32
}

func (a *arrayListIterator) HasNext() bool {
	return a.index < a.arrayList.used
}

func (a *arrayListIterator) Next() (interface{}, exceptions.Exception) {
	value, err := a.arrayList.Get(a.index)
	if err != nil {
		return nil, err
	}
	a.index++
	return value, nil
}

func (a *arrayListIterator) Remove() exceptions.Exception {
	err := a.arrayList.RemoveAt(a.index - 1)
	if err != nil {
		return err
	}
	a.index--
	return nil
}
