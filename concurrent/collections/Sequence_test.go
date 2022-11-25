/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

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
		for i := range sequence.RawChannel() {
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
	sequence.Close()
}
