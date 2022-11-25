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
	CopyOnWriteList[E lang.Object] struct {
		lang.BaseObject
		arr []E
	}

	copyOnWriteListIterator[E lang.Object] struct {
		lang.BaseObject
		list  *CopyOnWriteList[E]
		arr   []E
		index int
	}
)

func (l *CopyOnWriteList[E]) Iterator() Iterator[E] {
	return l.MutableIterator()
}

func (l *CopyOnWriteList[E]) Size() int {
	return len(l.arr)
}

func (l *CopyOnWriteList[E]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *CopyOnWriteList[E]) Contains(element E) bool {
	return Contains[E](l, element)
}

func (l *CopyOnWriteList[E]) ContainsAll(c Collection[E]) bool {
	return ContainsAll[E](l, c)
}

func (l *CopyOnWriteList[E]) MutableIterator() MutableIterator[E] {
	return &copyOnWriteListIterator[E]{
		list:  l,
		arr:   l.arr,
		index: 0,
	}
}

func (l *CopyOnWriteList[E]) Add(element E) bool {
	l.arr = append(l.arr, element)
	return true
}

func (l *CopyOnWriteList[E]) Remove(element E) exceptions.Exception {
	return Remove[E](l, element)
}

func (l *CopyOnWriteList[E]) AddAll(c Collection[E]) bool {
	_ = Loop[E](c, func(element E) exceptions.Exception {
		l.arr = append(l.arr, element)
		return nil
	})
	return true
}

func (l *CopyOnWriteList[E]) RemoveAll(c Collection[E]) bool {
	return RemoveAll[E](l, c)
}

func (l *CopyOnWriteList[E]) RetainAll(c Collection[E]) bool {
	return RetainAll[E](l, c)
}

func (l *CopyOnWriteList[E]) Clear() {
	l.arr = make([]E, 0)
}

func (l *CopyOnWriteList[E]) Get(index int) (E, exceptions.Exception) {
	if index < 0 || index >= l.Size() {
		return lang.Nil[E](), exceptions.NewIndexOutOfBound("", nil)
	}
	return l.arr[index], nil
}

func (l *CopyOnWriteList[E]) SubList(from, to int) List[E] {
	return l.SubList(from, to)
}

func (l *CopyOnWriteList[E]) Set(index int, element E) exceptions.Exception {
	if index < 0 || index >= l.Size() {
		return exceptions.NewIndexOutOfBound("", nil)
	}
	l.arr[index] = element
	return nil
}

func (l *CopyOnWriteList[E]) AddAtIndex(index int, element E) bool {
	//TODO implement me
	panic("implement me")
}

func (l *CopyOnWriteList[E]) RemoveAt(index int) exceptions.Exception {
	//TODO implement me
	panic("implement me")
}

func (l *CopyOnWriteList[E]) SubMutableList(from, to int) MutableList[E] {
	//TODO implement me
	panic("implement me")
}

func (c *copyOnWriteListIterator[E]) HasNext() bool {
	//TODO implement me
	panic("implement me")
}

func (c *copyOnWriteListIterator[E]) Next() (E, exceptions.Exception) {
	//TODO implement me
	panic("implement me")
}

func (c *copyOnWriteListIterator[E]) Remove() exceptions.Exception {
	//TODO implement me
	panic("implement me")
}
