package functional

type (
	Function interface {
		invoke()
	}

	FunctionErred interface {
		invoke() error
	}

	Function1[A1 any] interface {
		invoke(a1 A1)
	}

	Function1Erred[A1 any] interface {
		invoke(a1 A1) error
	}
)
