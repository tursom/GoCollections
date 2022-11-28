/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import (
	"unsafe"
)

//goland:noinspection GoVetUnsafePointer
func HashAddr[V any](v *V) int32 {
	p := uintptr(unsafe.Pointer(v))
	remain := int32(unsafe.Sizeof(*v))
	hash := int32(0)
	for remain-4 >= 0 {
		remain -= 4
		hash = hash*31 ^ *(*int32)(unsafe.Pointer(p))
		p += 4
	}
	for remain > 0 {
		hash = hash*31 ^ int32(*(*int8)(unsafe.Pointer(p)))
		p++
		remain--
	}

	return hash
}

func Hash32[T any](p *T) int32 {
	return HashInt32(*(*int32)(unsafe.Pointer(p)))
}

func Hash64[T any](p *T) int32 {
	i := *(*int64)(unsafe.Pointer(p))
	return HashInt64(i)
}

func HashString(str string) int32 {
	hashCode := int32(0)
	for _, r := range str {
		hashCode = 31*hashCode + r
	}
	return hashCode
}
