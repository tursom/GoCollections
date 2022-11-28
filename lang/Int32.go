/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type Int32 int32

type Rune = Int32

func CastInt32(v any) (int32, bool) {
	switch i := v.(type) {
	case int:
		return int32(i), true
	case int8:
		return int32(i), true
	case int16:
		return int32(i), true
	case int32:
		return i, true
	case int64:
		return int32(i), true
	case uint:
		return int32(i), true
	case uint8:
		return int32(i), true
	case uint16:
		return int32(i), true
	case uint32:
		return int32(i), true
	case uint64:
		return int32(i), true
	case float32:
		return int32(i), true
	case float64:
		return int32(i), true
	case Number:
		return int32(i.ToInt64().V()), true
	default:
		return 0, false
	}
}

func EqualsInt32(i1 Number, i2 any) bool {
	i2, ok := CastInt32(i2)
	return ok && i2 == i1.ToInt64().V()
}

func (i Int32) V() int32 {
	return int32(i)
}

func (i *Int32) P() *int32 {
	return (*int32)(i)
}

func (i Int32) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int32) AsObject() Object {
	return i
}

func (i Int32) Equals(e Object) bool {
	return EqualsInt32(i, e)
}

func (i Int32) ToString() String {
	return NewString(i.String())
}

func (i Int32) HashCode() int32 {
	return HashInt32(int32(i))
}

func (i Int32) Compare(t Int32) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i Int32) ToInt64() Int64 {
	return Int64(i)
}

func (i Int32) ToUInt64() UInt64 {
	return UInt64(i)
}

func (i Int32) ToFloat64() Float64 {
	return Float64(i)
}

func HashInt32(i int32) int32 {
	return i
}
