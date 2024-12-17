package sqlxx

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type StringSliceJSON []string

func (m *StringSliceJSON) Scan(v interface{}) error {
	val := fmt.Sprintf("%s", v)
	if len(val) == 0 {
		val = "[]"
	}

	if parsed := gjson.Parse(val); parsed.Type == gjson.Null {
		val = "[]"
	} else if !parsed.IsArray() {
		return errors.Errorf("expected JSON value to be an array but got type: %s", parsed.Type.String())
	}

	return errors.WithStack(json.Unmarshal([]byte(val), &m))
}

func (m StringSliceJSON) Value() (driver.Value, error) {
	if len(m) == 0 {
		return "[]", nil
	}

	encoded, err := json.Marshal(&m)
	return string(encoded), errors.WithStack(err)
}
