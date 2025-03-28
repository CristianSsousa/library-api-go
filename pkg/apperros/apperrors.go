package apperros

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your item already exists")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given param is not valid")
	// ErrUnathorized will throw if the user is not authorized
	ErrUnathorized = errors.New("Unauthorized")
	// ErrForbidden will throw if the user does not have the access
	ErrForbidden = errors.New("Forbidden")
	// ErrTimeOut will throw if the context is cancelled
	ErrTimeOut = errors.New("timeout")
)

func FormatValidationErrors(errs validator.ValidationErrors) []map[string]string {
	var errors []map[string]string
	for _, e := range errs {
		errorDetail := map[string]string{
			"field":   e.Field(),
			"message": fmt.Sprintf("O campo '%s' falhou na validação: %s", e.Field(), e.Tag()),
		}
		errors = append(errors, errorDetail)
	}
	return errors
}