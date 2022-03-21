package exceptions

type OperationNotSupportedException struct {
	RuntimeException
}

func NewOperationNotSupportedException(message interface{}, config *ExceptionConfig) *OperationNotSupportedException {
	config.AddSkipStack(1)
	return &OperationNotSupportedException{
		NewRuntimeException(
			message,
			"exception caused OperationNotSupportedException:",
			config,
		),
	}
}
