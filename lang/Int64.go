/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type Int64 int64

func CastInt64(v any) (int64, bool) {
	switch i := v.(type) {
	case int:
		return int64(i), true
	case int8:
		return int64(i), true
	case int16:
		return int64(i), true
	case int32:
		return int64(i), true
	case int64:
		return i, true
	case uint:
		return int64(i), true
	case uint8:
		return int64(i), true
	case uint16:
		return int64(i), true
	case uint32:
		return int64(i), true
	case uint64:
		return int64(i), true
	case float32:
		return int64(i), true
	case float64:
		return int64(i), true
	case Number:
		return i.ToInt64().V(), true
	default:
		return 0, false
	}
}

func EqualsInt64(i1 Number, i2 any) bool {
	i2, ok := CastInt64(i2)
	return ok && i2 == i1.ToUInt64().V()
}

func (i Int64) V() int64 {
	return int64(i)
}

func (i *Int64) P() *int64 {
	return (*int64)(i)
}

func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int64) AsObject() Object {
	return i
}

func (i Int64) Equals(e Object) bool {
	return EqualsInt64(i, e)
}

func (i Int64) ToString() String {
	return NewString(i.String())
}

func (i Int64) HashCode() int32 {
	return HashInt64(int64(i))
}

func (i Int64) Compare(t Int64) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i Int64) ToInt64() Int64 {
	return Int64(i)
}

func (i Int64) ToUInt64() UInt64 {
	return UInt64(i)
}

func (i Int64) ToFloat64() Float64 {
	return Float64(i)
}

func HashInt64(i int64) int32 {
	return int32(i) ^ int32(i>>32)
}
