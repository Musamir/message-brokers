//go:generate mockery --all --keeptree
package broker

import "github.com/streadway/amqp"

//AmqpChannel ...
type AmqpChannel interface {
	Close() error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
}

//AmqpConnection ...
type AmqpConnection interface {
	Channel() (AmqpChannel, error)
	Close() error
}

//AmqpDial ...
type AmqpDial func(url string) (AmqpConnection, error)
