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
