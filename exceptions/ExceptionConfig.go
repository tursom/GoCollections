package exceptions

type ExceptionConfig struct {
	SkipStack     int
	GetStackTrace bool
	Cause         any
	ExceptionName string
}

func Cfg() *ExceptionConfig {
	return DefaultExceptionConfig()
}

func DefaultExceptionConfig() *ExceptionConfig {
	return &ExceptionConfig{
		SkipStack:     0,
		GetStackTrace: true,
		Cause:         nil,
	}
}

func (c *ExceptionConfig) SetSkipStack(skipStack int) *ExceptionConfig {
	if c == nil {
		return &ExceptionConfig{SkipStack: skipStack}
	}
	c.SkipStack = skipStack
	return c
}

func (c *ExceptionConfig) SetGetStackTrace(getStackTrace bool) *ExceptionConfig {
	if c == nil {
		return &ExceptionConfig{GetStackTrace: getStackTrace}
	}
	c.GetStackTrace = getStackTrace
	return c
}

func (c *ExceptionConfig) SetCause(cause any) *ExceptionConfig {
	if c == nil {
		return &ExceptionConfig{Cause: cause}
	}
	c.Cause = cause
	return c
}

func (c *ExceptionConfig) AddSkipStack(skipStack int) *ExceptionConfig {
	if c == nil {
		return &ExceptionConfig{SkipStack: skipStack}
	}
	c.SkipStack += skipStack
	return c
}

func (c *ExceptionConfig) SetExceptionName(exceptionName string) *ExceptionConfig {
	if c == nil {
		return &ExceptionConfig{ExceptionName: exceptionName}
	}
	c.ExceptionName = exceptionName
	return c
}
