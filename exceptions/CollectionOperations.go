package exceptions

var ElementFound = NewRuntimeException("", "element found", DefaultExceptionConfig().SetGetStackTrace(false))
var ElementNotFound = NewRuntimeException("", "element not found", DefaultExceptionConfig().SetGetStackTrace(false))
var CollectionLoopFinished = NewRuntimeException("", "collection loop finished", DefaultExceptionConfig().SetGetStackTrace(false))
