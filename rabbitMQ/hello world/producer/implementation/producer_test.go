package implementation

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"testing"
)

//TestConsumer_Receive ...
func TestConsumer_Receive(t *testing.T) {
	m, err := NewMockProducer(zaptest.NewLogger(t).Sugar(), "test", "test queue")
	assert.Equal(t, nil, err, "err must be nil")

	err = m.Send("test")

	assert.Equal(t, nil, err, "err must be nil")
}
