package collections

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSequence_Alloc(t *testing.T) {
	var sequence Sequence[int]
	go func() {
		for i := range sequence.Channel() {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 10000; i++ {
		id := i
		sender := sequence.Alloc()
		go func() {
			time.Sleep(time.Duration(rand.Float32() * float32(time.Second) * 5 * (float32(id) / 10000)))
			sender.Send(id)
		}()
	}
	time.Sleep(time.Second * 10)
}
