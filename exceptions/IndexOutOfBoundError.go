package exceptions

type IndexOutOfBound struct {
	RuntimeException
}

func NewIndexOutOfBound(message string, config *ExceptionConfig) *IndexOutOfBound {
	return &IndexOutOfBound{
		NewRuntimeException(message, config.AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.IndexOutOfBound")),
	}
}

func CatchIndexOutOfBound[T any](f func() T, config *ExceptionConfig) (r T, err Exception) {
	defer func() {
		r := recover()
		if r != nil {
			err = NewIndexOutOfBound("", config.AddSkipStack(3).SetCause(r))
		}
	}()
	r = f()
	return
}
