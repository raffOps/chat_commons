package errs

type Err struct {
	Message string
	Code    int
}

func (e *Err) Error() string {
	return e.Message
}
