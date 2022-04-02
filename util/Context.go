package util

import (
	"github.com/tursom/GoCollections/lang"
	"github.com/tursom/GoCollections/lang/atomic"
	"sync"
)

var (
	contextId atomic.Int32
)

type (
	Context struct {
		lang.BaseObject
		contextId int32
		keyId     atomic.Int32
	}
	ContextMap interface {
		Get(key *ContextKey[any]) (any, bool)
		Set(key *ContextKey[any], value any)
	}
	ContextKey[T any] struct {
		lang.BaseObject
		contextId int32
		id        int32
	}
	contextMapImpl struct {
		lang.BaseObject
		contextId int32
		array     []any
	}
	concurrentContextMap struct {
		contextMapImpl
		lock sync.Mutex
	}
)

func NewContext() *Context {
	return &Context{
		contextId: contextId.Add(1),
	}
}

func AllocateContextKey[T any](ctx *Context) *ContextKey[T] {
	return &ContextKey[T]{
		contextId: ctx.contextId,
		id:        ctx.keyId.Add(1) - 1,
	}
}

func (c *Context) NewMap() ContextMap {
	return &contextMapImpl{
		contextId: c.contextId,
		//array:     make([]any, c.keyId.Load()),
	}
}

func (c *Context) NewConcurrentMap() ContextMap {
	return &concurrentContextMap{
		contextMapImpl: contextMapImpl{
			contextId: c.contextId,
			//array:     make([]any, c.keyId.Load()),
		},
	}
}

func (k *ContextKey[T]) Get(m ContextMap) T {
	value, _ := k.TryGet(m)
	return value
}

func (k *ContextKey[T]) TryGet(m ContextMap) (T, bool) {
	value, ok := m.Get(k.asNormalKey())
	return lang.Cast[T](value), ok
}

func (k *ContextKey[T]) Set(m ContextMap, value T) {
	m.Set(k.asNormalKey(), value)
}

func (m *contextMapImpl) Get(key *ContextKey[any]) (any, bool) {
	if len(m.array) < int(key.id) {
		return nil, false
	} else {
		return m.array[key.id], true
	}
}

func (m *contextMapImpl) Set(key *ContextKey[any], value any) {
	if len(m.array) < int(key.id) {
		newArray := make([]any, key.id)
		copy(newArray, m.array)
		m.array = newArray
	}
	m.array[key.id] = value
}

func (m *concurrentContextMap) Set(key *ContextKey[any], value any) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.contextMapImpl.Set(key, value)
}

func (k *ContextKey[T]) asNormalKey() *ContextKey[any] {
	return (*ContextKey[any])(k)
}
