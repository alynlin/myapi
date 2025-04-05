package errors

type GenericError struct {
	Err    error
	Format string
	Code   int
	Args   []interface{}
}

func (e *GenericError) Error() string {
	return e.Err.Error()
}

func New(err error, code int, format string, args ...interface{}) *GenericError {
	return &GenericError{
		Code:   code,
		Err:    err,
		Format: format,
		Args:   args,
	}
}
