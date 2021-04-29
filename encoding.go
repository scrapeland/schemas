package schemas

import "github.com/mitchellh/mapstructure"

func Encode(id string, source Source, schemaType SchemaType, data interface{}) (schema interface{}, err error) {
	switch schemaType {

	case SCHEMA_TYPE_POST:
		var postData PostData
		err = mapstructure.Decode(data, &postData)
		schema = NewPost(id, source, postData)

	case SCHEMA_TYPE_ANY:
		schema = NewAny(id, source, data)
	}

	return
}
