package exceptions

type ElementNotFoundException struct {
	RuntimeException
}

func NewElementNotFoundException(message string, getStackTrace bool) *ElementNotFoundException {
	return &ElementNotFoundException{
		NewRuntimeException(
			message,
			"exception caused ElementNotFoundException:",
			getStackTrace,
			nil,
		),
	}
}
