package implementation

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"hello-world/producer/broker"
	"hello-world/producer/broker/mocks"
)

//NewMockProducer ...
func NewMockProducer(logger *zap.SugaredLogger, rabbitMQURL string, queueName string) (*Producer, error) {
	_ = rabbitMQURL
	//m - mockConsumer
	m := Producer{
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

			data := []byte{0xa, 0x4, 0x74, 0x65, 0x73, 0x74} // marshaled Example (Message = "test")

			ch.On("Publish",
				"",        // exchange
				queueName, // routing key
				false,     // mandatory
				false,     // immediate
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

	return &m, m.connect()
}
