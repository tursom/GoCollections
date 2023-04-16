/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

type (
	BitSet interface {
		BitLength() uint
		SetBit(index uint, up bool) (old bool)
		GetBit(index uint) (ok bool)
	}
)
