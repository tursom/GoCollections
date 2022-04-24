package collections

import (
	"sync"

	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/concurrent"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	// GoSyncMap an map use go sync.Map to store hash slot
	GoSyncMap[K lang.Object, V any] struct {
		collections.NodeMap[K, V]
		m    sync.Map
		lock concurrent.RWLock
	}
)

func NewGoSyncMap[K lang.Object, V any]() *GoSyncMap[K, V] {
	m := &GoSyncMap[K, V]{lock: &sync.RWMutex{}}
	m.MapNodeFinder = collections.NewMapNodeFinderBySlot[K, V](m)
	return m
}

func (g *GoSyncMap[K, V]) ToString() lang.String {
	return collections.MapToString[K, V](g)
}

func (g *GoSyncMap[K, V]) FindSlot(k K) collections.MapNode[K, V] {
	hashCode := lang.HashCode(k)
	p, _ := g.m.Load(hashCode)
	root := lang.Cast[*collections.SimpleMapNode[K, V]](p)
	if root == nil {
		root = &collections.SimpleMapNode[K, V]{}
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
