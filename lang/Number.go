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

func GetBit[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 |
	Int8 | Int16 | Int32 | Int64 | UInt8 | UInt16 | UInt32 | UInt64](p T, index uint) (ok bool) {
	location := T(1) << index
	return p&location != 0
}

func SwapBit[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 |
	Int8 | Int16 | Int32 | Int64 | UInt8 | UInt16 | UInt32 | UInt64](p *T, index uint, new bool) (old bool) {
	location := T(1) << index
	oldValue := *p
	var newValue T
	if new {
		newValue = oldValue | location
	} else {
		newValue = oldValue & ^location
	}
	*p = newValue
	return oldValue&location != 0
}
