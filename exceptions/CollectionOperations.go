package exceptions

var ElementFound = ElementFoundThrowable{}
var ElementNotFound = ElementFoundThrowable{}
var CollectionLoopFinished = CollectionLoop{}

type ElementFoundThrowable struct {
}

func (e ElementFoundThrowable) Error() string {
	return "element found"
}

type ElementNotFoundThrowable struct {
}

func (e ElementNotFoundThrowable) Error() string {
	return "element not found"
}

type CollectionLoop struct {
}

func (e CollectionLoop) Error() string {
	return "collection loop finished"
}
