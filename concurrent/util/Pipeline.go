/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package util

import (
	"github.com/tursom/GoCollections/concurrent/collections"
	"github.com/tursom/GoCollections/lang"
)

type (
	pipeline[T, R any] struct {
		value  T
		sender collections.SequenceSender[R]
	}
)

func NewPipeline[T, R any](producer <-chan T, concurrency int, consumer func(T) R) lang.ReceiveChannel[R] {
	var sequence collections.Sequence[R]
	c := lang.NewChannel[pipeline[T, R]](0)
	go func() {
		defer close(c)
		for value := range producer {
			c.Send(pipeline[T, R]{
				value:  value,
				sender: sequence.Alloc(),
			})
		}
	}()

	for i := 0; i < concurrency; i++ {
		go func() {
			defer sequence.Close()
			for value := range c {
				r := consumer(value.value)
				value.sender.Send(r)
			}
		}()
	}
	return sequence.Channel()
}
