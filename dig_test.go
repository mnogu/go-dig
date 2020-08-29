package dig

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringKey(t *testing.T) {
	var jsonBlob = []byte(`{"foo": {"bar": {"baz": 1}}}`)
	var v interface{}
	if err := json.Unmarshal(jsonBlob, &v); err != nil {
		t.Fatal(err)
	}
	success, err := Dig(v, "foo", "bar", "baz")
	assert.Equal(t, float64(1), success, "foo.bar.baz should be 1")
	assert.Nil(t, err)

	failure, err := Dig(v, "foo", "qux", "quux")
	assert.Nil(t, failure)
	assert.NotNil(t, err)
}

func TestIntKey(t *testing.T) {
	var jsonBlob = []byte(`{"foo": [10, 11, 12]}`)
	var v interface{}
	if err := json.Unmarshal(jsonBlob, &v); err != nil {
		t.Fatal(err)
	}
	success, err := Dig(v, "foo", 1)
	assert.Equal(t, float64(11), success, "foo.bar.baz should be 1")
	assert.Nil(t, err)

	failure, err := Dig(v, "foo", 1, 0)
	assert.Nil(t, failure)
	assert.NotNil(t, err)

	failure2, err := Dig(v, "foo", "bar")
	assert.Nil(t, failure2)
	assert.NotNil(t, err)
}
