package collections

import "github.com/tursom/GoCollections/exceptions"

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
}

type Iterable interface {
	Iterator() Iterator
}

type MutableIterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove() error
}

type MutableIterable interface {
	Iterator() Iterator
	MutableIterator() MutableIterator
}

func Loop(iterable Iterable, f func(element interface{}) error) error {
	if f == nil {
		return exceptions.NewNPE("", true)
	}
	iterator := iterable.Iterator()
	for iterator.HasNext() {
		next, err := iterator.Next()
		if err != nil {
			return err
		}
		err = f(next)
		if err != nil {
			return err
		}
	}
	return nil
}

func LoopMutable(iterable MutableIterable, f func(element interface{}, iterator MutableIterator) (err error)) error {
	if f == nil {
		return nil
	}
	iterator := iterable.MutableIterator()
	for iterator.HasNext() {
		next, err := iterator.Next()
		if err != nil {
			return err
		}
		err = f(next, iterator)
		if err != nil {
			return err
		}
	}
	return nil
}
