package implementation

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"hello-world/consumer/broker"
	"hello-world/consumer/broker/mocks"
)

//NewMockConsumer ...
func NewMockConsumer(logger *zap.SugaredLogger, rabbitMQURL string, queueName string) (*Consumer, error) {
	_ = rabbitMQURL
	return &Consumer{
		Logger:    logger,
		HostDsn:   "testHost",
		QueueName: queueName,
		DialFunc: func(url string) (broker.AmqpConnection, error) {

			conn := new(mocks.AmqpConnection)

			ch := new(mocks.AmqpChannel)

			q := amqp.Queue{Name: "test queue"}

			var table amqp.Table

			// correct data
			ch.On("QueueDeclare",
				queueName, // name
				false,     // durable
				false,     // delete when unused
				false,     // exclusive
				false,     // no-wait
				table,     // arguments
			).Return(q, nil)

			msg := make(chan amqp.Delivery)
			close(msg)
			var msgChan <-chan amqp.Delivery = msg

			ch.On("Consume",
				queueName, // queue
				"",        // consumer
				true,      // auto-ack
				false,     // exclusive
				false,     // no-local
				false,     // no-wait
				table,     // args
			).Return(msgChan, nil)

			ch.On("Close").Return(nil)

			conn.On("Close").Return(nil)

			conn.On("Channel").Return(ch, nil)

			return conn, nil
		},
	}, nil
}
