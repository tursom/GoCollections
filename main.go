package main

import (
	"fmt"
	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/exceptions"
	"time"
)

func main() {
	_, err := exceptions.Try(func() (interface{}, exceptions.Exception) {
		panic("test")
	}, func(r interface{}) (interface{}, exceptions.Exception) {
		fmt.Println("recover from panic", r)
		return nil, exceptions.NewIndexOutOfBound(fmt.Sprint(r), true)
	})
	exceptions.Print(err)

	list := collections.NewConcurrentLinkedQueue()
	target := collections.NewArrayListByCapacity(10000)
	fmt.Println("list", list)

	go func() {
		for i := 0; i < 1000000; i++ {
			element, _ := list.Offer()
			//fmt.Println(offer)
			if element != nil {
				target.Add(element)
			}
		}
		fmt.Println("target:", target)
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			err = list.Push(i)
			//fmt.Println(err)
		}
		time.Sleep(time.Second * 2)
		fmt.Println(list)
	}()
	//
	//for i := 0; i < 100; i++ {
	//	fmt.Println(list)
	//}
	time.Sleep(time.Second * 10)
	fmt.Println("target:", target)

	//for i := 0; i < 20; i++ {
	//	list.Push(i)
	//	fmt.Println(list)
	//}
	//
	//for i := 0; i < 20; i++ {
	//	list.Offer()
	//	fmt.Println(list)
	//}
	//
	//for i := 0; i < 25; i++ {
	//	list.Push(i)
	//	fmt.Println(list)
	//}
	//
	//_ = collections.LoopMutable(list, func(element interface{}, iterator collections.MutableIterator) (err exceptions.Exception) {
	//	if element.(int)&1 == 0 {
	//		err = iterator.Remove()
	//	}
	//	fmt.Println(list)
	//	return
	//})
	//for i := 0; i < 10; i++ {
	//	list.Remove(i * 2)
	//	fmt.Println(list)
	//}
}
