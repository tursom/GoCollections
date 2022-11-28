/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type UInt32 uint32

func CastUInt32(v any) (uint32, bool) {
	switch i := v.(type) {
	case int:
		return uint32(i), true
	case int8:
		return uint32(i), true
	case int16:
		return uint32(i), true
	case int32:
		return uint32(i), true
	case int64:
		return uint32(i), true
	case uint:
		return uint32(i), true
	case uint8:
		return uint32(i), true
	case uint16:
		return uint32(i), true
	case uint32:
		return i, true
	case uint64:
		return uint32(i), true
	case float32:
		return uint32(i), true
	case float64:
		return uint32(i), true
	case Number:
		return uint32(i.ToUInt64().V()), true
	default:
		return 0, false
	}
}

func EqualsUInt32(i1 Number, i2 any) bool {
	i2, ok := CastUInt32(i2)
	return ok && i2 == i1.ToUInt64().V()
}

func (i UInt32) V() uint32 {
	return uint32(i)
}

func (i *UInt32) P() *uint32 {
	return (*uint32)(i)
}

func (i UInt32) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt32) AsObject() Object {
	return i
}

func (i UInt32) Equals(e Object) bool {
	return EqualsUInt32(i, e)
}

func (i UInt32) ToString() String {
	return NewString(i.String())
}

func (i UInt32) HashCode() int32 {
	return HashUInt32(uint32(i))
}

func (i UInt32) Compare(t UInt32) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i UInt32) ToInt64() Int64 {
	return Int64(i)
}

func (i UInt32) ToUInt64() UInt64 {
	return UInt64(i)
}

func (i UInt32) ToFloat64() Float64 {
	return Float64(i)
}

func HashUInt32(i uint32) int32 {
	return HashInt32(int32(i))
}
