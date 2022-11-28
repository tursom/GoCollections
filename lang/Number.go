/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

// Number interface for real number
type Number interface {
	ToInt64() Int64
	ToUInt64() UInt64
	ToFloat64() Float64
}
