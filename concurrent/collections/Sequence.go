package collections

import (
	"sync"

	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang/atomic"
)

type (
	// Sequence 一个用于并发生产场景下使消息有序的 Sequence
	// 数据的接收顺序与 Sender 的生成顺序保持一致
	Sequence[T any] struct {
		lock sync.Mutex
		ch   chan T
		head *sequenceNode[T]
		end  **sequenceNode[T]
	}

	// SequenceSender 用于给 Sequence 发送信息
	SequenceSender[T any] interface {
		Send(value T)
		Fail(cause exceptions.Exception)
	}

	sequenceNode[T any] struct {
		sent     bool
		value    T
		cause    exceptions.Exception
		sequence *Sequence[T]
		next     *sequenceNode[T]
	}
)

// channel 懒加载 Sequence channel
func (s *Sequence[T]) channel() chan T {
	// 经典懒加载单例写法
	if s.ch == nil {
		s.lock.Lock()
		defer s.lock.Unlock()
		if s.ch == nil {
			s.ch = make(chan T, 16)
		}
	}

	return s.ch
}

// Channel 获取用于读取 Sequence 数据的 channel
func (s *Sequence[T]) Channel() <-chan T {
	return s.channel()
}

func (s *Sequence[T]) Send(msg T) {
	s.Alloc().Send(msg)
}

// Alloc 预分配数据空间，数据将按照从 Sequence 分配时的顺序有序发送
func (s *Sequence[T]) Alloc() SequenceSender[T] {
	s.lock.Lock()
	defer s.lock.Unlock()

	node := &sequenceNode[T]{
		sequence: s,
	}
	if s.end == nil {
		s.end = &s.head
	}
	*s.end = node
	s.end = &node.next
	return node
}

// send 清空 Sequence 中可发送的消息
func (s *Sequence[T]) send() {
	channel := s.channel()
	s.lock.Lock()
	defer s.lock.Unlock()

	head := s.head
	for head != nil && head.sent {
		if head.cause == nil {
			channel <- head.value
		} else {
			close(channel)
			s.ch = nil
			panic(head.cause)
		}
		head = head.next
		// 重置 s.head 防止并发安全问题产生
		// 防止 sequenceNode.Send 判断自身是否是 head 产生的问题
		atomic.StorePointer(&s.head, head)
	}
}

func (s *sequenceNode[T]) Send(value T) {
	s.value = value
	s.sent = true
	if atomic.LoadPointer(&s.sequence.head) == s {
		s.sequence.send()
	}
}

func (s *sequenceNode[T]) Fail(cause exceptions.Exception) {
	s.cause = cause
	s.sent = true
	if atomic.LoadPointer(&s.sequence.head) == s {
		s.sequence.send()
	}
}
