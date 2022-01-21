package demo_map

import (
  "fmt"
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/zgo"
  "time"
)

/*
@Time : 2019-03-15 11:49
@Author : rubinus.chu
@File : demo
@project: origin
*/

func Demo() {
  config.InitConfig("local", "", "", "", "")

  err := zgo.Engine(&zgo.Options{
    CPath: config.Conf.CPath,
    Env:      config.Conf.Env,
    Project:  config.Conf.Project,
    Loglevel: config.Conf.Loglevel,
  })
  if err != nil {
    panic(err)
  }

  //zgo中的map是并发安全的
  sm := zgo.Map.New()

  for i := 0; i < 100; i++ {
    go func(i int) {
      sm.Set(i, i)
    }(i)

  }
  for i := 0; i < 100; i++ {
    go func(i int) {
      fmt.Println(sm.Get(i))
    }(i)
  }

  time.Sleep(1 * time.Second)
  for v := range sm.Range() {
    fmt.Println(v.Key, "range map == ", v.Val)
  }

  //不安全的map
  //sm := make(map[int]int)
  //for i := 0; i < 100; i++ {
  //	go func(i int) {
  //		sm[i] = i
  //	}(i)
  //}
  //for i := 0; i < 100; i++ {
  //	go func(i int) {
  //		fmt.Println(sm[i])
  //	}(i)
  //}

}
