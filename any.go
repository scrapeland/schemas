package schemas

type Any map[string]interface{}

func NewAny() Any {
	m := make(map[string]interface{})
	m["schemaType"] = SCHEMA_TYPE_ANY
	return Any(m)
}
