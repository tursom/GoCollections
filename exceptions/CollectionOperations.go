package exceptions

var ElementFound = NewRuntimeException("", DefaultExceptionConfig().SetGetStackTrace(false).SetExceptionName("ElementFound"))
var ElementNotFound = NewRuntimeException("", DefaultExceptionConfig().SetGetStackTrace(false).SetExceptionName("ElementNotFound"))
var CollectionLoopFinished = NewRuntimeException("", DefaultExceptionConfig().SetGetStackTrace(false).SetExceptionName("CollectionLoopFinished"))
