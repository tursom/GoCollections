package util

import (
	"fmt"
	"testing"
)

func TestContextKey(t *testing.T) {
	ctx := NewContext()
	key := AllocateContextKey[int](ctx)

	m := ctx.NewConcurrentMap()

	fmt.Println(key.Get(m))
	fmt.Println(key.TryGet(m))

	key.Set(m, 100)
	fmt.Println(key.Get(m))
	fmt.Println(key.TryGet(m))
}
