package schemas

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

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
