package main

import "github.com/tursom/GoCollections/exceptions"

func main() {
	exceptions.NewRuntimeException("test2", exceptions.DefaultExceptionConfig().SetCause(1)).PrintStackTrace()
}
