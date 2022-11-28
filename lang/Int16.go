/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type Int16 int16

func CastInt16(v any) (int16, bool) {
	switch i := v.(type) {
	case int:
		return int16(i), true
	case int8:
		return int16(i), true
	case int16:
		return i, true
	case int32:
		return int16(i), true
	case int64:
		return int16(i), true
	case uint:
		return int16(i), true
	case uint8:
		return int16(i), true
	case uint16:
		return int16(i), true
	case uint32:
		return int16(i), true
	case uint64:
		return int16(i), true
	case float32:
		return int16(i), true
	case float64:
		return int16(i), true
	case Number:
		return int16(i.ToInt64().V()), true
	default:
		return 0, false
	}
}

func EqualsInt16(i1 Number, i2 any) bool {
	i2, ok := CastInt16(i2)
	return ok && i2 == i1.ToInt64().V()
}

func (i Int16) V() int16 {
	return int16(i)
}

func (i *Int16) P() *int16 {
	return (*int16)(i)
}

func (i Int16) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int16) AsObject() Object {
	return i
}

func (i Int16) Equals(e Object) bool {
	return EqualsInt16(i, e)
}

func (i Int16) ToString() String {
	return NewString(i.String())
}

func (i Int16) HashCode() int32 {
	return HashInt16(int16(i))
}

func (i Int16) Compare(t Int16) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i Int16) ToInt64() Int64 {
	return Int64(i)
}

func (i Int16) ToUInt64() UInt64 {
	return UInt64(i)
}

func (i Int16) ToFloat64() Float64 {
	return Float64(i)
}

func HashInt16(i int16) int32 {
	return HashInt32(int32(i))
}
