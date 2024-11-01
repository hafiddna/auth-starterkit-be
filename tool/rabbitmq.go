package tool

import (
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQRPCRequest struct {
	Route string      `json:"route"`
	Param interface{} `json:"param"`
	Data  interface{} `json:"data"`
}

// RabbitMQRPCResponse
/* Status Codes
OK 0
CANCELLED 1
UNKNOWN 2
INVALID_ARGUMENT 3
DEADLINE_EXCEEDED 4
NOT_FOUND 5
ALREADY_EXISTS 6
PERMISSION_DENIED 7
RESOURCE_EXHAUSTED 8
FAILED_PRECONDITION 9
ABORTED 10
OUT_OF_RANGE 11
UNIMPLEMENTED 12
INTERNAL 13
UNAVAILABLE 14
DATA_LOSS 15
UNAUTHENTICATED 16
*/
type RabbitMQRPCResponse struct {
	StatusCode int         `json:"status_code"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type RabbitMQ interface {
	GetChannel() (*amqp.Channel, error)
}

type rabbitMQ struct {
	config config.CfgStruct
}

func NewRabbitMQ(config config.CfgStruct) RabbitMQ {
	return &rabbitMQ{
		config: config,
	}
}

func (r *rabbitMQ) GetChannel() (*amqp.Channel, error) {
	uri := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		r.config.App.RabbitMQ.Username,
		r.config.App.RabbitMQ.Password,
		r.config.App.RabbitMQ.Host,
		r.config.App.RabbitMQ.Port,
	)
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	return ch, nil
}
