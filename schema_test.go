package schemas

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		data string
		want interface{}
	}{
		{
			name: "post",
			data: `{
			  "id": "id-1",
			  "source": { "id": "1", "name": "source name" },
			  "schemaType": "POST",
			  "data": {
				"title": "my title",
				"link": "my link"
			  }
			}`,
			want: NewPost(
				"id-1",
				Source{ID: "1", Name: "source name"},
				PostData{Title: "my title", Link: "my link"},
			),
		},
		{
			name: "any",
			data: `{
			  "id": "id-1",
			  "source": { "id": "1", "name": "source name" },
			  "schemaType": "ANY",
			  "data": {
				"title": "my title",
				"link": "my link"
			  }
			}`,
			want: NewAny(
				"id-1",
				Source{ID: "1", Name: "source name"},
				map[string]interface{}{
					"title": "my title",
					"link":  "my link",
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []byte(tt.data)
			result, err := Decode(data)
			require.NoError(t, err)
			require.Equal(t, tt.want, result)
		})
	}
}
