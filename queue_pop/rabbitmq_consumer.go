package queue_pop

import (
  "github.com/gitcpu-io/zgo"
)

/*
@Time : 2019-06-11 10:12
@Author : rubinus.chu
@File : rabbitmq_consumer
@project: origin
*/

type Msg struct {
  ExchangeName string
  ExchangeType string
  RoutingKey   string
  QueueName    string
}

func (c *Msg) Consumer(label string) {

  consumer, err := zgo.MQ.Consumer(c.ExchangeName, c.ExchangeType, c.RoutingKey, c.QueueName)

  if err != nil {
    zgo.Log.Error(err)
    return
  }

  for val := range consumer {

    zgo.Log.Infof("ID: %s, 消息：%s, 类型:%s, 路由: %s, 交换机名: %s", val.MessageId, val.Body, val.Type, val.RoutingKey, val.Exchange)

    //todo something for u work

    dealRabbitMessage(val.Body)

  }

}

// 从main.go中调用
func RabbitmqConsumer() { //kafka topic 名字不能带有-
  //topic := fmt.Sprintf("%s_%s_%s", configs.MidPlatform, configs.Conf.Project, configs.Conf.KafkaTopics["noread"])
  zgo.Log.Info("---------------启动消费Rabbitmq---------------")
  c := Msg{
    ExchangeName: "exchangeName",
    ExchangeType: "topic",
    RoutingKey:   "routingKey",
    QueueName:    "queueName",
  }
  go c.Consumer("")
}

func dealRabbitMessage(body []byte) {

  //todo what you do

}
