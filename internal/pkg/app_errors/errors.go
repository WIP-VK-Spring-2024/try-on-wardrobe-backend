package app_errors

import (
	"errors"
	"log"
	"net/http"
	"runtime"
)

var (
	ErrNotFound            = errors.New("requested resource does not exist")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrAlreadyExists       = errors.New("resource already exists")
	ErrTokenMalformed      = errors.New("token malformed or missing")
	ErrInvalidSignature    = errors.New("token has invalid signature")
	ErrTokenExpired        = errors.New("token has expired")
	ErrUnimplemented       = errors.New("method unimplemented")
	ErrNoRelatedEntity     = errors.New("related resource not found")
	ErrConstraintViolation = errors.New("constraint violated")
	ErrNotOwner            = errors.New("must be the owner to delete or edit this resource")
)

var (
	ErrBadRequest = &ResponseError{
		Msg:  "bad request",
		Code: http.StatusBadRequest,
	}

	// ErrNotOwner = &ResponseError{
	// 	Msg:  "must be the owner to delete or edit this resource",
	// 	Code: http.StatusForbidden,
	// }

	ErrUnauthorized = &ResponseError{
		Msg:  "credentials missing or invalid",
		Code: http.StatusUnauthorized,
	}

	ErrUserIdInvalid = &ResponseError{
		Code: http.StatusBadRequest,
		Msg:  "user ID is missing or isn't a valid uuid",
	}

	ErrClothesIdInvalid = &ResponseError{
		Code: http.StatusBadRequest,
		Msg:  "clothes ID is missing or isn't a valid uuid",
	}

	ErrUserImageIdInvalid = &ResponseError{
		Code: http.StatusBadRequest,
		Msg:  "user image ID is missing or isn't a valid uuid",
	}

	ErrTryOnIdInvalid = &ResponseError{
		Code: http.StatusBadRequest,
		Msg:  "try on result ID is missing or isn't a valid uuid",
	}

	ErrOutfitIdInvalid = &ResponseError{
		Code: http.StatusBadRequest,
		Msg:  "outfit ID is missing or isn't a valid uuid",
	}
)

type InternalError struct {
	Err  error
	File string
	Line int
}

func (err *InternalError) Error() string {
	return err.Err.Error()
}

//easyjson:json
type ResponseError struct {
	Code int `json:"-"`
	Msg  string
}

func (err ResponseError) Error() string {
	return err.Msg
}

func New(err error) error {
	var code int

	switch {
	case errors.Is(err, ErrAlreadyExists):
		code = http.StatusConflict

	case errors.Is(err, ErrNotOwner):
		code = http.StatusForbidden

	case errors.Is(err, ErrNotFound) || errors.Is(err, ErrNoRelatedEntity):
		code = http.StatusNotFound

	case errors.Is(err, ErrInvalidCredentials):
		code = http.StatusForbidden

	case errors.Is(err, ErrConstraintViolation):
		code = http.StatusBadRequest

	default:
		log.Println("Nothing matches, internal error")
		_, file, line, _ := runtime.Caller(1)
		return &InternalError{
			Err:  err,
			File: file,
			Line: line,
		}
	}

	log.Println("Code is", code)

	return &ResponseError{
		Code: code,
		Msg:  err.Error(),
	}
}
