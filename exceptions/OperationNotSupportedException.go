package exceptions

type OperationNotSupportedException struct {
	RuntimeException
}

func NewOperationNotSupportedException(message any, config *ExceptionConfig) *OperationNotSupportedException {
	return &OperationNotSupportedException{
		NewRuntimeException(
			message,
			"exception caused OperationNotSupportedException:",
			config.AddSkipStack(1),
		),
	}
}
