package schemas

type Any struct {
	ID         string     `json:"id"`
	SchemaType SchemaType `json:"schemaType"`
	Source
	Data interface{}
}

func NewAny(id string, source Source, data interface{}) Any {
	return Any{
		SchemaType: SCHEMA_TYPE_ANY,
		ID:         id,
		Source:     source,
		Data:       data,
	}
}
