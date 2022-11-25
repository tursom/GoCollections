/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package collections

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	// Iterator an iterator over a collection or another entity that can be represented as a sequence of elements.
	// Allows to sequentially access the elements
	Iterator[T any] interface {
		// HasNext return true if the iterator has more elements
		HasNext() bool
		// Next returns the next element in the iteration
		// return exceptions.IndexOutOfBound if there is no more element
		Next() (T, exceptions.Exception)
	}

	// MutableIterator an iterator over a mutable collection.
	// Provides the ability to remove elements while iterating.
	MutableIterator[T any] interface {
		Iterator[T]
		// Remove removes from the underlying collection the last element returned by this iterator.
		Remove() exceptions.Exception
	}

	// ListIterator an iterator over a collection that supports indexed access
	ListIterator[T any] interface {
		Iterator[T]
		HasPrevious() bool
		Previous() (T, exceptions.Exception)
		NextIndex() int
		PreviousIndex() int
	}

	// MutableListIterator an iterator over a mutable collection that supports indexed access
	// Provides the ability to add, modify and remove elements while iterating.
	MutableListIterator[T any] interface {
		ListIterator[T]
		MutableIterator[T]
		Set(value T) exceptions.Exception
		Add(value T) exceptions.Exception
	}

	// Iterable classes that inherit from this interface can be represented as a sequence of elements that can be iterated over.
	// param T
	Iterable[T any] interface {
		Iterator() Iterator[T]
	}

	MutableIterable[T any] interface {
		Iterable[T]
		MutableIterator() MutableIterator[T]
	}
)

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

// SkipIterator skip [skip] items of [iterator]
// returns [iterator] itself
func SkipIterator[T any](iterator Iterator[T], skip int) (Iterator[T], exceptions.Exception) {
	for i := 0; i < skip; i++ {
		if iterator.HasNext() {
			_, err := iterator.Next()
			if err != nil {
				return nil, err
			}
		} else {
			break
		}
	}
	return iterator, nil
}
