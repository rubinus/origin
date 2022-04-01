package demo_rabbitmq

import (
  "github.com/gitcpu-io/origin/queue_push"
  "github.com/gitcpu-io/zgo"
  "testing"
  "time"
)

/*
@Time : 2019-09-28 12:48
@Author : rubinus.chu
@File : producer_test
@project: origin
*/

//var project = "1553240759"
var project = "1553240759"

type Msg struct {
  Name       string `json:"name"`
  Age        int    `json:"age"`
  CreateTime int64  `json:"create_time"`
}

func TestProducer(t *testing.T) {
  err := zgo.Engine(&zgo.Options{
    Env:     "dev",
    Project: project,

    //Rabbitmq: []string{
    //	label_bj,
    //	label_sh,
    //},
  }) //测试时表示使用rabbitmq，在origin中使用一次

  if err != nil {
    panic(err)
  }

  time.Sleep(3 * time.Second)

  //生产数据到rabbitmq
  go func() {
    for {
      time.Sleep(3 * time.Second)
      Producer() //第一：简单测试每3秒发送一条
      //ProducerMuLabel() //第三发送多个label时

      //***********第二：测试调用封装好的rabbit mq 生产数据
      msg := &Msg{
        Name:       "朱大仙儿",
        Age:        30,
        CreateTime: zgo.Utils.GetTimestamp(13), //获取13位unix时间戳
      }
      for i := 0; i < 100; i++ {
        queue_push.RabbitmqProducer("exchangeName", "topic", "routingKey", msg)
      }
      //***********测试调用封装
    }
  }()

  //************第二：测试封装好的rabbit mq 消费数据
  //queue_pop.RabbitmqConsumer()
  //************测试封装好的rabbit mq 消费数据

  go Consumer() //第一：简单测试
  //go ConsumerMuLabel() //第三：多个label

  for {
    time.Sleep(3 * time.Second)
  }

  //use 原始连接
  //UseOriginConn()

}

func TestProducer2(t *testing.T) {
  err := zgo.Engine(&zgo.Options{
    Env:     "dev",
    Project: project,

    Rabbitmq: []string{
      label_bj,
      label_sh,
    },
  }) //测试时表示使用rabbitmq，在origin中使用一次

  if err != nil {
    panic(err)
  }

  time.Sleep(2 * time.Second)

  //生产数据到rabbitmq
  go func() {
    for {
      time.Sleep(3 * time.Second)
      ProducerByQueue() //使用队列

    }
  }()
  go ConsumerByQueue() //使用队列

  for {
    time.Sleep(3 * time.Second)
  }

}
