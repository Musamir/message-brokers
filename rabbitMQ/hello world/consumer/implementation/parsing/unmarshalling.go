//Proto generate command
//go:generate protoc --go_out=. --go_opt=paths=source_relative response.proto

package parsing

import (
	"github.com/golang/protobuf/proto"
)

//Parser ...
type Parser struct{}

//NewParser ...
func NewParser() *Parser {
	return &Parser{}
}

//Parse ...
func (p *Parser) Parse(marshaled []byte, toUnmarshal *Response) error {
	if err := proto.Unmarshal(marshaled, toUnmarshal); err != nil {
		return err
	}
	return nil
}
