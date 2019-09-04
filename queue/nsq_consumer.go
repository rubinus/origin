package queue

import (
	"git.zhugefang.com/gobase/origin/config"
	"git.zhugefang.com/gocore/zgo"
)

/*
@Time : 2019-06-11 10:12
@Author : rubinus.chu
@File : nsq_consumer
@project: wechat
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

func deal(smg []byte) {

}

func NsqConsumer() {

	go func() {
		c := nsqer{
			Topic:   config.Conf.Project,
			Channel: config.Conf.Project,
		}
		c.Consumer()
	}()

}
