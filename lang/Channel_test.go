/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"fmt"
	"testing"

	"github.com/tursom/GoCollections/util/time"
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
