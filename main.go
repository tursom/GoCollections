package main

import (
	"fmt"
	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/exceptions"
)

func main() {
	_, err := exceptions.Try(func() (interface{}, error) {
		panic("test")
	}, func(r interface{}) (interface{}, error) {
		fmt.Println("recover from panic", r)
		return nil, exceptions.NewIndexOutOfBound(fmt.Sprint(r), true)
	})
	exceptions.NewRuntimeException(
		err,
		"test exception:",
		true,
		err,
	).PrintStackTrace()

	list := collections.NewArrayList()
	fmt.Println(list)
	for i := 0; i < 20; i++ {
		list.Add(i)
		//fmt.Println(list)
	}

	_ = collections.LoopMutable(list, func(element interface{}, iterator collections.MutableIterator) (err error) {
		if element.(int)&1 == 0 {
			err = iterator.Remove()
		}
		fmt.Println(list)
		return
	})
	//for i := 0; i < 10; i++ {
	//	list.Remove(i * 2)
	//	fmt.Println(list)
	//}
}
