package demo_rabbitmq

import (
	"context"
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"time"
)

/*
@Time : 2019-09-27 12:47
@Author : rubinus.chu
@File : producer
@project: origin
*/

const (
	label_bj = "mq_label_bj"
	label_sh = "mq_label_sh"
)

// ProducerByQueue 使用
func ProducerByQueue() {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	body := []byte(`{"name":"朱大仙儿","method":"使用队列名直接发送"}`)
	ch, err := zgo.MQ.ProducerByQueue(ctx, "queueName", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-ch

}

// Producer 经常使用
func Producer() {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	body := []byte(`{"name":"朱大仙儿"}`)
	ch, err := zgo.MQ.Producer(ctx, "exchangeName", "topic", "routingKey", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-ch

}

// ProducerMuLabel 多个label时
func ProducerMuLabel() {

	zgorabbitmq, err := zgo.MQ.New(label_bj)
	if err != nil {
		zgo.Log.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	body := []byte(`{"name":"朱大仙儿","label":"多个label--mq_label_bj"}`)
	ch, err := zgorabbitmq.Producer(ctx, "exchangeName", "topic", "routingKey", body)
	<-ch

}

// UseOriginConn使用原始连接
func UseOriginConn() {
	connChan, err := zgo.MQ.GetConnChan(label_bj) //获取原始连接chan
	if err != nil {
		zgo.Log.Error(err)
		return
	}

	if conn, ok := <-connChan; ok {
		c, err := conn.Channel()
		if err != nil {
			zgo.Log.Errorf("channel.open: %s", err)
		}

		err = c.ExchangeDeclare("logs", "topic", true, false, false, false, nil)
		if err != nil {
			zgo.Log.Errorf("exchange.declare: %v", err)
		}

		msg := zgo.RabbitmqPublishing{
			DeliveryMode: 2,
			Timestamp:    time.Now(),
			ContentType:  "text/plain",
			Body:         []byte("Go Go AMQP by zgo engine!"),
		}

		for {
			err = c.Publish("logs", "info", false, false, msg)
			if err != nil {
				zgo.Log.Errorf("basic.publish: %v", err)
			}
			fmt.Println(msg, "===")

			time.Sleep(2 * time.Second)
		}
	}
}
