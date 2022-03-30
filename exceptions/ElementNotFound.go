package exceptions

type ElementNotFoundException struct {
	RuntimeException
}

func NewElementNotFoundException(message string, config *ExceptionConfig) *ElementNotFoundException {
	return &ElementNotFoundException{
		NewRuntimeException(message, config.AddSkipStack(1).SetExceptionName("ElementNotFoundException")),
	}
}
