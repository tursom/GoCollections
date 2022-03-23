package exceptions

type NPE struct {
	RuntimeException
}

func NewNPE(message any, config *ExceptionConfig) *NPE {
	config.AddSkipStack(1)
	return &NPE{
		NewRuntimeException(
			message,
			"exception caused NullPointerException:",
			config,
		),
	}
}
