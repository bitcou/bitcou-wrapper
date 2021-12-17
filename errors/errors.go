package wrapper_errors

import "errors"

func New(err error) map[string]string {
	return map[string]string{
		"errorMessage": err.Error(),
	}
}

var (
	// ErrorInternalServer internal server error.
	ErrorInternalServer = errors.New("internal server error")
	// ErrorProductNotFound invalid id, product not found.
	ErrorProductNotFound = errors.New("the requested product was not found")
)
