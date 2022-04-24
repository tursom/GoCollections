package main

import (
	_ "github.com/tursom/GoCollections/collections"
	_ "github.com/tursom/GoCollections/concurrent"
	_ "github.com/tursom/GoCollections/concurrent/collections"
	"github.com/tursom/GoCollections/exceptions"
	_ "github.com/tursom/GoCollections/lang"
	_ "github.com/tursom/GoCollections/lang/atomic"
	_ "github.com/tursom/GoCollections/util"
)

func main() {
	exceptions.NewRuntimeException("test2", exceptions.DefaultExceptionConfig().SetCause(1)).PrintStackTrace()
}
