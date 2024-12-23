package httperrors

import (
	"fmt"
	domainerrors "github.com/tniah/iam-domain/errors"
	"net/http"
)

type HttpError struct {
	httpCode int
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Details  []any  `json:"details"`
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("http error: httpCode=%d, message=%s, code=%d", he.httpCode, he.Message, he.Code)
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

func (he *HttpError) WithErrorInfo(de *domainerrors.ErrorInfo) *HttpError {
	for i, detail := range he.Details {
		switch detail.(type) {
		case *domainerrors.ErrorInfo:
			he.Details[i] = de
			return he
		default:
			continue
		}
	}

	he.Details = append(he.Details, de)
	return he
}

func (he *HttpError) WithInvalidArgument(de *domainerrors.InvalidArgument) *HttpError {
	for i, detail := range he.Details {
		switch detail.(type) {
		case *domainerrors.InvalidArgument:
			he.Details[i] = de
			return he
		default:
			continue
		}
	}
	he.Details = append(he.Details, de)
	return he
}
