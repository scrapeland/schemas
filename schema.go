package schemas

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type SchemaType string

const (
	SCHEMA_TYPE_POST SchemaType = "POST"
	SCHEMA_TYPE_ANY  SchemaType = "ANY"
)

var AllSchemaTypes = []SchemaType{
	SCHEMA_TYPE_POST,
	SCHEMA_TYPE_ANY,
}

func (s SchemaType) IsValid() bool {
	switch s {
	case SCHEMA_TYPE_POST, SCHEMA_TYPE_ANY:
		return true
	default:
		return false
	}
}

func Decode(data []byte) (interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	schemaTypeIface, ok := m["schemaType"]
	if !ok {
		return nil, fmt.Errorf("there is no type property")
	}
	schemaTypeStr, ok := schemaTypeIface.(string)
	if !ok {
		return nil, fmt.Errorf("schema type value should be string")
	}

	schemaType := SchemaType(schemaTypeStr)

	var result interface{}

	switch schemaType {
	case SCHEMA_TYPE_POST:
		var post Post
		err = mapstructure.Decode(m, &post)
		result = post
	case SCHEMA_TYPE_ANY:
		var any Any
		err = mapstructure.Decode(m, &any)
		result = any
	}

	return result, err
}

type Source struct {
	ID   string
	Name string
}
