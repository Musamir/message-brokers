package producer

//Producer ...
type Producer interface {
	Send(message string) error
	Finish() error
}
