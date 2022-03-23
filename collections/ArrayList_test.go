package collections

import (
	"fmt"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"testing"
)

func TestArrayListAdd(t *testing.T) {
	list := NewArrayList[lang.Int]()
	for i := 0; i < 10; i++ {
		list.Add(lang.Int(i))
		fmt.Println(list)
	}
	for i := 0; i < 10; i++ {
		list.RemoveLast()
		fmt.Println(list)
	}
}

func TestLinkedList(t *testing.T) {
	list := NewLinkedList[lang.Int]()
	for i := 0; i < 10; i++ {
		list.Add(lang.Int(i))
		fmt.Println(list)
		for j := 0; j <= i; j++ {
			//exceptions.Exec2r0[
			//	func(func() (error, error, error)) (error, error),
			//	func() (error, error, error),
			//	error,
			//](
			exceptions.Exec2r0(
				//exceptions.Exec1r1,
				//exceptions.Exec0r2,
				exceptions.Exec1r1[func() (error, error, error), error, error],
				exceptions.Exec0r2[error, error, error],
				func() (error, error, error) {
					return nil, nil, nil
				},
			)
			fmt.Println(exceptions.Exec2r1((*LinkedList[lang.Int]).Get, list, j).AsInt())
		}
	}
	for i := 0; i < 10; i++ {
		list.RemoveLast()
		fmt.Println(list)
	}
}

func TestEqualsFunc(t *testing.T) {
	f := func() {}
	fmt.Println(equals(f, f))
}

func equals(v1, v2 any) bool {
	return v1 == v2
}
