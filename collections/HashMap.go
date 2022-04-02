package collections

import (
	"fmt"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	//HashMap
	HashMap[K lang.Object, V any] struct {
		NodeMap[K, V]
		slot       []*hashMapNode[K, V]
		loadFactor float32
		size       int
	}

	abstractHashMapNode[K lang.Object, V any] struct {
		m    *HashMap[K, V]
		hash int32
	}

	hashMapNode[K lang.Object, V any] struct {
		lang.BaseObject
		abstractHashMapNode[K, V]
		key   K
		value V
		next  *hashMapNode[K, V]
	}

	emptyHashMapSlot[K lang.Object, V any] struct {
		lang.BaseObject
		abstractHashMapNode[K, V]
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

func (m *HashMap[K, V]) String() string {
	return MapToString[K, V](m).String()
}

func (m *HashMap[K, V]) findSlot(k K) MapNode[K, V] {
	hashCode := lang.HashCode(k)
	hashCode ^= hashCode >> 16
	index := int(hashCode) % len(m.slot)
	root := m.slot[index]
	if root == nil {
		return &emptyHashMapSlot[K, V]{abstractHashMapNode: abstractHashMapNode[K, V]{m: m, hash: hashCode}}
	}
	return root
}

func (m *HashMap[K, V]) Loop(f func(K, V) exceptions.Exception) (err exceptions.Exception) {
	for _, node := range m.slot {
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

func (m *HashMap[K, V]) resize() {
	slot := m.slot
	resize := len(slot)
	load := int(float32(m.size) / m.loadFactor)
	for load < resize {
		resize >>= 1
	}
	for load > resize {
		resize <<= 1
	}
	if resize == len(slot) {
		return
	}
	m.slot = make([]*hashMapNode[K, V], resize)
	for _, node := range slot {
		for node != nil {
			next := node.next
			index := node.index()
			node.next = m.slot[index]
			m.slot[index] = node
			node = next
		}
	}
}

func (m *HashMap[K, V]) resizeFromRemove() {
	if len(m.slot) > int(float32(m.size)/m.loadFactor) {
		m.resize()
	}
}

func (m *HashMap[K, V]) resizeFromAdd() {
	if len(m.slot) < int(float32(m.size)/m.loadFactor) {
		m.resize()
	}
}

func (n *abstractHashMapNode[K, V]) index() int {
	return int(n.hash) % len(n.m.slot)
}

func (e *emptyHashMapSlot[K, V]) GetKey() K {
	return lang.Nil[K]()
}

func (e *emptyHashMapSlot[K, V]) GetValue() V {
	return lang.Nil[V]()
}

func (e *emptyHashMapSlot[K, V]) SetValue(_ V) {
}

func (e *emptyHashMapSlot[K, V]) CreateNext(key K) MapNode[K, V] {
	index := e.index()
	node := &hashMapNode[K, V]{
		key:                 key,
		abstractHashMapNode: abstractHashMapNode[K, V]{m: e.m, hash: lang.HashCode(key)},
		next:                e.m.slot[index],
	}
	e.m.slot[index] = node
	e.m.size++
	e.m.resizeFromAdd()
	return node
}

func (e *emptyHashMapSlot[K, V]) GetNext() MapNode[K, V] {
	node := e.m.slot[e.index()]
	// required
	if node == nil {
		return nil
	}
	return node
}

func (e *emptyHashMapSlot[K, V]) RemoveNext() {
	node := e.m.slot[e.index()]
	if node != nil {
		e.m.slot[e.index()] = node.next
		e.m.size--
		e.m.resizeFromRemove()
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
	s.next = &hashMapNode[K, V]{
		key:  key,
		next: s.next,
		abstractHashMapNode: abstractHashMapNode[K, V]{
			m:    s.m,
			hash: lang.HashCode(key),
		},
	}
	s.m.size++
	s.m.resizeFromAdd()
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
		s.m.size--
		s.m.resizeFromRemove()
	}
}
