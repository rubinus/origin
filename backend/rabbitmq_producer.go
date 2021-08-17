package backend

import (
	"context"
	"github.com/rubinus/zgo"
	"time"
)

/*
@Time : 2019-06-14 12:15
@Author : rubinus.chu
@File : rabbitmq_producer
@project: origin
*/

func RabbitmqProducer(exchangeName, exchangeType, routingKey string, body interface{}) chan uint8 {

	out := make(chan uint8, 1)

	bytes, err := zgo.Utils.Marshal(body)
	if err != nil {
		zgo.Log.Error(err)
		out <- 0
		return out
	}
	//zgo.Log.Info(string(bytes))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	out, err = zgo.MQ.Producer(ctx, exchangeName, exchangeType, routingKey, bytes)

	if err != nil {
		zgo.Log.Error(err)
		out <- 0
		return out
	}
	return out
}
