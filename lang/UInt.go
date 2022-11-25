/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type UInt uint

type AsUInt interface {
	Object
	AsUInt() UInt
}

func CastUInt(v any) (uint, bool) {
	switch i := v.(type) {
	case uint:
		return i, true
	case AsUInt:
		return i.AsUInt().V(), true
	default:
		return 0, false
	}
}

func EqualsUInt(i1 AsUInt, i2 any) bool {
	i2, ok := CastUInt(i2)
	return ok && i2 == i1.AsUInt().V()
}

func (i UInt) V() uint {
	return uint(i)
}

func (i *UInt) P() *uint {
	return (*uint)(i)
}

func (i UInt) AsUInt() UInt {
	return i
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
	return int32(i)
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
