package service

import (
	"go.uber.org/zap"
	"hello-world/consumer"
	consumerImp "hello-world/consumer/implementation"
	"hello-world/producer"
	producerImp "hello-world/producer/implementation"
	"sync"
)

//Application ...
type Application struct {
	Logger   *zap.SugaredLogger
	producer producer.Producer
	consumer consumer.Consumer
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

	RabbitMQURL, err := InitRabbitMQConfig(config)
	if err != nil {
		logger.Errorf("%v", err)
		return nil, err
	}

	logger.Infof("RabbitMQURL=[%s]", RabbitMQURL)

	logger.Info("Starting init producer")

	p, err := producerImp.NewProducer(logger, RabbitMQURL, config.RabbitMQ.Queue)
	if err != nil {
		logger.Errorf("%v", err)
		return nil, FailedInit
	}

	logger.Info("producer was initialized")

	logger.Info("Starting init consumer")

	c, err := consumerImp.NewConsumer(logger, RabbitMQURL, config.RabbitMQ.Queue)
	if err != nil {
		logger.Errorf("%v", err)
		return nil, FailedInit
	}
	logger.Info("consumer was initialized")

	return &Application{
		Logger:   logger,
		producer: p,
		consumer: c,
	}, nil
}

//testData ...
var testData = []string{
	"Hello World!!",
	"Hope you've got these messages",
}

//Run ...
func (a *Application) Run() error {
	a.Logger.Info("Starting the application ...\n")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		a.Logger.Info("Starting to send messages")
		var err error
		for _, msg := range testData {
			if err = a.producer.Send(msg); err != nil {
				a.Logger.Errorf("Couldn't send the message %s: %v", msg, err)
			} else {
				a.Logger.Infof("The message [%s] successfuly was sent", msg)
			}
		}
		a.Logger.Info("Completing sending messages")
		if err = a.producer.Finish(); err != nil {
			a.Logger.Errorf("%v", err)
		} else {
			a.Logger.Info("Completion has been completed")
		}
	}()

	go func() {
		defer wg.Done()
		a.Logger.Info("Starting receiving messages ")
		if err := a.consumer.Receive(); err != nil {
			return
		}
	}()

	wg.Wait()

	a.Logger.Info("Stopping the program ...")

	if err := a.Logger.Sync(); err != nil {
		a.Logger.Errorf("error to flush buffered log%v", err)
		return err
	}
	return nil
}
