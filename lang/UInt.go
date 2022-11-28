/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type UInt uint

func CastUInt(v any) (uint, bool) {
	switch i := v.(type) {
	case int:
		return uint(i), true
	case int8:
		return uint(i), true
	case int16:
		return uint(i), true
	case int32:
		return uint(i), true
	case int64:
		return uint(i), true
	case uint:
		return i, true
	case uint8:
		return uint(i), true
	case uint16:
		return uint(i), true
	case uint32:
		return uint(i), true
	case uint64:
		return uint(i), true
	case float32:
		return uint(i), true
	case float64:
		return uint(i), true
	case Number:
		return uint(i.ToUInt64().V()), true
	default:
		return 0, false
	}
}

func EqualsUInt(i1 Number, i2 any) bool {
	i2, ok := CastUInt(i2)
	return ok && i2 == i1.ToUInt64().V()
}

func (i UInt) V() uint {
	return uint(i)
}

func (i *UInt) P() *uint {
	return (*uint)(i)
}

func (i UInt) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt) AsObject() Object {
	return i
}

func (i UInt) Equals(e Object) bool {
	return EqualsUInt(i, e)
}

func (i UInt) ToString() String {
	return NewString(i.String())
}

func (i UInt) HashCode() int32 {
	return HashUInt(uint(i))
}

func (i UInt) Compare(t UInt) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i UInt) ToInt64() Int64 {
	return Int64(i)
}

func (i UInt) ToUInt64() UInt64 {
	return UInt64(i)
}

func (i UInt) ToFloat64() Float64 {
	return Float64(i)
}

func HashUInt(i uint) int32 {
	return HashInt(int(i))
}
