package demo_nsq

import (
  "fmt"
	"github.com/gitcpu-io/origin/configs"
  "github.com/gitcpu-io/zgo"
)

type chat struct {
  Topic   string
  Channel string
}

func (c *chat) Consumer() {

  //n, err := zgo.Nsq.New("nsq_label_bj")
  n, err := zgo.Nsq.New()
  if err != nil {
    panic(err)
  }
  n.Consumer(c.Topic, c.Channel, c.Deal)

}

//处理消息
func (c *chat) Deal(msg zgo.NsqMessage) error {

  fmt.Println("接收到NSQ", msg.NSQDAddress, ",message:", string(msg.Body))

  //todo something for u work

  return nil
}

func Consumer() {

  go func() {
    c := chat{
      Topic:   configs.Conf.Project,
      Channel: configs.Conf.Project,
    }
    c.Consumer()
  }()
}
