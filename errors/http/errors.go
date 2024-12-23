package httperrors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	httpCode int
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Details  []any  `json:"details,omitempty"`
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("http error: httpCode=%d message=%s code=%d", he.httpCode, he.Message, he.Code)
}

func (he *HttpError) HttpCode() int {
	return he.httpCode
}

func New(httpCode int, msg string) *HttpError {
	he, err := FromHttpCode(httpCode, msg)
	if err != nil {
		panic(err)
	}

	return he
}

func FromHttpCode(httpCode int, msg string) (*HttpError, error) {
	statusTxt := http.StatusText(httpCode)
	if statusTxt == "" {
		return nil, fmt.Errorf("invalid http code: httpCode=%d", httpCode)
	}

	if msg == "" {
		msg = statusTxt
	}

	return &HttpError{
		httpCode: httpCode,
		Code:     httpCode,
		Message:  msg,
	}, nil
}

func (he *HttpError) WithDetails(details ...any) *HttpError {
	he.Details = append(he.Details, details...)
	return he
}
