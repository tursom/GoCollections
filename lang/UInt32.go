/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type UInt32 uint32

type AsUInt32 interface {
	Object
	AsUInt32() UInt32
}

func CastUInt32(v any) (uint32, bool) {
	switch i := v.(type) {
	case uint32:
		return i, true
	case AsUInt32:
		return i.AsUInt32().V(), true
	default:
		return 0, false
	}
}

func EqualsUInt32(i1 AsUInt32, i2 any) bool {
	i2, ok := CastUInt32(i2)
	return ok && i2 == i1.AsUInt32().V()
}

func (i UInt32) V() uint32 {
	return uint32(i)
}

func (i *UInt32) P() *uint32 {
	return (*uint32)(i)
}

func (i UInt32) AsUInt32() UInt32 {
	return i
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
	return int32(i)
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
