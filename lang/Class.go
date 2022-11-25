/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package lang

import "reflect"

type (
	Class struct {
		t       reflect.Type
		methods map[string]Method
		fields  map[string]Field
	}
	Method struct {
	}
	Field struct {
	}
)

func (c Class) GetType() reflect.Type {
	return c.t
}

func (c Class) GetName() String {
	return NewString(c.t.Name())
}

func GenerateClass(t reflect.Type) *Class {
	//TODO impl
	//t.Method(1).Func.Call()
	return nil
}
