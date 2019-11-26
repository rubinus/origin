package demo_kafka

import (
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {

	err := zgo.Engine(&zgo.Options{
		Env:     "local",
		Project: "1553240759",
		Kafka: []string{
			label_bj,
			label_sh,
		},
		Loglevel: "info",
	})

	if err != nil {
		panic(err)
	}

	//测试读取kafka数据，wait for sdk init connection
	time.Sleep(3 * time.Second)

	c := chat{
		Topic:   label_bj,
		GroupId: label_bj,
	}
	go c.Consumer(label_bj)
	c2 := chat{
		Topic:   label_sh,
		GroupId: label_sh,
	}
	go c2.Consumer(label_sh)

	for {
		select {
		case <-time.Tick(3 * time.Second):
			fmt.Println("一直在消费着")
		}
	}
}
