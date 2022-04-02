package lang

import (
	"fmt"
	"testing"
)

func TestThreadLocalImpl(t1 *testing.T) {
	local := NewThreadLocal[int]()
	fmt.Println(local.Get())
	local.Put(1)
	fmt.Println(local.Get())
	local.Remove()
	fmt.Println(local.Get())
}
