package exceptions

var ElementFound = NewRuntimeException("", DefaultExceptionConfig().SetGetStackTrace(false).SetExceptionName("github.com.tursom.GoCollections.exceptions.ElementFound"))
var ElementNotFound = NewElementNotFoundException("", nil)
var CollectionLoopFinished = NewRuntimeException("", DefaultExceptionConfig().SetGetStackTrace(false).SetExceptionName("github.com.tursom.GoCollections.exceptions.CollectionLoopFinished"))
