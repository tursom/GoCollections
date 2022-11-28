/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type UInt8 uint8

func CastUInt8(v any) (uint8, bool) {
	switch i := v.(type) {
	case int:
		return uint8(i), true
	case int8:
		return uint8(i), true
	case int16:
		return uint8(i), true
	case int32:
		return uint8(i), true
	case int64:
		return uint8(i), true
	case uint:
		return uint8(i), true
	case uint8:
		return i, true
	case uint16:
		return uint8(i), true
	case uint32:
		return uint8(i), true
	case uint64:
		return uint8(i), true
	case float32:
		return uint8(i), true
	case float64:
		return uint8(i), true
	case Number:
		return uint8(i.ToUInt64().V()), true
	default:
		return 0, false
	}
}

func EqualsUInt8(i1 Number, i2 any) bool {
	i2, ok := CastUInt8(i2)
	return ok && i2 == i1.ToUInt64().V()
}

func (i UInt8) V() uint8 {
	return uint8(i)
}

func (i *UInt8) P() *uint8 {
	return (*uint8)(i)
}

func (i UInt8) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt8) AsObject() Object {
	return i
}

func (i UInt8) Equals(e Object) bool {
	return EqualsUInt8(i, e)
}

func (i UInt8) ToString() String {
	return NewString(i.String())
}

func (i UInt8) HashCode() int32 {
	return HashUInt8(uint8(i))
}

func (i UInt8) Compare(t UInt8) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i UInt8) ToInt64() Int64 {
	return Int64(i)
}

func (i UInt8) ToUInt64() UInt64 {
	return UInt64(i)
}

func (i UInt8) ToFloat64() Float64 {
	return Float64(i)
}

func HashUInt8(i uint8) int32 {
	return HashUInt32(uint32(i))
}
