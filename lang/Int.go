/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type Int int

func CastInt(v any) (int, bool) {
	switch i := v.(type) {
	case int:
		return i, true
	case int8:
		return int(i), true
	case int16:
		return int(i), true
	case int32:
		return int(i), true
	case int64:
		return int(i), true
	case uint:
		return int(i), true
	case uint8:
		return int(i), true
	case uint16:
		return int(i), true
	case uint32:
		return int(i), true
	case uint64:
		return int(i), true
	case float32:
		return int(i), true
	case float64:
		return int(i), true
	case Number:
		return int(i.ToInt64().V()), true
	default:
		return 0, false
	}
}

func EqualsInt(i1 Number, i2 any) bool {
	i2, ok := CastInt(i2)
	return ok && i2 == i1.ToInt64().V()
}

func (i Int) V() int {
	return int(i)
}

func (i *Int) P() *int {
	return (*int)(i)
}

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int) AsObject() Object {
	return i
}

func (i Int) Equals(e Object) bool {
	return EqualsInt(i, e)
}

func (i Int) ToString() String {
	return NewString(i.String())
}

func (i Int) HashCode() int32 {
	return HashInt(int(i))
}

func (i Int) Compare(t Int) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i Int) ToInt64() Int64 {
	return Int64(i)
}

func (i Int) ToUInt64() UInt64 {
	return UInt64(i)
}

func (i Int) ToFloat64() Float64 {
	return Float64(i)
}

func HashInt(i int) int32 {
	return HashInt64(int64(i))
}
