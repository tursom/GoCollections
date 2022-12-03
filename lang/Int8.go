/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type Int8 int8

func CastInt8(v any) (int8, bool) {
	switch i := v.(type) {
	case int:
		return int8(i), true
	case int8:
		return i, true
	case int16:
		return int8(i), true
	case int32:
		return int8(i), true
	case int64:
		return int8(i), true
	case uint:
		return int8(i), true
	case uint8:
		return int8(i), true
	case uint16:
		return int8(i), true
	case uint32:
		return int8(i), true
	case uint64:
		return int8(i), true
	case float32:
		return int8(i), true
	case float64:
		return int8(i), true
	case Number:
		return int8(i.ToFloat64().V()), true
	default:
		return 0, false
	}
}

func EqualsInt8(i1 Number, i2 any) bool {
	i2, ok := CastInt8(i2)
	return ok && i2 == i1.ToInt64().V()
}

func (i Int8) V() int8 {
	return int8(i)
}

func (i *Int8) P() *int8 {
	return (*int8)(i)
}

func (i Int8) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int8) AsObject() Object {
	return i
}

func (i Int8) Equals(e Object) bool {
	return EqualsInt8(i, e)
}

func (i Int8) ToString() String {
	return NewString(i.String())
}

func (i Int8) HashCode() int32 {
	return HashInt8(int8(i))
}

func (i Int8) Compare(t Int8) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}

func (i Int8) ToInt64() Int64 {
	return Int64(i)
}

func (i Int8) ToUInt64() UInt64 {
	return UInt64(i)
}

func (i Int8) ToFloat64() Float64 {
	return Float64(i)
}

func (i *Int8) BitLength() int {
	return 8
}

func (i *Int8) SetBit(bit int, up bool) (old bool) {
	//TODO implement me
	panic("implement me")
}

func HashInt8(i int8) int32 {
	return HashInt32(int32(i))
}
