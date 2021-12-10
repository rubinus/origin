package demo_kafka

import (
	"fmt"
	"github.com/gitcpu-io/origin/config"
	"github.com/gitcpu-io/zgo"
)

type chat struct {
	Topic   string
	GroupId string
}

func (c *chat) Consumer(label string) {
	//client, _ := zgo.Kafka.New(label)
	//client, _ := zgo.Kafka.New()

	consumer, _ := zgo.Kafka.Consumer(c.Topic, c.GroupId)
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

						fmt.Printf("==message===%d %s\n", msg.Offset, msg.Value)
						//zgo.Log.Info("==message===%d %s\n", msg.Offset, msg.Value)

						//todo something for u work

					}
				}(part)

			case msg, ok := <-consumer.Messages():
				if ok {
					fmt.Printf("==message===%d %s\n", msg.Offset, msg.Value)
					//zgo.Log.Infof("==message===%d %s\n", msg.Offset, msg.Value)

					//todo something for u work

				}

			}
		}
	}()

}

func Consumer() {
	c := chat{
		Topic:   config.Conf.Project,
		GroupId: config.Conf.Project,
	}
	go c.Consumer("")
}
