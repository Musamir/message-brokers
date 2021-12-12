package parsing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Example ...
var Example = Response{Message: "test"}

//protoMarshaled - marshaled Example
var protoMarshaled = []byte{0xa, 0x4, 0x74, 0x65, 0x73, 0x74}

func TestParser_Parse(t *testing.T) {
	marshaledExample := protoMarshaled
	expected := Example.Message
	actualProtoUnmarshalled := &Response{}

	parser := Parser{}

	err := parser.Parse(marshaledExample, actualProtoUnmarshalled)
	assert.Equal(t, nil, err, "must be nil!!!")
	actualMessage := actualProtoUnmarshalled.Message
	assert.Equal(t, expected, actualMessage, "not equal!!!")
}
