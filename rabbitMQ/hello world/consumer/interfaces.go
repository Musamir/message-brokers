package consumer

//Consumer ...
type Consumer interface {
	Receive() error
}
