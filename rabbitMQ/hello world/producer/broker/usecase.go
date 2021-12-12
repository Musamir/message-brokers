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
