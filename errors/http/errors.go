package httperrors

import (
	"fmt"
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

func NewInvalidArgumentError(msg string, domain, service string, violations ...*InvalidField) *HttpError {
	details := []any{
		&ErrorInfo{
			Reason: ReasonInvalidArgument,
			Domain: domain,
			Metadata: map[string]string{
				FieldService: service,
			},
		},
	}

	if len(violations) > 0 {
		details = append(details, &InvalidArgument{
			Fields: violations,
		})
	}

	return &HttpError{
		httpCode: http.StatusBadRequest,
		Code:     http.StatusBadRequest,
		Message:  msg,
		Details:  details,
	}
}

func NewInternalServerError(msg, domain, service string) *HttpError {
	details := []any{
		&ErrorInfo{
			Reason: ReasonInternalError,
			Domain: domain,
			Metadata: map[string]string{
				FieldService: service,
			},
		},
	}

	return &HttpError{
		httpCode: http.StatusInternalServerError,
		Code:     http.StatusInternalServerError,
		Message:  msg,
		Details:  details,
	}
}
