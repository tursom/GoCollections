/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package atomic

type Bool Int32

func (v *Bool) Load() (val bool) {
	return (*Int32)(v).Load() != 0
}

func (v *Bool) Store(val bool) {
	if val {
		(*Int32)(v).Store(1)
	} else {
		(*Int32)(v).Store(0)
	}
}

func (v *Bool) Swap(new bool) (old bool) {
	n := int32(0)
	if new {
		n = 1
	}
	return (*Int32)(v).Swap(n) != 0
}

func (v *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	o := int32(0)
	if old {
		o = 1
	}
	n := int32(0)
	if new {
		n = 1
	}
	return (*Int32)(v).CompareAndSwap(o, n)
}
