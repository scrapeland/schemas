package schemas

type Post struct {
	SchemaType SchemaType `json:"schemaType"`
	Title      string     `json:"title"`
	Link       string     `json:"link"`
}

func NewPost(title, link string) Post {
	return Post{
		SchemaType: SCHEMA_TYPE_POST,
		Title:      title,
		Link:       link,
	}
}
