package schemas

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		data string
		want interface{}
	}{
		{
			data: `{
			  "source": { "id": "1", "name": "source name" },
			  "schemaType": "POST",
			  "title": "my title",
			  "link": "my link"
			}`,
			want: NewPost(
				Source{ID: "1", Name: "source name"},
				PostData{Title: "my title", Link: "my link"},
			),
		},
		{
			data: `{"schemaType": "ANY"}`,
			want: NewAny(),
		},
	}

	for _, tt := range tests {
		data := []byte(tt.data)
		result, err := Decode(data)
		require.NoError(t, err)
		require.Equal(t, tt.want, result)
	}
}
