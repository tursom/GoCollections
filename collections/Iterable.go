package collections

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Iterable interface {
	Iterator() Iterator
}

type MutableIterator interface {
	HasNext() bool
	Next() interface{}
	Remove() bool
}

type MutableIterable interface {
	Iterator() Iterator
	MutableIterator() MutableIterator
}

func Loop(iterable Iterable, f func(element interface{}) error) error {
	if f == nil {
		return nil
	}
	iterator := iterable.Iterator()
	for iterator.HasNext() {
		err := f(iterator.Next())
		if err != nil {
			return err
		}
	}
	return nil
}

func LoopMutable(iterable MutableIterable, f func(element interface{}, iterator MutableIterator) error) error {
	if f == nil {
		return nil
	}
	iterator := iterable.MutableIterator()
	for iterator.HasNext() {
		err := f(iterator.Next(), iterator)
		if err != nil {
			return err
		}
	}
	return nil
}
