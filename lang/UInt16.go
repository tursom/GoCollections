/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "strconv"

type UInt16 uint16

type AsUInt16 interface {
	Object
	AsUInt16() UInt16
}

func CastUInt16(v any) (uint16, bool) {
	switch i := v.(type) {
	case uint16:
		return i, true
	case AsUInt16:
		return i.AsUInt16().V(), true
	default:
		return 0, false
	}
}

func EqualsUInt16(i1 AsUInt16, i2 any) bool {
	i2, ok := CastUInt16(i2)
	return ok && i2 == i1.AsUInt16().V()
}

func (i UInt16) V() uint16 {
	return uint16(i)
}

func (i *UInt16) P() *uint16 {
	return (*uint16)(i)
}

func (i UInt16) AsUInt16() UInt16 {
	return i
}

func (i UInt16) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i UInt16) AsObject() Object {
	return i
}

func (i UInt16) Equals(e Object) bool {
	return EqualsUInt16(i, e)
}

func (i UInt16) ToString() String {
	return NewString(i.String())
}

func (i UInt16) HashCode() int32 {
	return int32(i)
}

func (i UInt16) Compare(t UInt16) int {
	switch {
	case i > t:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
