package collections

import (
	"fmt"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"strings"
)

type (
	Collection[T any] interface {
		Iterable[T]
		Size() int
		IsEmpty() bool
		Contains(element T) bool
		ContainsAll(c Collection[T]) bool
	}

	MutableCollection[T any] interface {
		Collection[T]
		MutableIterable[T]

		Add(element T) bool
		Remove(element T) exceptions.Exception
		AddAll(c Collection[T]) bool
		RemoveAll(c Collection[T]) bool
		RetainAll(c Collection[T]) bool
		Clear()
	}

	List[T any] interface {
		Collection[T]

		Get(index int) (T, exceptions.Exception)
		SubList(from, to int) List[T]
	}

	MutableList[T any] interface {
		List[T]
		MutableCollection[T]

		Set(index int, element T) exceptions.Exception
		AddAtIndex(index int, element T) bool
		RemoveAt(index int) exceptions.Exception
		SubMutableList(from, to int) MutableList[T]
	}
)

func ListGet[T lang.Object](list List[T], index int) T {
	get, err := list.Get(index)
	if err != nil {
		panic(err)
	}
	return get
}

func ContainsAll[T lang.Object](l Collection[T], collection Collection[T]) bool {
	return Loop[T](collection, func(e T) exceptions.Exception {
		if l.Contains(e) {
			return nil
		} else {
			return exceptions.ElementNotFound
		}
	}) == nil
}

func AddAll[T any](l MutableCollection[T], collection Collection[T]) bool {
	return Loop[T](collection, func(e T) exceptions.Exception {
		if !l.Add(e) {
			return exceptions.CollectionLoopFinished
		}
		return nil
	}) == nil
}

func String[T any](l Iterable[T]) string {
	iterator := l.Iterator()
	if !iterator.HasNext() {
		return "[]"
	}

	builder := strings.Builder{}
	builder.WriteString("[")
	next, _ := iterator.Next()
	builder.WriteString(fmt.Sprint(next))
	for iterator.HasNext() {
		builder.WriteString(", ")
		next, _ = iterator.Next()
		builder.WriteString(fmt.Sprint(next))
	}
	builder.WriteString("]")
	return builder.String()
}
