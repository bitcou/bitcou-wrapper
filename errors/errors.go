package wrapper_errors

import (
	"errors"
	"strings"
)

func New(err error) map[string]string {
	return map[string]string{
		"errorMessage": err.Error(),
	}
}

func GetGqlError(err error) error {
	fi := strings.TrimPrefix(err.Error(), "Message: ")
	lastInd := strings.Index(fi, ", Locations: ")
	return errors.New(fi[:lastInd])
}

var (
	// ErrorInternalServer internal server error.
	ErrorInternalServer = errors.New("internal server error")
	// ErrorProductNotFound invalid id, product not found.
	ErrorProductNotFound = errors.New("the requested product was not found")
	// ErrorInvalidCategory invalid category id.
	ErrorInvalidCategory = errors.New("invalid category id")
	// ErrorInvalidCategory invalid country id.
	ErrorInvalidCountry = errors.New("no country was found with that id")
	// ErrorInvalidCategory invalid country id.
	ErrorInvalidJson = errors.New("there is an error in the body, please check each field and its data type")
)
