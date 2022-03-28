package exceptions

type IllegalParameterException struct {
	RuntimeException
}

func NewIllegalParameterException(message any, config *ExceptionConfig) *IllegalParameterException {
	return &IllegalParameterException{
		NewRuntimeException(
			message,
			"exception caused ElementNotFoundException:",
			config.AddSkipStack(1),
		),
	}
}
