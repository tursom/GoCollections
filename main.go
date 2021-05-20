package main

import (
	"collections/collections"
	"fmt"
)

func main() {
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
