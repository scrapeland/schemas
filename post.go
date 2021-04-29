package schemas

type PostData struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type Post struct {
	ID         string     `json:"id"`
	SchemaType SchemaType `json:"schemaType"`
	Source     Source
	Data       PostData
}

func NewPost(id string, source Source, data PostData) Post {
	return Post{
		SchemaType: SCHEMA_TYPE_POST,
		ID:         id,
		Source:     source,
		Data:       data,
	}
}
