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
	GoMap[K lang.Object, V any] struct {
		NodeMap[K, V]
		m map[int32]*SimpleMapNode[K, V]
	}
	//goMapNode[K lang.Object, V any] struct {
	//	key   K
	//	value V
	//	next  *goMapNode[K, V]
	//}
)

func NewGoMap[K lang.Object, V any]() *GoMap[K, V] {
	m := &GoMap[K, V]{m: make(map[int32]*SimpleMapNode[K, V])}
	m.NodeMap.MapNodeFinder = NewMapNodeFinderBySlot[K, V](m)
	return m
}

func (g *GoMap[K, V]) ToString() lang.String {
	return MapToString[K, V](g)
}

func (g *GoMap[K, V]) String() string {
	return g.ToString().GoString()
}

func (g *GoMap[K, V]) FindSlot(k K) MapNode[K, V] {
	hashCode := lang.HashCode(k)
	root := g.m[hashCode]
	if root == nil {
		root = &SimpleMapNode[K, V]{}
		g.m[hashCode] = root
	}
	return root
}

//func (g *GoMap[K, V]) Put(k K, v V) (bool, exceptions.Exception) {
//	node, _ := g.FindNode(k, true)
//	node.value = v
//	return true, nil
//}
//
//func (g *GoMap[K, V]) Get(k K) (V, bool, exceptions.Exception) {
//	node, _ := g.FindNode(k, false)
//	if node == nil {
//		return g.nil()
//	} else {
//		return node.value, true, nil
//	}
//}
//
//func (g *GoMap[K, V]) Remove(k K) (V, bool, exceptions.Exception) {
//	node, prev := g.FindNode(k, false)
//	if node == nil {
//		return g.nil()
//	} else {
//		prev.next = node.next
//		return node.value, true, nil
//	}
//}

func (g *GoMap[K, V]) Loop(f func(K, V) exceptions.Exception) (err exceptions.Exception) {
	for _, node := range g.m {
		for node.next != nil {
			node = node.next
			err = f(node.key, node.value)
			if err != nil {
				return
			}
		}
	}
	return
}

func (g *GoMap[K, V]) nil() (V, bool, exceptions.Exception) {
	return lang.Nil[V](), false, nil
}
