package exceptions

type IndexOutOfBound struct {
	RuntimeException
}

func NewIndexOutOfBound(message any, config *ExceptionConfig) *IndexOutOfBound {
	config.AddSkipStack(1)
	return &IndexOutOfBound{
		NewRuntimeException(
			message,
			"exception caused IndexOutOfBound:",
			config,
		),
	}
}

func CatchIndexOutOfBound[T any](f func() T, config *ExceptionConfig) (r T, err Exception) {
	defer func() {
		r := recover()
		if r != nil {
			err = NewIndexOutOfBound(r, config.AddSkipStack(3))
		}
	}()
	r = f()
	return
}
