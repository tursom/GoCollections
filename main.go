package main

import (
	"collections/collections"
	"collections/exceptions"
	"fmt"
)

func main() {
	fmt.Println(exceptions.Try(func() (interface{}, error) {
		panic("test")
	}, func(r interface{}) interface{} {
		fmt.Println("recover from panic", r)
		return exceptions.NewIndexOutOfBound("", true)
	}))

	list := collections.NewLinkedList()
	fmt.Println(list)
	for i := 0; i < 20; i++ {
		list.Add(i)
		//fmt.Println(list)
	}

	_ = collections.LoopMutable(list, func(element interface{}, iterator collections.MutableIterator) error {
		if element.(int)&1 == 0 {
			iterator.Remove()
		}
		fmt.Println(list)
		return nil
	})
	//for i := 0; i < 10; i++ {
	//	list.Remove(i * 2)
	//	fmt.Println(list)
	//}
}
