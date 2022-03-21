package exceptions

type ExceptionConfig struct {
	SkipStack     int
	GetStackTrace bool
	Cause         interface{}
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
		return DefaultExceptionConfig().SetSkipStack(skipStack)
	}
	c.SkipStack = skipStack
	return c
}

func (c *ExceptionConfig) SetGetStackTrace(getStackTrace bool) *ExceptionConfig {
	if c == nil {
		return DefaultExceptionConfig().SetGetStackTrace(getStackTrace)
	}
	c.GetStackTrace = getStackTrace
	return c
}

func (c *ExceptionConfig) SetCause(cause interface{}) *ExceptionConfig {
	if c == nil {
		return DefaultExceptionConfig().SetCause(cause)
	}
	c.Cause = cause
	return c
}

func (c *ExceptionConfig) AddSkipStack(skipStack int) *ExceptionConfig {
	if c == nil {
		return DefaultExceptionConfig().AddSkipStack(skipStack)
	}
	c.SkipStack += skipStack
	return c
}
