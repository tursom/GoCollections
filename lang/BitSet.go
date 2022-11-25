/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

type (
	BitSet interface {
		BitLength() int
		SetBit(bit int, up bool) (old bool)
	}
)
