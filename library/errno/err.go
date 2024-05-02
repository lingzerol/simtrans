package errno

type CodeErr struct {
	code   int
	errmsg string
	err    error
}

func (c *CodeErr) Error() string {
	if c == nil || c.err == nil {
		return ""
	}
	return c.err.Error()
}

func (c *CodeErr) Code() int {
	if c == nil {
		return 0
	}
	return c.code
}

func (c *CodeErr) Message() string {
	if c == nil {
		return ""
	}
	return c.errmsg
}
