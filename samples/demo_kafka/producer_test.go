package demo_kafka

import (
  "fmt"
  "github.com/gitcpu-io/zgo"
  "testing"
  "time"
)

const (
  label_bj = "kafka_label_bj"
  label_sh = "kafka_label_sh"
)

var project = "origin"

func TestProducer(t *testing.T) {
  err := zgo.Engine(&zgo.Options{
    Env:     "dev",
    Project: "1553240759",
    Kafka: []string{
      label_bj,
      label_sh,
    },
  }) //测试时表示使用kafka，在origin中使用一次

  if err != nil {
    panic(err)
  }

  //测试读取kafka数据，wait for sdk init connection
  time.Sleep(2 * time.Second)

  clientBj, err := zgo.Kafka.New()
  if err != nil {
    panic(err)
  }
  clientSh, err := zgo.Kafka.New()
  if err != nil {
    panic(err)
  }
  //map[string][]map[string]string

  var replyChan = make(chan int)
  var countChan = make(chan int)
  l := 100 //暴力测试50000个消息，时间10秒，本本的并发每秒5000

  count := []int{}
  total := []int{}
  stime := time.Now()

  for i := 0; i < l; i++ {
    go func(i int) {
      countChan <- i //统计开出去的goroutine
      if i%2 == 0 {
        ch := Producer(project, clientBj, i, false)
        reply := <-ch
        replyChan <- reply

      } else {
        ch := Producer(project, clientSh, i, false)
        reply := <-ch
        replyChan <- reply
      }
    }(i)
  }

  go func() {
    for v := range replyChan {
      if v != 10001 { //10001表示超时
        count = append(count, v) //成功数
      }
    }
  }()

  go func() {
    for v := range countChan { //总共的goroutine
      total = append(total, v)
    }
  }()

  for _, v := range count {
    if v != 1 {
      fmt.Println("有不成功的")
    }
  }

  for {
    if len(count) == l {
      var timeLen = time.Since(stime)
      fmt.Printf("总消耗时间：%s, 成功：%d, 总共开出来的goroutine：%d\n", timeLen, len(count), len(total))
      break
    }

    select {
    case <-time.Tick(time.Duration(1000 * time.Millisecond)):
      fmt.Println("处理进度每1000毫秒", len(count))
    default:

    }
  }
  time.Sleep(2 * time.Second)
}
