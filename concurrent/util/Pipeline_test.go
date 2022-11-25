/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

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
