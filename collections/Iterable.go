package collections

import "github.com/tursom/GoCollections/exceptions"

type Iterator interface {
	HasNext() bool
	Next() (interface{}, exceptions.Exception)
}

type Iterable interface {
	Iterator() Iterator
}

type MutableIterator interface {
	HasNext() bool
	Next() (interface{}, exceptions.Exception)
	Remove() exceptions.Exception
}

type MutableIterable interface {
	Iterator() Iterator
	MutableIterator() MutableIterator
}

func Loop(iterable Iterable, f func(element interface{}) exceptions.Exception) exceptions.Exception {
	if f == nil || iterable == nil {
		return exceptions.NewNPE("", true)
	}
	return LoopIterator(iterable.Iterator(), f)
}

func LoopMutable(iterable MutableIterable, f func(element interface{}, iterator MutableIterator) (err exceptions.Exception)) exceptions.Exception {
	if f == nil || iterable == nil {
		return exceptions.NewNPE("", true)
	}
	return LoopMutableIterator(iterable.MutableIterator(), f)
}

func LoopIterator(iterator Iterator, f func(element interface{}) exceptions.Exception) exceptions.Exception {
	if f == nil || iterator == nil {
		return exceptions.NewNPE("", true)
	}
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

func LoopMutableIterator(iterator MutableIterator, f func(element interface{}, iterator MutableIterator) (err exceptions.Exception)) exceptions.Exception {
	if f == nil || iterator == nil {
		return exceptions.NewNPE("", true)
	}
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

func Size(iterable Iterable) (size uint32, err exceptions.Exception) {
	iterator := iterable.Iterator()
	for iterator.HasNext() {
		_, err = iterator.Next()
		if err != nil {
			return
		}
		size++
	}
	return
}

func Contains(l Iterable, element interface{}) bool {
	return Loop(l, func(e interface{}) exceptions.Exception {
		if e == element {
			return exceptions.ElementFound
		}
		return nil
	}) != nil
}

func Remove(l MutableIterable, element interface{}) exceptions.Exception {
	return LoopMutable(l, func(e interface{}, iterator MutableIterator) (err exceptions.Exception) {
		if element == e {
			return iterator.Remove()
		}
		return nil
	})
}

func RemoveAll(l MutableIterable, collection Collection) bool {
	return LoopMutable(l, func(element interface{}, iterator MutableIterator) (err exceptions.Exception) {
		if collection.Contains(element) {
			return iterator.Remove()
		}
		return nil
	}) == nil
}

func RetainAll(l MutableIterable, collection Collection) bool {
	return LoopMutable(l, func(element interface{}, iterator MutableIterator) exceptions.Exception {
		if !collection.Contains(element) {
			return iterator.Remove()
		}
		return nil
	}) == nil
}

func Clear(l MutableIterable) exceptions.Exception {
	return LoopMutable(l, func(element interface{}, iterator MutableIterator) (err exceptions.Exception) {
		return iterator.Remove()
	})
}
