package atomic

type Bool struct {
	i Int32
}

func (v *Bool) Load() (val bool) {
	return v.i.Load() != 0
}

func (v *Bool) Store(val bool) {
	if val {
		v.i.Store(1)
	} else {
		v.i.Store(0)
	}
}

func (v *Bool) Swap(new bool) (old bool) {
	n := int32(0)
	if new {
		n = 1
	}
	return v.i.Swap(n) != 0
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
	return v.i.CompareAndSwap(o, n)
}
