package atomic

import (
	"fmt"
	"testing"
)

func TestAtomic_Store(t *testing.T) {
	a := NewReference[int](nil)
	var i = 1
	a.Store(&i)
	i = 2
	fmt.Println(*a.Load())
}

func TestReferenceOf(t *testing.T) {
	one := 1

	var p *int = nil
	ref := ReferenceOf(&p)

	ref.Store(&one)
	fmt.Println(ref.Load())
	fmt.Println(*ref.Load())

	_ = *ref.AsUintptr() + 1
}
