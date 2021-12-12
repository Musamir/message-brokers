//Proto generate command
//go:generate protoc --go_out=. --go_opt=paths=source_relative request.proto

package decoding

import (
	"github.com/golang/protobuf/proto"
)

//Decoder ...
type Decoder struct{}

//NewDecoder ...
func NewDecoder() *Decoder {
	return &Decoder{}
}

//Decode ...
func (p *Decoder) Decode(data *Request) ([]byte, error) {
	marshaled, err := proto.Marshal(data)
	if err != nil {
		return []byte{}, nil
	}
	return marshaled, nil
}
