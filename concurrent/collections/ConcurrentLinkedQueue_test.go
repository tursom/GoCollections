/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package collections

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
	"unsafe"

	"github.com/tursom/GoCollections/collections"
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

	for i := 0; i < 16; i++ {
		testConcurrentLinkedQueueIteratorRemove(t, queue, 1000)
		t.Logf("Test_concurrentLinkedQueueIterator_Remove passed on %d loop", i+1)
	}
}

func testConcurrentLinkedQueueIteratorRemove(t *testing.T, queue *ConcurrentLinkedQueue[lang.Int], nodeNumber int) {
	nodes := make(chan collections.StackNode[lang.Int], nodeNumber)
	nodesIndex := make([]bool, nodeNumber)
	b := rand.Int()&1 == 1
	if b {
		for i := 0; i < 4; i++ {
			go func() {
				for node := range nodes {
					index := exceptions.Exec0r1(node.RemoveAndGet)
					if nodesIndex[index] {
						t.Fatalf("TODO Fatalf")
					}
					nodesIndex[index] = true
				}
			}()
		}
	}

	for i := 0; i < nodeNumber; i++ {
		node, _ := queue.OfferAndGetNode(lang.Int(i))
		nodes <- node
	}

	close(nodes)

	if !b {
		for i := 0; i < 4; i++ {
			go func() {
				for node := range nodes {
					index := exceptions.Exec0r1(node.RemoveAndGet)
					if nodesIndex[index] {
						t.Fatalf("TODO Fatalf")
					}
					nodesIndex[index] = true
				}
			}()
		}
	}

	time.Sleep(time.Second / 5)

	head := queue.head
	if head != nil && head.deleted {
		head = head.next
	}
	if head != nil {
		t.Fatalf(fmt.Sprintf("queue remain %d element, is not thread safe", queue.Size()))
	}
}
