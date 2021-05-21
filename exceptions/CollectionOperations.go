package exceptions

var ElementFound = NewRuntimeException("", "element found", false, nil)
var ElementNotFound = NewRuntimeException("", "element not found", false, nil)
var CollectionLoopFinished = NewRuntimeException("", "collection loop finished", false, nil)
