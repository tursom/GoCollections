package exceptions

type WrongCallHostException struct {
	RuntimeException
}

func NewWrongCallHostException(message string) WrongCallHostException {
	return WrongCallHostException{
		NewRuntimeException(nil, message, DefaultExceptionConfig().AddSkipStack(1)),
	}
}
