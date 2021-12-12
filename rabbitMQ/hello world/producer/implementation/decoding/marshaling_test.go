package decoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Example ...
var Example = Request{Message: "test"}

//protoMarshaled - marshaled Example
var protoMarshaled = []byte{0xa, 0x4, 0x74, 0x65, 0x73, 0x74}

func TestDecoder_Decode(t *testing.T) {
	decoder := Decoder{}

	expected := protoMarshaled

	actualMarshaled, err := decoder.Decode(&Example)
	assert.Equal(t, nil, err, "must be nil!!!")
	assert.Equal(t, expected, actualMarshaled, "not equal!!!")
}
