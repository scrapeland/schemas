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
			data: `{"schemaType": "POST", "title": "my title", "link": "my link"}`,
			want: NewPost("my title", "my link"),
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
