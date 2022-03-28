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
