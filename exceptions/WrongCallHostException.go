package exceptions

type WrongCallHostException struct {
	RuntimeException
}

func NewWrongCallHostException(message string) WrongCallHostException {
	return WrongCallHostException{
		NewRuntimeException(message, DefaultExceptionConfig().AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.WrongCallHostException")),
	}
}
