package error

type basedError struct {
	Code    string
	Message string
}

type DefinedError struct {
	basedError
	HttpStatus int
	Detail     interface{}
}

func (err DefinedError) Error() string {
	return err.Message
}

func NewDefinedError(code string, message string, httpStatus int, detail interface{}) DefinedError {
	return DefinedError{
		basedError: basedError{
			Code:    code,
			Message: message,
		},
		HttpStatus: httpStatus,
		Detail:     detail,
	}
}

type FieldValidationError struct {
	basedError
	FieldErrors map[string]string
}

func (err FieldValidationError) Error() string {
	return err.Message
}

func NewFieldValidationError(fieldErrors map[string]string) FieldValidationError {
	return FieldValidationError{
		basedError: basedError{
			Code:    "400",
			Message: "Field validation error",
		},
		FieldErrors: fieldErrors,
	}
}
