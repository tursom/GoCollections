package exceptions

type ElementNotFoundException struct {
	RuntimeException
}

func NewElementNotFoundException(message interface{}, config *ExceptionConfig) *ElementNotFoundException {
	config.AddSkipStack(1)
	return &ElementNotFoundException{
		NewRuntimeException(
			message,
			"exception caused ElementNotFoundException:",
			config,
		),
	}
}
