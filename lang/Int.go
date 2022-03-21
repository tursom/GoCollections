package lang

type Int int

func (i Int) Equals(e Object) bool {
	i2, ok := e.(Int)
	if !ok {
		return false
	}
	return i == i2
}

func (i Int) Compare(t Int) int {
	return int(t - i)
}
