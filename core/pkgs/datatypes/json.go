package datatypes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"gitlab.com/meta-node/mail/core/pkgs/errors"
)

// JSON Custom type to mapping from json of postgres to map[string]bool
type JSON map[string]interface{}

// Scan implements the sql.Scanner interface.
func (m *JSON) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, m)
}

// GetValuePath return value
func (m JSON) GetValuePath(items ...string) (interface{}, error) {
	value := m
	for idx, item := range items {
		nestedValue, ok := value[item]
		if !ok {
			return nil, errors.New("Could not get field at %s", item)
		}

		if idx == len(items)-1 {
			return nestedValue, nil
		}

		nestedMap, ok := nestedValue.(map[string]interface{})
		if !ok {
			return nil, errors.New("Could not get nested item after field %s", item)
		}

		value = nestedMap
	}

	return nil, nil
}

// Value implements the driver.Valuer interface.
func (m JSON) Value() ([]byte, error) {
	return json.Marshal(m)
}

func (m JSON) ToString() (string, error) {
	s, err := json.Marshal(m)
	if err != nil {
		return "", errors.Wrap(err, "Convert to string failed")
	}
	return string(s), err
}

// ConvertStructToJSON convert struct to JSON
func ConvertStructToJSON(obj interface{}) (JSON, error) {
	if obj == nil {
		return JSON{}, nil
	}

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.Wrap(err, "Convert struct to json failed")
	}
	var result JSON
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, errors.Wrap(err, "Convert json data to JSON failed ")
	}

	return result, nil
}

func (m *JSON) UnmarshalGQL(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, m)
}

func (m JSON) MarshalGQL(w io.Writer) {
	bytes, _ := json.Marshal(m)
	w.Write(bytes)
}

func ParseToJson(s []byte) (map[string]interface{}, error) {
	var result JSON
	err := json.Unmarshal(s, &result)
	if err != nil {
		return nil, errors.Wrap(err, "Convert json data to JSON failed")
	}
	return result, nil
}

func ConvertJSONtoStruct(data JSON, value interface{}) {
	str, _ := data.Value()
	json.Unmarshal(str, &value)
}

func ConvertMaptoStruct(data interface{}, value interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &value)
	if err != nil {
		panic(err)
	}
	log.Println(value)
}
