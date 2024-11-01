package tool

import (
	"testing"

	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/stretchr/testify/assert"
)

var (
	rabbitMQConfig = config.NewConfig().GetConfig()

	rabbitMQTest = NewRabbitMQ(rabbitMQConfig)
)

func TestRabbitMQ_GetChannel(t *testing.T) {
	ch, err := rabbitMQTest.GetChannel()

	assert.Nil(t, err)
	assert.NotNil(t, ch)
}
