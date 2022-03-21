package exceptions

type PackageException struct {
	RuntimeException
	err interface{}
}

func NewPackageException(err interface{}, exceptionMessage string, config *ExceptionConfig) *PackageException {
	config.AddSkipStack(1)
	return &PackageException{
		RuntimeException: NewRuntimeException(err, exceptionMessage, config),
		err:              err,
	}
}

func (p *PackageException) Err() interface{} {
	return p.err
}

func UnpackException(err interface{}) interface{} {
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
