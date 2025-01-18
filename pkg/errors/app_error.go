package errors

import (
	"encoding/json"
	"net/http"
	"strings"
)

type AppError struct {
	ID string
	// The params for i18n
	params map[string]any
	// Message to be display to the end user without debugging information
	Message string `json:"message"`
	// Internal error string to help the developer
	DetailedError string `json:"detailed_error"`
	// The RequestID that's also set in the header
	RequestID string `json:"request_id,omitempty"`
	// The http status code
	StatusCode int `json:"status_code,omitempty"`
	// The biz code
	Code int `json:"code,omitempty"`
	// The function where it happened in the form of Struct.Func
	Where string `json:"-"`
	// The wrapped error
	wrapped error
}

func NewAppError(
	where string,
	id string,
	params map[string]any,
	details string,
	status int,
	code int,
) *AppError {
	ap := &AppError{
		ID:            id,
		params:        params,
		Message:       id,
		Where:         where,
		DetailedError: details,
		StatusCode:    status,
		Code:          code,
	}

	return ap
}

const maxErrorLength = 1024

func (er *AppError) Error() string {
	var sb strings.Builder

	// render the error information
	if er.Where != "" {
		_, _ = sb.WriteString(er.Where)
		_, _ = sb.WriteString(": ")
	}

	// only render the detailed error when it's present
	if er.DetailedError != "" {
		_, _ = sb.WriteString(er.DetailedError)
	}

	// render the wrapped error
	err := er.wrapped
	if err != nil {
		_, _ = sb.WriteString(", ")
		_, _ = sb.WriteString(err.Error())
	}

	res := sb.String()
	if len(res) > maxErrorLength {
		res = res[:maxErrorLength] + "..."
	}
	return res
}

func (er *AppError) ToJSON() string {
	// turn the wrapped error into a detailed message
	detailed := er.DetailedError
	defer func() {
		er.DetailedError = detailed
	}()

	er.wrappedToDetailed()

	b, err := json.Marshal(er)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (er *AppError) wrappedToDetailed() {
	if er.wrapped == nil {
		return
	}

	if er.DetailedError != "" {
		er.DetailedError += ", "
	}

	er.DetailedError += er.wrapped.Error()
}

func (er *AppError) Unwrap() error {
	return er.wrapped
}

func (er *AppError) Wrap(err error) *AppError {
	er.wrapped = err

	return er
}

func (er *AppError) WipeDetailed() {
	er.wrapped = nil
	er.DetailedError = ""
}

func MakeBindError(where string, details string) *AppError {
	return NewAppError(where, "app.bind.error", nil, details, http.StatusUnprocessableEntity, ErrBind)
}

func MakeUnknownError(where string, details string) *AppError {
	return NewAppError(where, "app.unknown.error", nil, details, http.StatusUnprocessableEntity, ErrUnknown)
}

func MakeTokenInvalid(where string) *AppError {
	return NewAppError(where, "app.token.invalid", nil, "", http.StatusUnauthorized, ErrTokenInvalid)
}
