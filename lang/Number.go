/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

// Number interface for real number
type Number interface {
	ToInt64() Int64
	ToUInt64() UInt64
	ToFloat64() Float64
}

func SwapBit[T int32 | int64 | uint32 | uint64](p *T, bit int, new bool) (old bool) {
	location := T(1) << bit
	oldValue := *p
	var newValue T
	if new {
		newValue = oldValue | location
	} else {
		newValue = oldValue & ^location
	}
	*p = newValue
	return oldValue&newValue != 0
}
