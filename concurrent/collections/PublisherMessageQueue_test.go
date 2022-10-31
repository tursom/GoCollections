package collections

import (
	"fmt"
	"testing"
	"time"

	"github.com/tursom/GoCollections/util/unsafe"
)

func TestPublisherMessageQueueNode_Sizeof(t *testing.T) {
	fmt.Println(unsafe.Sizeof[publisherMessageQueueNode[int]]())
}

func TestMessageQueue_Subscribe(t *testing.T) {
	var mq PublisherMessageQueue[string]
	for i := 0; i < 3; i++ {
		id := i
		subscribe := mq.Subscribe()
		go func() {
			fmt.Println("run subscriber", id)
			defer func() {
				subscribe.Close()
				fmt.Println("subscriber", id, "closed")
			}()

			for true {
				msg := subscribe.Receive()
				fmt.Println(id, msg)
			}
		}()
	}
	time.Sleep(time.Second / 10)

	for i := 0; i < 3; i++ {
		index := i
		go func() {
			for j := 0; j < 16; j++ {
				mq.Send(fmt.Sprintf("%d-%d", index, j))
			}
		}()
	}
	time.Sleep(time.Second * 10)
}
