package exceptions

type OperationNotSupportedException struct {
	RuntimeException
}

func NewOperationNotSupportedException(message string, getStackTrace bool) OperationNotSupportedException {
	return OperationNotSupportedException{
		NewRuntimeException(
			message,
			"exception caused OperationNotSupportedException:",
			getStackTrace,
			nil,
		),
	}
}
