package collections

import "github.com/tursom/GoCollections/exceptions"
import "github.com/tursom/GoCollections/lang"

type Iterator[T any] interface {
	HasNext() bool
	Next() (T, exceptions.Exception)
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type MutableIterator[T any] interface {
	Iterator[T]
	Remove() exceptions.Exception
}

type MutableIterable[T any] interface {
	Iterable[T]
	MutableIterator() MutableIterator[T]
}

func Loop[T any](iterable Iterable[T], f func(element T) exceptions.Exception) exceptions.Exception {
	if f == nil || iterable == nil {
		return exceptions.NewNPE("", nil)
	}
	return LoopIterator(iterable.Iterator(), f)
}

func LoopMutable[T any](
	iterable MutableIterable[T],
	f func(element T, iterator MutableIterator[T]) (err exceptions.Exception),
) exceptions.Exception {
	if f == nil || iterable == nil {
		return exceptions.NewNPE("", nil)
	}
	return LoopMutableIterator(iterable.MutableIterator(), f)
}

func LoopIterator[T any](iterator Iterator[T], f func(element T) exceptions.Exception) exceptions.Exception {
	if f == nil || iterator == nil {
		return exceptions.NewNPE("", nil)
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

func LoopMutableIterator[T any](
	iterator MutableIterator[T],
	f func(element T, iterator MutableIterator[T]) (err exceptions.Exception),
) exceptions.Exception {
	if f == nil || iterator == nil {
		return exceptions.NewNPE("", nil)
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

func Size[T any](iterable Iterable[T]) (size int, err exceptions.Exception) {
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

func Contains[T lang.Object](l Iterable[T], element T) bool {
	return Loop(l, func(e T) exceptions.Exception {
		if lang.Equals(element, e) {
			return exceptions.ElementFound
		}
		return nil
	}) != nil
}

func Remove[T lang.Object](l MutableIterable[T], element T) exceptions.Exception {
	return LoopMutable(l, func(e T, iterator MutableIterator[T]) (err exceptions.Exception) {
		if lang.Equals(element, e) {
			return iterator.Remove()
		}
		return nil
	})
}

func RemoveAll[T any](l MutableIterable[T], collection Collection[T]) bool {
	return LoopMutable(l, func(element T, iterator MutableIterator[T]) (err exceptions.Exception) {
		if collection.Contains(element) {
			return iterator.Remove()
		}
		return nil
	}) == nil
}

func RetainAll[T any](l MutableIterable[T], collection Collection[T]) bool {
	return LoopMutable(l, func(element T, iterator MutableIterator[T]) exceptions.Exception {
		if !collection.Contains(element) {
			return iterator.Remove()
		}
		return nil
	}) == nil
}

func Clear[T any](l MutableIterable[T]) exceptions.Exception {
	return LoopMutable(l, func(element T, iterator MutableIterator[T]) (err exceptions.Exception) {
		return iterator.Remove()
	})
}
