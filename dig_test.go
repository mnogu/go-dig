package dig

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringDig(t *testing.T) {
	var jsonBlob = []byte(`{"foo": {"bar": {"baz": 1}}}`)
	var h map[string]interface{}
	if err := json.Unmarshal(jsonBlob, &h); err != nil {
		t.Fatal(err)
	}
	success, err := Dig(h, "foo", "bar", "baz")
	assert.Equal(t, float64(1), success, "foo.bar.baz should be 1")
	assert.Nil(t, err)

	failure, err := Dig(h, "foo", "qux", "quux")
	assert.Nil(t, failure)
	assert.NotNil(t, err)
}

func TestIntDig(t *testing.T) {
	var jsonBlob = []byte(`{"foo": [10, 11, 12]}`)
	var h map[string]interface{}
	if err := json.Unmarshal(jsonBlob, &h); err != nil {
		t.Fatal(err)
	}
	success, err := Dig(h, "foo", 1)
	assert.Equal(t, float64(11), success, "foo.bar.baz should be 1")
	assert.Nil(t, err)

	failure, err := Dig(h, "foo", 1, 0)
	assert.Nil(t, failure)
	assert.NotNil(t, err)

	failure2, err := Dig(h, "foo", "bar")
	assert.Nil(t, failure2)
	assert.NotNil(t, err)
}
