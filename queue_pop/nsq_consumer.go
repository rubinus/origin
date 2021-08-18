package queue_pop

import (
	"github.com/rubinus/zgo"
)

/*
@Time : 2019-06-11 10:12
@Author : rubinus.chu
@File : nsq_consumer
@project: origin
*/

type nsqer struct {
	Topic   string
	Channel string
}

func (c *nsqer) Consumer() {

	n, err := zgo.Nsq.New()
	if err != nil {
		panic(err)
	}
	n.Consumer(c.Topic, c.Channel, c.Deal)

}

func (c *nsqer) Deal(msg zgo.NsqMessage) error {

	//fmt.Println("接收到NSQ", msg.NSQDAddress, ",message:", string(msg.Body))

	//todo something for u work

	deal(msg.Body)

	return nil
}

func deal(body []byte) {

}

func NsqConsumer() {
	zgo.Log.Info("---------------启动消费Nsq---------------")

	go func() {
		c := nsqer{
			Topic:   "topic",
			Channel: "topic",
		}
		c.Consumer()
	}()

}
