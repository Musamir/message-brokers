package implementation

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"hello-world/consumer/broker"
	"hello-world/consumer/broker/mocks"
	"testing"
)

//TestConsumer_Receive ...
func TestConsumer_Receive(t *testing.T) {

	//m - mockConsumer
	m := Consumer{
		Logger:    zaptest.NewLogger(t).Sugar(),
		HostDsn:   "testHost",
		QueueName: "test",
		DialFunc: func(url string) (broker.AmqpConnection, error) {

			conn := new(mocks.AmqpConnection)

			ch := new(mocks.AmqpChannel)

			q := amqp.Queue{Name: "test"}

			var table amqp.Table

			// correct data
			ch.On("QueueDeclare",
				"test", // name
				false,  // durable
				false,  // delete when unused
				false,  // exclusive
				false,  // no-wait
				table,  // arguments
			).Return(q, nil)

			msg := make(chan amqp.Delivery)
			close(msg)
			var msgChan <-chan amqp.Delivery = msg

			ch.On("Consume",
				"test", // queue
				"",     // consumer
				true,   // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				table,  // args
			).Return(msgChan, nil)

			ch.On("Close").Return(nil)

			conn.On("Close").Return(nil)

			conn.On("Channel").Return(ch, nil)

			return conn, nil
		},
	}

	err := m.Receive()

	assert.Equal(t, nil, err, "err must be nil")
}
