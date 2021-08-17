package queue

import (
	"github.com/rubinus/zgo"
)

/*
@Time : 2019-06-11 10:12
@Author : rubinus.chu
@File : kafka_consumer
@project: origin
*/

type noRead struct {
	Topic   string
	GroupId string
}

func (c *noRead) Consumer(label string) {
	//client, _ := zgo.Kafka.New(label)
	//client, _ := zgo.Kafka.New()

	consumer, err := zgo.Kafka.Consumer(c.Topic, c.GroupId)

	if err != nil {
		zgo.Log.Error(err)
		return
	}
	go func() {
		for {
			select {
			case part, ok := <-consumer.Partitions():

				if !ok {
					return
				}
				// start a separate goroutine to consume messages
				go func(pc zgo.PartitionConsumer) {
					for msg := range pc.Messages() {

						zgo.Log.Infof("==partition message===%s %d %s", msg.Topic, msg.Offset, msg.Value)

						//todo something for u work

						dealMessage(msg.Value)

					}
				}(part)

			case msg, ok := <-consumer.Messages():
				if ok {
					zgo.Log.Infof("----message----%s %d %s", msg.Topic, msg.Offset, msg.Value)

					//todo something for u work

					dealMessage(msg.Value)

				}

			}
		}
	}()

}

func KafkaConsumer() { //kafka topic 名字不能带有-
	//topic := fmt.Sprintf("%s_%s_%s", config.MidPlatform, config.Conf.Project, config.Conf.KafkaTopics["noread"])
	zgo.Log.Info("---------------启动消费Kafka---------------")

	topic := ""
	c := noRead{
		Topic:   topic,
		GroupId: topic,
	}
	go c.Consumer("")
}

func dealMessage(body []byte) {

}
