package implementation

import (
	"fmt"
	"go.uber.org/zap"
	"hello-world/consumer/broker"
	"hello-world/consumer/implementation/parsing"
)

//NewConsumer ...
func NewConsumer(logger *zap.SugaredLogger, rabbitMQURL string, queueName string) (*Consumer, error) {
	if logger == nil {
		return nil, fmt.Errorf("nil memory address")
	}
	return &Consumer{
		Logger:    logger,
		HostDsn:   rabbitMQURL,
		DialFunc:  broker.AmqpDialWrapper,
		QueueName: queueName,
	}, nil
}

//Consumer ...
type Consumer struct {
	Logger    *zap.SugaredLogger
	HostDsn   string
	DialFunc  broker.AmqpDial
	QueueName string
}

//Receive ...
func (c *Consumer) Receive() (err error) {

	conn, err := c.DialFunc(c.HostDsn)

	if err != nil {
		c.Logger.Errorf("Failed to dial with the broker: %v", err)
		return FailedToDial
	}

	defer func() {
		if err := conn.Close(); err != nil {
			c.Logger.Errorf("An error occurred while closing the connection: %s", err)
			err = SomethingWentWrong
		}
	}()

	ch, err := conn.Channel()
	if err != nil {
		c.Logger.Errorf("Couldn't open a unique, concurrent server channel: %v", err)
		return SomethingWentWrong
	}

	defer func() {
		if err := ch.Close(); err != nil {
			c.Logger.Errorf("Couldn't close ch: %v", err)
			err = SomethingWentWrong
		}
	}()

	q, err := ch.QueueDeclare(
		c.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		c.Logger.Errorf("Couldn't deaclare a queue: %v", err)
		return SomethingWentWrong
	}

	msg, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		c.Logger.Errorf("Couldn't start delivering messages: %v", err)
		return SomethingWentWrong
	}

	parser := parsing.NewParser()

	response := &parsing.Response{}
	for d := range msg {
		if err := parser.Parse(d.Body, response); err != nil {
			c.Logger.Errorf("couldn't unmarshal data %v", err)
		} else {
			c.Logger.Infof("got a message: %s", response.Message)
		}
	}

	return
}
