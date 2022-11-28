/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type UInt64 uint64

func CastUInt64(v any) (uint64, bool) {
	switch i := v.(type) {
	case int:
		return uint64(i), true
	case int8:
		return uint64(i), true
	case int16:
		return uint64(i), true
	case int32:
		return uint64(i), true
	case int64:
		return uint64(i), true
	case uint:
		return uint64(i), true
	case uint8:
		return uint64(i), true
	case uint16:
		return uint64(i), true
	case uint32:
		return uint64(i), true
	case uint64:
		return uint64(i), true
	case float32:
		return uint64(i), true
	case float64:
		return uint64(i), true
	case Number:
		return i.ToUInt64().V(), true
	default:
		return 0, false
	}
}

func EqualsUInt64(i1 Number, i2 any) bool {
	i2, ok := CastUInt64(i2)
	return ok && i2 == i1.ToUInt64().V()
}

func (i UInt64) V() uint64 {
	return uint64(i)
}

func (i *UInt64) P() *uint64 {
	return (*uint64)(i)
}

func (i UInt64) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt64) AsObject() Object {
	return i
}

func (i UInt64) Equals(e Object) bool {
	return EqualsUInt64(i, e)
}

func (i UInt64) ToString() String {
	return NewString(i.String())
}

func (i UInt64) HashCode() int32 {
	return HashUInt64(uint64(i))
}

func (i UInt64) Compare(t UInt64) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i UInt64) ToInt64() Int64 {
	return Int64(i)
}

func (i UInt64) ToUInt64() UInt64 {
	return i
}

func (i UInt64) ToFloat64() Float64 {
	return Float64(i)
}

func HashUInt64(i uint64) int32 {
	return HashInt64(int64(i))
}
