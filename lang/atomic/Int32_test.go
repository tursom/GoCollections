package atomic

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/tursom/GoCollections/lang"
)

func TestInt32_SetBit(t *testing.T) {
	var i Int32
	for j := 0; j < 32; j++ {
		if i.SetBit(j, true) {
			t.Fail()
		}
		if i.Load() != 1<<j {
			t.Fail()
		}
		fmt.Println(i.Load())
		if !i.SetBit(j, false) {
			t.Fail()
		}
		if i.Load() != 0 {
			t.Fail()
		}
	}
}

func TestInt32_CompareAndSwapBit(t *testing.T) {
	var i Int32
	cond := sync.WaitGroup{}
	cond.Add(1)
	for k := 0; k < 128; k++ {
		id := k
		go func() {
			cond.Wait()
			for j := 0; j < 32; j++ {
				for !i.CompareAndSwapBit(j, false, true) {
				}
				for !i.CompareAndSwapBit(j, true, false) {
				}
			}
			fmt.Printf("%d finished\n", id)
		}()
	}
	cond.Done()
	time.Sleep(1 * time.Second)
	fmt.Println(i.Load())
}

func TestTypeConv(t *testing.T) {
	var intV lang.Int32
	var i any = intV
	_, ok := i.(lang.AsInt32)
	fmt.Println(reflect.TypeOf(i), ok)

	i = &intV
	_, ok = i.(lang.AsInt32)
	fmt.Println(reflect.TypeOf(i), ok)
}
