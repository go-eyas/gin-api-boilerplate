package srv

// rabbit mq

import (
	"fmt"
	"github.com/go-eyas/toolkit/amqp"
)

type MQSrv struct {
	App *amqp.MQApp
}

var MQ = &MQSrv{}

func (mq *MQSrv) Init(conf *AmqpConfig) {
	app, err := amqp.NewApp(conf)
	if err != nil {
		panic(fmt.Errorf("mq init fail: %v", err))
	}
	mq.App = app

}
