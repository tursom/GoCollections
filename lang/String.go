package lang

type String struct {
	string
	hashCode int32
}

func NewString(str string) String {
	return String{string: str}
}

func (i String) AsString() string {
	return i.string
}

func (i String) GoString() string {
	return i.string
}

func (i String) String() string {
	return i.string
}

func (i String) AsObject() Object {
	return i
}

func (i String) Equals(e Object) bool {
	i2, ok := e.(String)
	if !ok {
		return false
	}
	return i == i2
}

func (i String) ToString() String {
	return i
}

func (i String) HashCode() int32 {
	if i.hashCode != 0 {
		return i.hashCode
	}
	i.hashCode = HashString(i.string)
	return i.hashCode
}

func (i String) Compare(t String) int {
	switch {
	case i.string > t.string:
		return 1
	case i == t:
		return 0
	default:
		return -1
	}
}
