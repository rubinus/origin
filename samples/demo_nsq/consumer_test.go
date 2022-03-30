package demo_nsq

import (
  "fmt"
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/zgo"
  "net/url"
  "testing"
  "time"
)

func TestConsumer(t *testing.T) {
  config.InitConfig("local", "", "", "", "")

  err := zgo.Engine(&zgo.Options{
    Project: config.Conf.Project,
    Env:     config.Conf.Env,
    Nsq: []string{
      label_bj,
      label_sh,
    },
  })
  if err != nil {
    panic(err)
  }

  c := chat{
    Topic:   "origin",
    Channel: label_bj,
  }
  c.Consumer()
  c2 := chat{
    Topic:   label_sh,
    Channel: label_sh,
  }
  c2.Consumer()

  interval := time.Duration(3) * time.Second
  ticker := time.NewTicker(interval)
  for val := range ticker.C {
    fmt.Println("一直在消费着",val)
    ticker.Reset(interval)
  }
}

func TestConsumer2(t *testing.T) {

  values, _ := url.ParseQuery("http://baidu.com?state=send_back&q=q101&ext=kuozhanziduan")
  q := values.Get("q")     //队列名
  ext := values.Get("ext") //扩展字段
  s := values.Get("s")     //h5、open
  fmt.Println(q)
  fmt.Println(ext)
  fmt.Println(s)

  fmt.Println(url.QueryEscape("http://baidu.com"))

  err := zgo.Engine(&zgo.Options{
    Project: "1559734565",
    Env:     "dev",
  })
  if err != nil {
    panic(err)
  }

  time.Sleep(1 * time.Second)

  c := chat{
    Topic:   "queue-101",
    Channel: "queue-101",
  }
  c.Consumer()

  for {
    if false {
      fmt.Println(111)
    }
    select {
    case <-time.NewTicker(time.Duration(1 * time.Minute)).C:
      fmt.Println("一直在消费着")
    default:

    }
  }
}
