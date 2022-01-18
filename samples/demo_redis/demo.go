package demo_redis

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/zgo"
)

func Subscribe() {
  //ch,err := zgo.Redis.Subscribe(context.TODO(),"__keyevent@0__:expired") //订阅过期key的事件通知，用于定时执行任务
  ch, err := zgo.Redis.PSubscribe(context.TODO(), "my*")
  if err != nil {
    zgo.Log.Error(err)
    return
  }

  for {
    select {
    case msg := <-ch:
      fmt.Println(msg.Channel)
      fmt.Println(string(msg.Message))
    }
  }

}

func Xadd() {
  m := make(map[string]string)
  m["aaa"] = "aaa123"
  m["bbb"] = "bbb123"
  m["ccc"] = "ccc123"

  s, _ := zgo.Redis.XAdd(context.TODO(), "key-101", "19000000000010", m)
  //s, _ := zgo.Redis.XAdd(context.TODO(), "key-101", "*", m) //自动生成id
  fmt.Println(s, "--xadd")

  fmt.Println(zgo.Redis.XLen(context.TODO(), "key-101"))
}

func Xdel() {
  ids := []string{
    "1561371910929-0",
    "1561372099154-0",
  }
  i, _ := zgo.Redis.XDel(context.TODO(), "key-101", ids)
  fmt.Println(i, "xdel个数")

  fmt.Println(zgo.Redis.XLen(context.TODO(), "key-101"))

}

func Xrange() {
  strings, _ := zgo.Redis.XRange(context.TODO(), "key-101", "1561372594375-0", "1561372671389-0")
  //strings, _ := zgo.Redis.XRange(context.TODO(), "key-101", "-", "+")
  for _, v := range strings {
    for k, vv := range v {
      fmt.Println("key:", k)
      fmt.Println("value", vv)
    }
  }
  fmt.Println()
  fmt.Println()

  stringsRev, _ := zgo.Redis.XRevRange(context.TODO(), "key-101", "+", "-")
  for _, v := range stringsRev {
    for k, vv := range v {
      fmt.Println("key:", k)
      fmt.Println("value", vv)
    }
  }
}

func GroupCreate() {
  s, _ := zgo.Redis.XGroupCreate(context.TODO(), "key-101", "group-101", "$")
  //d, _ := zgo.Redis.XGroupDestroy(context.TODO(), "key-101", "group-101")
  fmt.Println(s, "--XGroupCreate--")
}

func Xack() {
  ids := []string{
    "19000000000010-0",
    "1561372594375-0",
  }
  s, _ := zgo.Redis.XAck(context.TODO(), "key-101", "group-101", ids)
  fmt.Println(s, "--XAck--")
}

func Read() {

  streams := []string{
    "lol",
  }
  streamReader, err := zgo.Redis.NewStreamReader(streams, "group-101", "101")
  if err != nil {
    zgo.Log.Error(err)
    return
  }

  if streamReader.Err() != nil {
    zgo.Log.Error(streamReader.Err())
    return
  }

  for {
    if _, entries, ok := streamReader.Next(); ok == true {
      if len(entries) > 0 {
        fmt.Println("-----", entries)
      }
    }
  }
}

func ReadNew() {
  var streamName = "key-101"
  var groupName = "group-102"

  zgo.Redis.XGroupCreate(context.TODO(), streamName, groupName, "0") //从0开始

  streams := []string{
    //streamName,
    //"lol",
  }
  streamReader, err := zgo.Redis.NewStreamReader(streams, groupName, "101")
  if err != nil {
    zgo.Log.Error(err)
    return
  }

  if streamReader.Err() != nil {
    zgo.Log.Error(streamReader.Err())
    return
  }

  for {
    if _, entries, ok := streamReader.Next(); ok == true {
      if len(entries) > 0 {
        fmt.Println(groupName, "===", entries)
      }
    }
  }
}
