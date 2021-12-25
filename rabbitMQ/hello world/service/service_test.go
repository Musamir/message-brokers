package service

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	consumerImp "hello-world/consumer/implementation"
	producerImp "hello-world/producer/implementation"
	"testing"
)

//TestApplication_Run ...
func TestApplication_Run(t *testing.T) {
	p, err := producerImp.NewMockProducer(zaptest.NewLogger(t).Sugar(), "test", "test queue")
	assert.Equal(t, nil, err, "err must be nil")

	c, err := consumerImp.NewMockConsumer(zaptest.NewLogger(t).Sugar(), "test", "test queue")
	assert.Equal(t, nil, err, "err must be nil")

	//mockService ...
	mServ := &Application{
		Logger:   zaptest.NewLogger(t).Sugar(),
		producer: p,
		consumer: c,
	}

	testData = []string{"test"}

	err = mServ.Run()
	assert.Equal(t, nil, err, "err must be nil")
}
