package service

import (
	"go.uber.org/zap"
	consumerImp "hello-world/consumer/implementation"
	producerImp "hello-world/producer/implementation"
	"os"
	"os/signal"
)

//Application ...
type Application struct {
	Logger *zap.SugaredLogger
}

//NewApplication ...
func NewApplication(path string) (*Application, error) {
	err := InitConfig(path)
	if err != nil {
		return nil, err
	}
	logger, err := InitLogger(config)
	if err != nil {
		return nil, err
	}
	return &Application{
		Logger: logger,
	}, nil
}

//Run ...
func (a *Application) Run() error {
	a.Logger.Info("Starting the application ...\n")
	RabbitMQURL, err := a.InitRabbitMQConfig(config)
	if err != nil {
		return err
	}

	a.Logger.Infof("RabbitMQURL=[%s]", RabbitMQURL)

	a.Logger.Info("Starting init producer")

	producer, err := producerImp.NewProducer(a.Logger, RabbitMQURL, config.RabbitMQ.Queue)
	if err != nil {
		a.Logger.Errorf("%v", err)
		return FailedInit
	}

	a.Logger.Info("producer was initialized")

	go func() {
		a.Logger.Info("Starting to send messages")
		//testData ...
		var testData = []string{
			"Hello World!!",
			"Hope you've got these messages",
		}
		var err error
		for _, msg := range testData {
			if err = producer.Send(msg); err != nil {
				a.Logger.Errorf("Couldn't send the message %s: %v", msg, err)
			} else {
				a.Logger.Infof("The message [%s] successfuly was sent", msg)
			}
		}
		a.Logger.Info("Completing sending messages")
		if err = producer.Finish(); err != nil {
			a.Logger.Errorf("%v", err)
		} else {
			a.Logger.Info("Completion has been completed")
		}
	}()

	a.Logger.Info("Starting init consumer")

	consumer, err := consumerImp.NewConsumer(a.Logger, RabbitMQURL, config.RabbitMQ.Queue)
	if err != nil {
		a.Logger.Errorf("%v", err)
		return FailedInit
	}

	go func() {
		a.Logger.Info("Starting receiving messages ")
		if err = consumer.Receive(); err != nil {
			return
		}
	}()

	if err != nil {
		return FailedInit
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.Logger.Info("Stopping the program ...")

	if err := a.Logger.Sync(); err != nil {
		a.Logger.Errorf("error to flush buffered log%v", err)
		return err
	}
	return nil
}
