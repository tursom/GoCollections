/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "time"

type (
	convChannel[T1, T2 any] struct {
		channel  Channel[T1]
		recvConv func(T1) T2
		sendConv func(T2) T1
	}

	convSendChannel[T1, T2 any] struct {
		channel  SendChannel[T1]
		sendConv func(T2) T1
	}

	convReceiveChannel[T1, T2 any] struct {
		channel  ReceiveChannel[T1]
		recvConv func(T1) T2
	}

	filterReceiveChannel[T any] struct {
		ReceiveChannel[T]
		filter func(T) bool
	}
)

func ConvChannel[T1, T2 any](c1 Channel[T1], recvConv func(T1) T2, sendConv func(T2) T1) ChannelProxy[T2, T1] {
	return &convChannel[T1, T2]{c1, recvConv, sendConv}
}

func ConvSendChannel[T1, T2 any](c1 SendChannel[T1], sendConv func(T2) T1) SendChannelProxy[T2, T1] {
	return &convSendChannel[T1, T2]{c1, sendConv}
}

func ConvReceiveChannel[T1, T2 any](c1 ReceiveChannel[T1], recvConv func(T1) T2) ReceiveChannelProxy[T2, T1] {
	return &convReceiveChannel[T1, T2]{c1, recvConv}
}

func FilterReceiveChannel[T any](c1 ReceiveChannel[T], filter func(T) bool) ReceiveChannelProxy[T, T] {
	return &filterReceiveChannel[T]{c1, filter}
}

func (c *convChannel[T1, T2]) Close() {
	c.channel.Close()
}

func (c *convChannel[T1, T2]) SCh() chan<- T2 {
	return nil
}

func (c *convChannel[T1, T2]) Send(obj T2) {
	c.channel.Send(c.sendConv(obj))
}

func (c *convChannel[T1, T2]) TrySend(obj T2) bool {
	return c.channel.TrySend(c.sendConv(obj))
}

func (c *convChannel[T1, T2]) SendTimeout(obj T2, timeout time.Duration) bool {
	return c.channel.SendTimeout(c.sendConv(obj), timeout)
}

func (c *convChannel[T1, T2]) RCh() <-chan T2 {
	return nil
}

func (c *convChannel[T1, T2]) Receive() T2 {
	return c.recvConv(c.channel.Receive())
}

func (c *convChannel[T1, T2]) TryReceive() (T2, bool) {
	obj, b := c.channel.TryReceive()
	if !b {
		return Nil[T2](), false
	}

	return c.recvConv(obj), true
}

func (c *convChannel[T1, T2]) ReceiveTimeout(timeout time.Duration) (T2, bool) {
	obj, b := c.channel.ReceiveTimeout(timeout)
	if !b {
		return Nil[T2](), false
	}

	return c.recvConv(obj), true
}

func (c *convChannel[T1, T2]) Ch() chan T2 {
	return nil
}

func (c *convChannel[T1, T2]) ProxySendChannel() SendChannel[T1] {
	return c.channel
}

func (c *convChannel[T1, T2]) ProxyReceiveChannel() ReceiveChannel[T1] {
	return c.channel
}

func (c *convChannel[T1, T2]) ProxyChannel() Channel[T1] {
	return c.channel
}

func (c *convSendChannel[T1, T2]) Close() {
	c.channel.Close()
}

func (c *convSendChannel[T1, T2]) SCh() chan<- T2 {
	return nil
}

func (c *convSendChannel[T1, T2]) Send(obj T2) {
	c.channel.Send(c.sendConv(obj))
}

func (c *convSendChannel[T1, T2]) TrySend(obj T2) bool {
	return c.channel.TrySend(c.sendConv(obj))
}

func (c *convSendChannel[T1, T2]) SendTimeout(obj T2, timeout time.Duration) bool {
	return c.channel.SendTimeout(c.sendConv(obj), timeout)
}

func (c *convSendChannel[T1, T2]) ProxySendChannel() SendChannel[T1] {
	return c.channel
}

func (c *convReceiveChannel[T1, T2]) Close() {
	c.channel.Close()
}

func (c *convReceiveChannel[T1, T2]) RCh() <-chan T2 {
	return nil
}

func (c *convReceiveChannel[T1, T2]) Receive() T2 {
	return c.recvConv(c.channel.Receive())
}

func (c *convReceiveChannel[T1, T2]) TryReceive() (T2, bool) {
	obj, b := c.channel.TryReceive()
	if !b {
		return Nil[T2](), false
	}

	return c.recvConv(obj), true
}

func (c *convReceiveChannel[T1, T2]) ReceiveTimeout(timeout time.Duration) (T2, bool) {
	obj, b := c.channel.ReceiveTimeout(timeout)
	if !b {
		return Nil[T2](), false
	}

	return c.recvConv(obj), true
}

func (c *convReceiveChannel[T1, T2]) ProxyReceiveChannel() ReceiveChannel[T1] {
	return c.channel
}

func (c *filterReceiveChannel[T]) RCh() <-chan T {
	return nil
}

func (f *filterReceiveChannel[T]) Receive() T {
	obj := f.ReceiveChannel.Receive()
	for !f.filter(obj) {
		obj = f.ReceiveChannel.Receive()
	}

	return obj
}

func (f *filterReceiveChannel[T]) TryReceive() (T, bool) {
	obj, ok := f.ReceiveChannel.TryReceive()
	if !ok {
		return Nil[T](), false
	}

	for !f.filter(obj) {
		obj, ok = f.ReceiveChannel.TryReceive()
		if !ok {
			return Nil[T](), false
		}
	}

	return obj, true
}

func (f *filterReceiveChannel[T]) ReceiveTimeout(timeout time.Duration) (T, bool) {
	t1 := time.Now()

	obj, ok := f.ReceiveChannel.ReceiveTimeout(timeout)
	if !ok {
		return Nil[T](), false
	}

	for !f.filter(obj) {
		duration := timeout - t1.Sub(time.Now())
		if duration <= 0 {
			return Nil[T](), false
		}

		obj, ok = f.ReceiveChannel.ReceiveTimeout(duration)
		if !ok {
			return Nil[T](), false
		}
	}

	return obj, true
}

func (f *filterReceiveChannel[T]) ProxyReceiveChannel() ReceiveChannel[T] {
	return f.ReceiveChannel
}
