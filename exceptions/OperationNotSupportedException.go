package exceptions

type OperationNotSupportedException struct {
	RuntimeException
}

func NewOperationNotSupportedException(message string, config *ExceptionConfig) *OperationNotSupportedException {
	return &OperationNotSupportedException{
		NewRuntimeException(message, config.AddSkipStack(1).SetExceptionName("OperationNotSupportedException")),
	}
}
