package util

import (
	"fmt"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	c := make(chan int)
	r := NewPipeline(c, 4, func(t int) int {
		return t * t * t * t
	})

	go func() {
		defer close(c)
		for i := 0; i < 1024; i++ {
			c <- i
		}
		time.Sleep(time.Second)
	}()

	for i := range r.RCh() {
		fmt.Println(i)
	}
}
