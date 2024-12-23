package httperrors

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ErrorDetails struct{}

type ErrorInfo struct {
	Reason   string            `json:"reason"`
	Domain   string            `json:"domain"`
	Metadata map[string]string `json:"metadata"`
}

func (e *ErrorInfo) MarshalJSON() ([]byte, error) {
	type Alias ErrorInfo
	bytes, err := json.Marshal(&struct {
		Type string `json:"@type"`
		*Alias
	}{
		Type:  reflect.TypeOf(e).String(),
		Alias: (*Alias)(e),
	})
	return bytes, err
}

func (e *ErrorInfo) Error() string {
	return fmt.Sprintf("reason=%s | domain=%s | metadata=%v", e.Reason, e.Domain, e.Metadata)
}

type InvalidArgument struct {
	Fields []*InvalidField `json:"fields"`
}

type InvalidField struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}
