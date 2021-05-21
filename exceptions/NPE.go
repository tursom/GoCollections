package exceptions

type NPE struct {
	RuntimeException
}

func NewNPE(message string, getStackTrace bool) *NPE {
	return &NPE{
		NewRuntimeException(
			message,
			"exception caused NullPointerException:",
			getStackTrace,
			nil,
		),
	}
}
