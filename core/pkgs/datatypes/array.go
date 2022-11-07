package datatypes

import (
	"database/sql/driver"
	"encoding/json"
	"io"

	"gitlab.com/meta-node/mail/core/pkgs/errors"
)

// StringArray represent list of string
type StringArray []string

// Scan implements the sql.Scanner interface.
func (m *StringArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, m)
}

// Value implements the driver.Valuer interface.
func (m StringArray) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *StringArray) UnmarshalGQL(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, m)
}

func (m StringArray) MarshalGQL(w io.Writer) {
	bytes, _ := json.Marshal(m)
	w.Write(bytes)
}

// Int64Array represent list of int64
type Int64Array []int64

// Scan implements the sql.Scanner interface.
func (m *Int64Array) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, m)
}

// Value implements the driver.Valuer interface.
func (m Int64Array) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Int64Array) UnmarshalGQL(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, m)
}

func (m Int64Array) MarshalGQL(w io.Writer) {
	bytes, _ := json.Marshal(m)
	w.Write(bytes)
}
