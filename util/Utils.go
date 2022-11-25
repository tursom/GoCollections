/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package util

type (
	Stringer struct {
		stringer func() string
	}
)

func NewStringer(stringer func() string) Stringer {
	return Stringer{
		stringer: stringer,
	}
}

func (s Stringer) String() string {
	if s.stringer == nil {
		return "nil"
	}
	return s.stringer()
}
