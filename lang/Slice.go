/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

type Slice[T Object] []T

func NewSlice[T Object](size, cap int) Slice[T] {
	return make(Slice[T], size, cap)
}

func (s Slice[T]) Contains(value T) bool {
	for _, e := range s {
		if e.Equals(value) {
			return true
		}
	}

	return false
}

func (s *Slice[T]) Append(value T) {
	*s = append(*s, value)
}

func (s Slice[T]) Size() int {
	return len(s)
}
