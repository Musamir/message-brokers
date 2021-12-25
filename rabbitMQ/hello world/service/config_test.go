package service

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

//testConfig ...
var testConfig = Conf{
	Logger: struct {
		LogLevel string `yaml:"loglevel"`
	}{LogLevel: "debug"},
	RabbitMQ: struct {
		User  string `yaml:"user"`
		Host  string `yaml:"host"`
		Port  string `yaml:"port"`
		Queue string `yaml:"queue"`
	}{User: "user", Host: "rabbitmq", Port: "5672", Queue: "test queue"},
}

func TestInitConfig(t *testing.T) {
	testPath := "testConf.yaml"

	file, _ := os.Create("./testConf.yaml")
	defer func() {
		_ = file.Close()
		_ = os.Remove("./testConf.yaml")
	}()
	data := `
logger:
  loglevel: debug
rabbitmq:
  user: user
  host: rabbitmq
  port: '5672'
  queue: test queue`
	_, _ = file.WriteString(strings.TrimSpace(data))
	err := InitConfig(testPath)

	expected := testConfig
	assert.Equal(t, nil, err, "err must be nil")
	actual := *config
	assert.Equal(t, expected, actual)
}

//TestApplication_InitRabbitMQConfig ...
func TestApplication_InitRabbitMQConfig(t *testing.T) {
	_ = os.Setenv("RABBITMQ_PASS", "")
	actualURL, err := InitRabbitMQConfig(&testConfig)
	assert.NotEqual(t, nil, err, "err mustn't be nil")
	pass := "qwerty"
	_ = os.Setenv("RABBITMQ_PASS", pass)
	actualURL, _ = InitRabbitMQConfig(&testConfig)
	expectedURL := "amqp://user:qwerty@rabbitmq:5672"
	assert.Equal(t, expectedURL, actualURL, "not equal!!!")
}
