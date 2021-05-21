package exceptions

type IndexOutOfBound struct {
	RuntimeException
}

func NewIndexOutOfBound(message interface{}, getStackTrace bool) IndexOutOfBound {
	return IndexOutOfBound{
		NewRuntimeException(
			message,
			"exception caused IndexOutOfBound:",
			getStackTrace,
			nil,
		),
	}
}
