package collections

import (
	"github.com/tursom/GoCollections/concurrent"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"sync"
)

type (
	GoSyncMap[K lang.Object, V any] struct {
		NodeMap[K, V]
		m    sync.Map
		lock concurrent.RWLock
	}
)

func NewGoSyncMap[K lang.Object, V any]() *GoSyncMap[K, V] {
	m := &GoSyncMap[K, V]{lock: &sync.RWMutex{}}
	m.MapNodeFinder = NewMapNodeFinderBySlot[K, V](m)
	return m
}

func (g *GoSyncMap[K, V]) ToString() lang.String {
	return MapToString[K, V](g)
}

func (g *GoSyncMap[K, V]) findSlot(k K) MapNode[K, V] {
	hashCode := lang.HashCode(k)
	p, _ := g.m.Load(hashCode)
	root := lang.Cast[*SimpleMapNode[K, V]](p)
	if root == nil {
		root = &SimpleMapNode[K, V]{}
		g.m.Store(hashCode, root)
	}
	return root
}

func (g *GoSyncMap[K, V]) Put(k K, v V) (bool, exceptions.Exception) {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.NodeMap.Put(k, v)
}

func (g *GoSyncMap[K, V]) Get(k K) (V, bool, exceptions.Exception) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.NodeMap.Get(k)
}

func (g *GoSyncMap[K, V]) Remove(k K) (V, bool, exceptions.Exception) {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.NodeMap.Remove(k)
}

func (g *GoSyncMap[K, V]) Loop(f func(K, V) exceptions.Exception) (err exceptions.Exception) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	g.m.Range(func(key, value any) bool {
		err = f(lang.Cast[K](key), lang.Cast[V](value))
		return err == nil
	})
	return
}
