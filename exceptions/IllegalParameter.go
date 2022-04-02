package exceptions

type IllegalParameterException struct {
	RuntimeException
}

func NewIllegalParameterException(message string, config *ExceptionConfig) *IllegalParameterException {
	return &IllegalParameterException{
		NewRuntimeException(message, config.AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.IllegalParameterException")),
	}
}
