package broker

import "github.com/streadway/amqp"

//AmqpDialWrapper ...
func AmqpDialWrapper(url string) (AmqpConnection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return AmqpConnectionWrapper{conn}, nil
}

//AmqpConnectionWrapper ...
type AmqpConnectionWrapper struct {
	conn *amqp.Connection
}

//Channel ...
func (w AmqpConnectionWrapper) Channel() (AmqpChannel, error) {
	return w.conn.Channel()
}

//Close ...
func (w AmqpConnectionWrapper) Close() error {
	return w.conn.Close()
}

//
////AmqpConnectionImpl ...
//type AmqpConnectionImpl struct {
//	conn *amqp.Connection
//}
//
////Channel ...
//func (a *AmqpConnectionImpl) Channel() (AmqpChannel, error) {
//	return a.conn.Channel()
//}
//
////Close ...
//func (a *AmqpConnectionImpl) Close() error {
//	return a.conn.Close()
//}
//
////AmqpChannelImpl ...
//type AmqpChannelImpl struct {
//	ch AmqpChannel
//}
//
////Close ...
//func (a *AmqpChannelImpl) Close() error {
//	return a.ch.Close()
//}
//
////QueueDeclare ...
//func (a *AmqpChannelImpl) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
//	return a.ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
//}
//
////Consume ...
//func (a *AmqpChannelImpl) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
//	return a.ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
//}
