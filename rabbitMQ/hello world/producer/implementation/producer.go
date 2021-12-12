package implementation

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"hello-world/producer/broker"
	"hello-world/producer/implementation/decoding"
)

//NewProducer ...
func NewProducer(logger *zap.SugaredLogger, rabbitMQURL string, queueName string) (*Producer, error) {
	if logger == nil {
		return nil, NilMemoryAddress
	}

	p := &Producer{
		Logger:    logger,
		HostDsn:   rabbitMQURL,
		DialFunc:  broker.AmqpDialWrapper,
		QueueName: queueName,
	}
	err := p.connect()
	if err != nil {
		return nil, err
	}

	return p, nil
}

//Producer ...
type Producer struct {
	Logger    *zap.SugaredLogger
	HostDsn   string
	DialFunc  broker.AmqpDial
	QueueName string
	conn      broker.AmqpConnection
	ch        broker.AmqpChannel
	decoder   *decoding.Decoder
}

//connect ...
func (p *Producer) connect() error {

	conn, err := p.DialFunc(p.HostDsn)

	if err != nil {
		p.Logger.Errorf("Failed to dial with the broker: %v", err)
		return FailedToDial
	}

	p.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		p.Logger.Errorf("Couldn't open a unique, concurrent server channel: %v", err)
		return SomethingWentWrong
	}

	_, err = ch.QueueDeclare(
		p.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		p.Logger.Errorf("Couldn't deaclare a queue: %v", err)
		return SomethingWentWrong
	}

	p.decoder = decoding.NewDecoder()
	p.ch = ch

	return nil
}

//Send ...
func (p *Producer) Send(message string) error {

	msg := &decoding.Request{Message: message}

	if p.decoder == nil || p.ch == nil {
		return FailedToDial
	}

	data, err := p.decoder.Decode(msg)
	if err != nil {
		return err
	}

	return p.ch.Publish(
		"",          // exchange
		p.QueueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "application/x-protobuf",
			Body:        data,
		})
}

//Finish ...
func (p *Producer) Finish() error {
	if p.conn == nil || p.ch == nil {
		return NilMemoryAddress
	}

	if err := p.ch.Close(); err != nil {
		p.Logger.Errorf("Couldn't close ch: %v", err)
		err = SomethingWentWrong
	}

	if err := p.conn.Close(); err != nil {
		p.Logger.Errorf("An error occurred while closing the connection: %s", err)
		err = SomethingWentWrong
	}
	return nil
}
