package schemas

type PostData struct {
	ID    string
	Title string
	Link  string
}

type Post struct {
	Source
	ID         string     `json:"id"`
	SchemaType SchemaType `json:"schemaType"`
	Title      string     `json:"title"`
	Link       string     `json:"link"`
}

func NewPost(source Source, data PostData) Post {
	return Post{
		Source:     source,
		SchemaType: SCHEMA_TYPE_POST,
		ID:         data.ID,
		Title:      data.Title,
		Link:       data.Link,
	}
}
