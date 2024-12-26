package httperrors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	httpCode int
	// reason - the reason of the error.
	reason  string
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details []any  `json:"details,omitempty"`
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("http error: httpCode=%d message=%s code=%d", he.httpCode, he.Message, he.Code)
}

func (he *HttpError) HttpCode() int {
	return he.httpCode
}

func (he *HttpError) Reason() string {
	return he.reason
}

func New(httpCode int, msg, reason string) *HttpError {
	he, err := FromHttpCode(httpCode, msg, reason)
	if err != nil {
		panic(err)
	}

	return he
}

func FromHttpCode(httpCode int, msg, reason string) (*HttpError, error) {
	statusTxt := http.StatusText(httpCode)
	if statusTxt == "" {
		return nil, fmt.Errorf("invalid http code: httpCode=%d", httpCode)
	}

	if msg == "" {
		msg = statusTxt
	}

	return &HttpError{
		httpCode: httpCode,
		reason:   reason,
		Code:     httpCode,
		Message:  msg,
	}, nil
}

func (he *HttpError) WithDetails(details ...any) {
	he.Details = append(he.Details, details...)
}
