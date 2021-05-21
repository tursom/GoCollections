package exceptions

type NPE struct {
	RuntimeException
}

func NewNPE(message interface{}, getStackTrace bool) *NPE {
	return &NPE{
		NewRuntimeException(
			message,
			"exception caused NullPointerException:",
			getStackTrace,
			nil,
		),
	}
}
