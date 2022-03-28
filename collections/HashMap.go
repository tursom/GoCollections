package collections

import (
	"fmt"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	//HashMap
	//TODO impl
	HashMap[K lang.Object, V any] struct {
		NodeMap[K, V]
		slot       []*hashMapNode[K, V]
		loadFactor float32
	}

	hashMapNode[K lang.Object, V any] struct {
		lang.BaseObject
		m     *HashMap[K, V]
		key   K
		value V
		next  *hashMapNode[K, V]
		hash  int32
	}

	emptyHashMapSlot[K lang.Object, V any] struct {
		lang.BaseObject
		m        *HashMap[K, V]
		hashCode int32
		index    int
	}
)

func NewHashMap[K lang.Object, V any]() *HashMap[K, V] {
	return NewHashMapInitCap[K, V](16, 0.75)
}

func NewHashMapInitCap[K lang.Object, V any](initialCapacity int, loadFactor float32) *HashMap[K, V] {
	m := &HashMap[K, V]{
		slot:       make([]*hashMapNode[K, V], initialCapacity),
		loadFactor: loadFactor,
	}
	m.MapNodeFinder = NewMapNodeFinderBySlot[K, V](m)
	return m
}

func (h *HashMap[K, V]) String() string {
	return MapToString[K, V](h).String()
}

func (h *HashMap[K, V]) findSlot(k K) MapNode[K, V] {
	hashCode := lang.HashCode(k)
	hashCode ^= hashCode >> 16
	index := int(hashCode) % len(h.slot)
	root := h.slot[index]
	if root == nil {
		return &emptyHashMapSlot[K, V]{m: h, hashCode: hashCode, index: index}
	}
	return root
}

func (h *HashMap[K, V]) Loop(f func(K, V) exceptions.Exception) (err exceptions.Exception) {
	for _, node := range h.slot {
		for node != nil {
			err = f(node.GetKey(), node.GetValue())
			if err != nil {
				return
			}
			node = node.next
		}
	}
	return
}

func (e *emptyHashMapSlot[K, V]) GetKey() K {
	//TODO implement me
	panic("implement me")
}

func (e *emptyHashMapSlot[K, V]) GetValue() V {
	//TODO implement me
	panic("implement me")
}

func (e *emptyHashMapSlot[K, V]) SetValue(value V) {
	//TODO implement me
	panic("implement me")
}

func (e *emptyHashMapSlot[K, V]) CreateNext(key K) MapNode[K, V] {
	node := &hashMapNode[K, V]{
		key:  key,
		hash: lang.HashCode(key),
	}
	e.m.slot[e.index] = node
	return node
}

func (e *emptyHashMapSlot[K, V]) GetNext() MapNode[K, V] {
	node := e.m.slot[e.index]
	if node == nil {
		return nil
	}
	return node
}

func (e *emptyHashMapSlot[K, V]) RemoveNext() {
	node := e.m.slot[e.index]
	if node != nil {
		e.m.slot[e.index] = node.next
	}
}

func (s *hashMapNode[K, V]) String() string {
	return "hashMapNode{key: " + s.key.String() + ", value: " + fmt.Sprint(s.value) + "}"
}

func (s *hashMapNode[K, V]) GetKey() K {
	return s.key
}

func (s *hashMapNode[K, V]) GetValue() V {
	return s.value
}

func (s *hashMapNode[K, V]) SetValue(value V) {
	s.value = value
}

func (s *hashMapNode[K, V]) CreateNext(key K) MapNode[K, V] {
	s.next = &hashMapNode[K, V]{key: key, next: s.next, hash: lang.HashCode(key)}
	return s.next
}

func (s *hashMapNode[K, V]) GetNext() MapNode[K, V] {
	if s.next == nil {
		return nil
	}
	return s.next
}

func (s *hashMapNode[K, V]) RemoveNext() {
	if s.next != nil {
		s.next = s.next.next
	}
}
