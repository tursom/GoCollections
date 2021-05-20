package collections

import (
	"collections/exceptions"
	"fmt"
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
		Remove(element interface{}) bool
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

		Get(index uint32) (interface{}, error)
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
		Remove(element interface{}) bool
		AddAll(c Collection) bool
		RemoveAll(c Collection) bool
		RetainAll(c Collection) bool
		Clear()

		Get(index uint32) (interface{}, error)
		SubList(from, to uint32) List

		Set(index uint32, element interface{}) bool
		AddAtIndex(index uint32, element interface{}) bool
		RemoveAt(index uint32) bool
		SubMutableList(from, to uint32) MutableList
	}
)

func Contains(l Collection, element interface{}) bool {
	return Loop(l, func(e interface{}) error {
		if e == element {
			return exceptions.ElementFound
		}
		return nil
	}) != nil
}

func ContainsAll(l Collection, collection Collection) bool {
	return Loop(collection, func(e interface{}) error {
		if l.Contains(e) {
			return nil
		} else {
			return exceptions.ElementNotFound
		}
	}) == nil
}

func AddAll(l MutableCollection, collection Collection) bool {
	return Loop(collection, func(e interface{}) error {
		if !l.Add(e) {
			return exceptions.CollectionLoopFinished
		}
		return nil
	}) == nil
}

func RemoveAll(l MutableCollection, collection Collection) bool {
	return Loop(collection, func(e interface{}) error {
		if !l.Remove(e) {
			return exceptions.CollectionLoopFinished
		}
		return nil
	}) == nil
}

func RetainAll(l MutableCollection, collection Collection) bool {
	_ = LoopMutable(l, func(element interface{}, iterator MutableIterator) error {
		if !collection.Contains(element) {
			iterator.Remove()
		}
		return nil
	})
	return true
}

func String(l List) string {
	if l.IsEmpty() {
		return "[]"
	}

	builder := strings.Builder{}
	builder.WriteString("[")
	iterator := l.Iterator()
	next := iterator.Next()
	builder.WriteString(fmt.Sprint(next))
	for iterator.HasNext() {
		builder.WriteString(", ")
		next = iterator.Next()
		builder.WriteString(fmt.Sprint(next))
	}
	builder.WriteString("]")
	return builder.String()
}
