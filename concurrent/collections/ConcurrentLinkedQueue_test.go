package collections

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"

	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type data struct {
	lang.BaseObject
	id    int
	index int
}

func (d *data) String() string {
	return fmt.Sprintf("(%d,%d)", d.id, d.index)
}

func TestConcurrentLinkedQueue_NodeSize(t *testing.T) {
	fmt.Println(unsafe.Alignof(linkedStackNode[*data]{}))
	fmt.Println(unsafe.Sizeof(linkedStackNode[*data]{}))
}

func TestConcurrentLinkedQueue_Push(t *testing.T) {
	queue := NewLinkedQueue[*data]()
	cond := sync.WaitGroup{}
	cond.Add(1)
	for i := 0; i < 10; i++ {
		id := i
		go func() {
			cond.Wait()
			for i := 0; i < 100; i++ {
				_ = queue.Offer(&data{id: id, index: i})
			}
		}()
	}
	time.Sleep(time.Second)
	cond.Done()
	time.Sleep(time.Second)
	fmt.Println(queue)
}

func TestConcurrentLinkedQueue_ThreadSafe(t *testing.T) {
	times := 400000
	queue := NewLinkedQueue[*data]()
	for i := 0; i < 100; i++ {
		id := i
		go func() {
			//nodes := make([]QueueNode[*data], 0)
			for j := 0; j < times; j++ {
				_ = queue.Offer(&data{id: id, index: j})
				//node, _ := queue.OfferAndGetNode(&data{id: id, index: j})
				//nodes = append(nodes, node)
				//fmt.Println(queue)
			}
			//time.Sleep(time.Second * 1)
			fmt.Println(queue.Size())
			//for _, node := range nodes {
			//	exceptions.Exec0r0(node.Remove)
			//}
			for j := 0; j < times; j++ {
				offer, _ := queue.Poll()
				if offer == nil {
					panic("offer is nil")
				}
			}
		}()
	}

	time.Sleep(time.Second * 10)
	if queue.Size() != 0 {
		t.Fatalf(fmt.Sprintf("queue remain %d element, is not thread safe", queue.Size()))
	}
	//for !queue.IsEmpty() {
	//	fmt.Println(queue.Offer())
	//}
}

func Test_concurrentLinkedQueueIterator_Remove(t *testing.T) {
	queue := NewLinkedQueue[lang.Int]()
	nodes := make([]QueueNode[lang.Int], 0)
	for i := 0; i < 1000; i++ {
		node, _ := queue.OfferAndGetNode(lang.Int(i))
		nodes = append(nodes, node)
		//fmt.Println(queue)
	}
	for _, node := range nodes {
		fmt.Println(exceptions.Exec0r1(node.RemoveAndGet))
	}
	if queue.Size() != 0 {
		t.Fatalf(fmt.Sprintf("queue remain %d element, is not thread safe", queue.Size()))
	}
}
