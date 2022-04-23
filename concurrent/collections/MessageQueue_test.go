package collections

import (
	"fmt"
	"testing"
	"time"
)

func TestMessageQueue_Subscribe(t *testing.T) {
	var mq MessageQueue[int]
	for i := 0; i < 3; i++ {
		id := i
		go func() {
			fmt.Println("run subscriber", id)
			subscribe, canceler := mq.Subscribe()
			for i := 0; i < 10; i++ {
				msg := <-subscribe
				fmt.Println(id, msg)
			}
			canceler()
		}()
	}
	for i := 0; i < 16; i++ {
		mq.Send(i)
	}
	time.Sleep(time.Second * 10)
}
