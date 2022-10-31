package lang

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel_Send(t *testing.T) {
	ch := NewChannel[int](0)

	fmt.Println(ch.TrySend(0))

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	ch.Send(1)
	ch.SCh() <- 2
	time.Sleep(time.Second / 10)
	fmt.Println(ch.TrySend(3))

	ch.Close()
}
