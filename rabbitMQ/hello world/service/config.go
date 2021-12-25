package service

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

//config ...
var config *Conf

//Conf ...
type Conf struct {
	Logger struct {
		LogLevel string `yaml:"loglevel"`
	} `yaml:"logger"`
	RabbitMQ struct {
		User  string `yaml:"user"`
		Host  string `yaml:"host"`
		Port  string `yaml:"port"`
		Queue string `yaml:"queue"`
	} `yaml:"rabbitmq"`
}

//InitConfig ...
func InitConfig(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	conf := &Conf{}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case conf.RabbitMQ.User == "":
		return fmt.Errorf("user field is empty")
	case conf.RabbitMQ.Host == "":
		return fmt.Errorf("host field is empty")
	case conf.RabbitMQ.Port == "":
		return fmt.Errorf("port field is empty")
	case conf.RabbitMQ.Queue == "":
		return fmt.Errorf("queue field is empty")
	}
	config = conf
	return nil
}

//InitLogger ...
func InitLogger(config *Conf) (Logger *zap.SugaredLogger, err error) {
	writer := zapcore.AddSync(os.Stdout)

	atom := zap.NewAtomicLevel()

	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.RFC3339TimeEncoder // TimeEncoderOfLayout(time.RFC3339)
	conf.EncodeLevel = zapcore.CapitalLevelEncoder
	zapEncoder := zapcore.NewConsoleEncoder(conf)
	core := zapcore.NewCore(zapEncoder, writer, atom)
	logger := zap.New(core, zap.AddCaller())

	err = atom.UnmarshalText([]byte(config.Logger.LogLevel))
	if err != nil {
		return nil, err
	}

	Logger = logger.Sugar()
	Logger.Info("Logger initialized!")
	return
}

//InitRabbitMQConfig ...
func InitRabbitMQConfig(conf *Conf) (RabbitMQURL string, err error) {
	pass := os.Getenv("RABBITMQ_PASS")
	if pass == "" {
		return RabbitMQURL, fmt.Errorf("%v(%s)", EnvironmentNotSet, "RABBITMQ_PASS")
	}
	//"amqp://guest:guest@localhost:5672/"
	return fmt.Sprintf("amqp://%s:%s@%s:%s", conf.RabbitMQ.User, pass, conf.RabbitMQ.Host, conf.RabbitMQ.Port), nil
}
