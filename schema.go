package schemas

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type SchemaType string

const (
	SCHEMA_TYPE_POST SchemaType = "POST"
)

type Post struct {
	Type  SchemaType `json:"type"`
	Title string     `json:"title"`
	Link  string     `json:"link"`
}

func NewPost(title, link string) Post {
	return Post{
		Type:  SCHEMA_TYPE_POST,
		Title: title,
		Link:  link,
	}
}

func Decode(data []byte) (interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	schemaTypeIface, ok := m["type"]
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
	}

	return result, err
}
