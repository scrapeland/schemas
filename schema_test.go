package schemas

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodePost(t *testing.T) {
	data := []byte(`{"schemaType": "POST", "title": "my title", "link": "my link"}`)
	result, err := Decode(data)
	require.NoError(t, err)
	post, ok := result.(Post)
	require.True(t, ok)
	p := NewPost("my title", "my link")
	require.Equal(t, p, post)
}
