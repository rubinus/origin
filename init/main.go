package main

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/gitcpu-io/zgo"
  "github.com/gitcpu-io/zgo/config"
  "go.etcd.io/etcd/client/v3"
  "os"
  "time"
)

var cli *clientv3.Client

func CreateClient() (*clientv3.Client, error) {
  return clientv3.New(clientv3.Config{
    Endpoints: []string{
      "127.0.0.1:3379",
      //"测试机dev环境:2379",
    },
    DialTimeout: 10 * time.Second,
  })
}

func main() {
  c, err := CreateClient()
  if err != nil {
    return
  }
  cli = c

  getwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  err = zgo.Engine(&zgo.Options{
    CPath: fmt.Sprintf("%s/%s",getwd,"init"),
    Env:     "local",
    Project: "origin",
  })
  if err != nil {
    panic(err)
  }

  for _, v := range config.Conf.Nsq {
    k := v.Key
    value := v.Values
    key := "zgo/project/origin/nsq/" + k
    val, _ := json.Marshal(value)
    res, err := cli.KV.Put(context.TODO(), key, string(val), clientv3.WithPrevKV())
    if err != nil {
      panic(err)
    }
    fmt.Println(res)
  }

  for _, v := range config.Conf.Mongo {
    k := v.Key
    value := v.Values
    key := "zgo/project/origin/mongo/" + k
    val, _ := json.Marshal(value)
    _, err := cli.KV.Put(context.TODO(), key, string(val))
    if err != nil {
      fmt.Println(err)
    }
  }

  for _, v := range config.Conf.Es {
    k := v.Key
    value := v.Values
    key := "zgo/project/origin/es/" + k
    val, _ := json.Marshal(value)
    _, err := cli.KV.Put(context.TODO(), key, string(val))
    if err != nil {
      fmt.Println(err)
    }
  }
  for _, v := range config.Conf.Mysql {
    k := v.Key
    value := v.Values
    key := "zgo/project/origin/mysql/" + k
    val, _ := json.Marshal(value)
    _, err := cli.KV.Put(context.TODO(), key, string(val))
    if err != nil {
      fmt.Println(err)
    }
  }
  for _, v := range config.Conf.Etcd {
    k := v.Key
    value := v.Values
    key := "zgo/project/origin/etcd/" + k
    val, _ := json.Marshal(value)
    _, err := cli.KV.Put(context.TODO(), key, string(val))
    if err != nil {
      fmt.Println(err)
    }
  }

  for _, v := range config.Conf.Kafka {
    k := v.Key
    value := v.Values
    key := "zgo/project/origin/kafka/" + k
    val, _ := json.Marshal(value)
    _, err := cli.KV.Put(context.TODO(), key, string(val))
    if err != nil {
      fmt.Println(err)
    }
  }

  for _, v := range config.Conf.Redis {
    k := v.Key
    value := v.Values
    key := "zgo/project/origin/redis/" + k
    val, _ := json.Marshal(value)
    _, err := cli.KV.Put(context.TODO(), key, string(val))
    if err != nil {
      fmt.Println(err)
    }
  }

  for _, v := range config.Conf.Postgres {
    k := v.Key
    value := v.Values
    key := "zgo/project/origin/postgres/" + k
    val, _ := json.Marshal(value)
    _, err := cli.KV.Put(context.TODO(), key, string(val))
    if err != nil {
      fmt.Println(err)
    }
  }

  key := "zgo/project/origin/cache"
  val, _ := json.Marshal(config.Conf.Cache)
  _, err = cli.KV.Put(context.TODO(), key, string(val))
  if err != nil {
    fmt.Println(err)
  }

  key_log := "zgo/project/origin/log"
  val_log, _ := json.Marshal(config.Conf.Log)
  //val_log := "{\"c\": \"日志存储\",\"start\": 1,\"dbType\": \"nsq\",\"label\":\"nsq_label_bj\"}"
  res, err := cli.KV.Put(context.TODO(), key_log, string(val_log))
  fmt.Println(res, err)

  fmt.Println("all config to etcd done")

}
