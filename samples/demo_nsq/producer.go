package demo_nsq

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/zgo/zgonsq"
  "time"
)

func Producer(label string, nsqClient zgonsq.Nsqer, i int, b bool) chan int {

  //还需要一个上下文用来控制开出去的goroutine是否超时
  ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  //输入参数：上下文ctx，nsqClientChan里面是client的连接，args具体的查询操作参数

  body := []byte(fmt.Sprintf("msg--%s--%d", label, i))
  bodyMulti := []byte(fmt.Sprintf("msg-multi-%s--%d", label, i))
  var rch chan uint8
  var err error
  if b == true { //一次发送多条
    bodyMutil := [][]byte{
      body,
      bodyMulti,
    }
    rch, err = nsqClient.ProducerMulti(ctx, label, bodyMutil)

  } else {
    rch, err = nsqClient.Producer(ctx, label, body)

  }
  if err != nil {
    panic(err)
  }

  out := make(chan int, 1)
  select {
  case <-ctx.Done():
    fmt.Println(label, "超时")
    out <- 10001
    return out
  case b := <-rch:
    if b == 1 {
      out <- 1

    } else {
      out <- 10001
    }
  }

  return out

}
