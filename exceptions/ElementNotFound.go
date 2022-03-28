package exceptions

type ElementNotFoundException struct {
	RuntimeException
}

func NewElementNotFoundException(message any, config *ExceptionConfig) *ElementNotFoundException {
	return &ElementNotFoundException{
		NewRuntimeException(
			message,
			"exception caused ElementNotFoundException:",
			config.AddSkipStack(1),
		),
	}
}
