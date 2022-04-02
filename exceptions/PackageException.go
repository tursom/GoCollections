package exceptions

import (
	"fmt"
	"reflect"
)

type PackageException struct {
	RuntimeException
	err any
}

func NewPackageException(err any, config *ExceptionConfig) *PackageException {
	message := ""
	switch e := err.(type) {
	case error:
		message = e.Error()
	default:
		message = fmt.Sprint(e)
	}
	t := reflect.TypeOf(err)
	message = fmt.Sprintf("%s (%s)", message, t.Name())
	return &PackageException{
		RuntimeException: NewRuntimeException(message, config.AddSkipStack(1).
			SetExceptionName("github.com.tursom.GoCollections.exceptions.PackageException")),
		err: err,
	}
}

func (p *PackageException) Err() any {
	return p.err
}

func UnpackException(err any) any {
	for err != nil {
		switch err.(type) {
		case *PackageException:
			err = err.(*PackageException).Err()
			return err
		case Exception:
			err = err.(Exception).Cause()
		default:
			return err
		}
	}
	return nil
}
