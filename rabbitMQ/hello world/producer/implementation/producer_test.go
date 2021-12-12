package implementation

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"hello-world/producer/broker"
	"hello-world/producer/broker/mocks"
	"testing"
)

//TestConsumer_Receive ...
func TestConsumer_Receive(t *testing.T) {

	//m - mockConsumer
	m := Producer{
		Logger:    zaptest.NewLogger(t).Sugar(),
		HostDsn:   "testHost",
		QueueName: "test queue",
		DialFunc: func(url string) (broker.AmqpConnection, error) {

			conn := new(mocks.AmqpConnection)

			ch := new(mocks.AmqpChannel)

			q := amqp.Queue{Name: "test queue"}

			var table amqp.Table

			// correct data
			ch.On("QueueDeclare",
				"test queue", // name
				false,        // durable
				false,        // delete when unused
				false,        // exclusive
				false,        // no-wait
				table,        // arguments
			).Return(q, nil)

			data := []byte{0xa, 0x4, 0x74, 0x65, 0x73, 0x74} // marshaled Example (Message = "test")

			ch.On("Publish",
				"",           // exchange
				"test queue", // routing key
				false,        // mandatory
				false,        // immediate
				amqp.Publishing{
					ContentType: "application/x-protobuf",
					Body:        data,
				},
			).Return(nil)

			ch.On("Close").Return(nil)

			conn.On("Close").Return(nil)

			conn.On("Channel").Return(ch, nil)

			return conn, nil
		},
	}

	err := m.connect()

	assert.Equal(t, nil, err, "err must be nil")

	err = m.Send("test")

	assert.Equal(t, nil, err, "err must be nil")
}
