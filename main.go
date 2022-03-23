package main

import (
	"fmt"
	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"time"
)

func main() {
	_, err := exceptions.Try(func() (any, exceptions.Exception) {
		panic("test")
	}, func(r any) (any, exceptions.Exception) {
		fmt.Println("recover from panic", r)
		return nil, exceptions.NewIndexOutOfBound(fmt.Sprint(r), nil)
	})
	exceptions.Print(err)

	list := collections.NewConcurrentLinkedQueue[lang.Int]()
	target := collections.NewArrayListByCapacity[lang.Int](10000)
	fmt.Println("list", list)

	go func() {
		for i := 0; i < 1000000; i++ {
			list.Offer()
			//fmt.Println(offer)
			//if element != nil {
			//	target.Add(element)
			//}
		}
		fmt.Println("target:", target)
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			err = list.Push(lang.Int(i))
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
	//_ = collections.LoopMutable(list, func(element any, iterator collections.MutableIterator) (err exceptions.Exception) {
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
