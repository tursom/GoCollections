package lang

type Object interface {
	Equals(e Object) bool
}

func Equals(e Object, t Object) bool {
	if e == nil {
		return t == nil
	}
	return e.Equals(t)
}
