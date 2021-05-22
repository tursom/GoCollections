package collections

import (
	"fmt"
	"github.com/tursom/GoCollections/exceptions"
	"strings"
)

type (
	Collection interface {
		Iterator() Iterator
		Size() uint32
		IsEmpty() bool
		Contains(element interface{}) bool
		ContainsAll(c Collection) bool
	}

	MutableCollection interface {
		Iterator() Iterator
		Size() uint32
		IsEmpty() bool
		Contains(element interface{}) bool
		ContainsAll(c Collection) bool

		MutableIterator() MutableIterator
		Add(element interface{}) bool
		Remove(element interface{}) exceptions.Exception
		AddAll(c Collection) bool
		RemoveAll(c Collection) bool
		RetainAll(c Collection) bool
		Clear()
	}

	List interface {
		Iterator() Iterator
		Size() uint32
		IsEmpty() bool
		Contains(element interface{}) bool
		ContainsAll(c Collection) bool

		Get(index uint32) (interface{}, exceptions.Exception)
		SubList(from, to uint32) List
	}

	MutableList interface {
		Iterator() Iterator
		Size() uint32
		IsEmpty() bool
		Contains(element interface{}) bool
		ContainsAll(c Collection) bool

		MutableIterator() MutableIterator
		Add(element interface{}) bool
		Remove(element interface{}) exceptions.Exception
		AddAll(c Collection) bool
		RemoveAll(c Collection) bool
		RetainAll(c Collection) bool
		Clear()

		Get(index uint32) (interface{}, exceptions.Exception)
		SubList(from, to uint32) List

		Set(index uint32, element interface{}) exceptions.Exception
		AddAtIndex(index uint32, element interface{}) bool
		RemoveAt(index uint32) exceptions.Exception
		SubMutableList(from, to uint32) MutableList
	}
)

func ContainsAll(l Collection, collection Collection) bool {
	return Loop(collection, func(e interface{}) exceptions.Exception {
		if l.Contains(e) {
			return nil
		} else {
			return exceptions.ElementNotFound
		}
	}) == nil
}

func AddAll(l MutableCollection, collection Collection) bool {
	return Loop(collection, func(e interface{}) exceptions.Exception {
		if !l.Add(e) {
			return exceptions.CollectionLoopFinished
		}
		return nil
	}) == nil
}

func String(l Iterable) string {
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
