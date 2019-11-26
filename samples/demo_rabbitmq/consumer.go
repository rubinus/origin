package demo_rabbitmq

import (
	"fmt"
	"git.zhugefang.com/gocore/zgo"
)

/*
@Time : 2019-09-27 12:47
@Author : rubinus.chu
@File : consumer
@project: origin
*/

// ConsumerByQueue经常使用
func ConsumerByQueue() {

	ch, err := zgo.MQ.ConsumerByQueue("queueName")
	if err != nil {
		fmt.Println(err)
		return
	}
	for val := range ch {
		fmt.Printf("ID: %s, 消息：%s, 类型:%s, 路由: %s, 交换机名: %s\n", val.MessageId, val.Body, val.Type, val.RoutingKey, val.Exchange)
	}

}

// Consumer经常使用
func Consumer() {

	ch, err := zgo.MQ.Consumer("exchangeName", "topic", "routingKey", "queueName")
	if err != nil {
		fmt.Println(err)
		return
	}
	for val := range ch {
		fmt.Printf("ID: %s, 消息：%s, 类型:%s, 路由: %s, 交换机名: %s\n", val.MessageId, val.Body, val.Type, val.RoutingKey, val.Exchange)
	}

}

// ConsumerMuLabel消费多个label
func ConsumerMuLabel() {

	zgorabbitmq, err := zgo.MQ.New(label_bj)
	if err != nil {
		zgo.Log.Error(err)
		return
	}

	ch, err := zgorabbitmq.Consumer("exchangeName", "topic", "routingKey", "queueName")
	for val := range ch {
		fmt.Printf("ID: %s, 消息：%s, 类型:%s, 路由: %s, 交换机名: %s\n", val.MessageId, val.Body, val.Type, val.RoutingKey, val.Exchange)
	}

}
