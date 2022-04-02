package collections

import (
	"fmt"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"strings"
)

type (
	MapLooper[K lang.Object, V any] interface {
		Loop(func(K, V) exceptions.Exception) exceptions.Exception
	}

	Map[K lang.Object, V any] interface {
		MapLooper[K, V]
		Get(k K) (V, bool, exceptions.Exception)
	}

	MutableMap[K lang.Object, V any] interface {
		Map[K, V]
		Put(k K, v V) (bool, exceptions.Exception)
		Remove(k K) (V, bool, exceptions.Exception)
	}

	MapNodeFinder[K lang.Object, V any] interface {
		findNode(k K, createIfNotExist bool) (node, prev MapNode[K, V])
	}

	MapSlotFinder[K lang.Object, V any] interface {
		findSlot(k K) (root MapNode[K, V])
	}

	MapNode[K lang.Object, V any] interface {
		GetKey() K
		GetValue() V
		SetValue(value V)
		CreateNext(key K) MapNode[K, V]
		GetNext() MapNode[K, V]
		RemoveNext()
	}

	SimpleMapNode[K lang.Object, V any] struct {
		lang.BaseObject
		key   K
		value V
		next  *SimpleMapNode[K, V]
	}

	NodeMap[K lang.Object, V any] struct {
		lang.BaseObject
		MapNodeFinder MapNodeFinder[K, V]
	}

	MapNodeFinderBySlot[K lang.Object, V any] struct {
		lang.BaseObject
		slotFinder MapSlotFinder[K, V]
	}
)

func NewMapNodeFinderBySlot[K lang.Object, V any](slotFinder MapSlotFinder[K, V]) MapNodeFinder[K, V] {
	return &MapNodeFinderBySlot[K, V]{slotFinder: slotFinder}
}

func MapToString[K lang.Object, V any](m MapLooper[K, V]) lang.String {
	builder := strings.Builder{}
	builder.WriteString("{")
	_ = m.Loop(func(k K, v V) exceptions.Exception {
		if builder.Len() != 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(k.String())
		builder.WriteString(": ")
		if ov, ok := lang.TryCast[lang.Object](v); ok {
			builder.WriteString(ov.String())
		} else {
			builder.WriteString(fmt.Sprint(v))
		}
		return nil
	})
	builder.WriteString("}")
	return lang.NewString(builder.String())
}

func (g *NodeMap[K, V]) findNode(k K, createIfNotExist bool) (node, prev MapNode[K, V]) {
	return g.MapNodeFinder.findNode(k, createIfNotExist)
}

func (g *NodeMap[K, V]) Put(k K, v V) (bool, exceptions.Exception) {
	node, _ := g.findNode(k, true)
	node.SetValue(v)
	return true, nil
}

func (g *NodeMap[K, V]) Get(k K) (V, bool, exceptions.Exception) {
	node, _ := g.findNode(k, false)
	if node == nil {
		return g.nil()
	} else {
		return node.GetValue(), true, nil
	}
}

func (g *NodeMap[K, V]) Remove(k K) (V, bool, exceptions.Exception) {
	node, prev := g.findNode(k, false)
	if node == nil {
		return g.nil()
	} else {
		if prev != nil {
			prev.RemoveNext()
		}
		return node.GetValue(), true, nil
	}
}

func (g *NodeMap[K, V]) nil() (V, bool, exceptions.Exception) {
	return lang.Nil[V](), false, nil
}

func (m *MapNodeFinderBySlot[K, V]) findNode(k K, createIfNotExist bool) (node, prev MapNode[K, V]) {
	prev = m.slotFinder.findSlot(k)
	if prev != nil {
		node = prev.GetNext()
		for node != nil {
			if node.GetKey().Equals(k) {
				return
			}
			prev = node
			node = node.GetNext()
		}
	}
	if createIfNotExist {
		node = prev.CreateNext(k)
	}
	return
}

func (s *SimpleMapNode[K, V]) String() string {
	return "SimpleMapNode{key: " + s.key.String() + ", value: " + fmt.Sprint(s.value) + "}"
}

func (s *SimpleMapNode[K, V]) GetKey() K {
	return s.key
}

func (s *SimpleMapNode[K, V]) GetValue() V {
	return s.value
}

func (s *SimpleMapNode[K, V]) SetValue(value V) {
	s.value = value
}

func (s *SimpleMapNode[K, V]) CreateNext(key K) MapNode[K, V] {
	s.next = &SimpleMapNode[K, V]{key: key, next: s.next}
	return s.next
}

func (s *SimpleMapNode[K, V]) GetNext() MapNode[K, V] {
	if s.next == nil {
		return nil
	}
	return s.next
}

func (s *SimpleMapNode[K, V]) RemoveNext() {
	if s.next != nil {
		s.next = s.next.next
	}
}
