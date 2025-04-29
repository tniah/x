package sqlxx

import (
	"database/sql"
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

type NullString string

func NewNullString(s string) NullString {
	return NullString(s)
}

func (n *NullString) Scan(value interface{}) error {
	var v sql.NullString
	if err := (&v).Scan(value); err != nil {
		return err
	}

	*n = NullString(v.String)
	return nil
}

func (ns NullString) Value() (driver.Value, error) {
	if len(ns) == 0 {
		return sql.NullString{}.Value()
	}

	return sql.NullString{Valid: true, String: string(ns)}.Value()
}

func (ns NullString) String() string {
	return string(ns)
}

type NullInt64 struct {
	Int   int64
	Valid bool
}

func (ns *NullInt64) Scan(value interface{}) error {
	var d = sql.NullInt64{}
	if err := (&d).Scan(value); err != nil {
		return err
	}

	ns.Int = d.Int64
	ns.Valid = d.Valid
	return nil
}

func (ns NullInt64) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.Int, nil
}

func (ns NullInt64) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ns.Int)
}

func (ns *NullInt64) UnmarshalJSON(data []byte) error {
	if ns == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}

	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	ns.Valid = true
	return errors.WithStack(json.Unmarshal(data, &ns.Int))
}
