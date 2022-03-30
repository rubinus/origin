package demo_limiter

import (
  "fmt"
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/zgo"
  "time"
)

/*
@Time : 2019-04-29 14:34
@Author : rubinus.chu
@File : demo
@project: origin
*/

func CallLimiter() {
  config.InitConfig("local", "", "", "", "")

  err := zgo.Engine(&zgo.Options{
    Env:      config.Conf.Env,
    Project:  config.Conf.Project,
    Loglevel: config.Conf.Loglevel,
  })

  if err != nil {
    panic(err)
  }

  //###############################################################开始使用limiter
  b := zgo.Limiter.NewBucket(3) //生成3个长度的bucket

  //用法：
  //b.Get() //如果返回>0，表示可用。
  //b.Release() //搞定业务后，释放，让其它请求用

  go func() {
    for {
      select {
      case <-time.NewTicker(200 * time.Millisecond).C:
        b.Resize(10, 0)
        fmt.Println("============================重置长度10", b.Cap(), b.Len())
      default:
        time.Sleep(1*time.Millisecond)
      }
    }
  }()

  go func() {
    for {
      select {
      case <-time.NewTicker(100 * time.Millisecond).C:
        b.Resize(6, 0)
        fmt.Println("------------------------------重置长度6", b.Cap(), b.Len())
      default:
        time.Sleep(1*time.Millisecond)
      }
    }
  }()

  go func() {
    for {
      select {
      case <-time.NewTicker(100 * time.Millisecond).C:
        b.Release(1) //每1秒释放一次token
        fmt.Println("容量:", b.Cap(), "长度:", b.Len())
      default:
        time.Sleep(1*time.Millisecond)
      }
    }
  }()

  for {
    if false {
      fmt.Println(111)
    }
    select {
    case <-time.NewTicker(100 * time.Millisecond).C:
      token := b.Get(1)
      if token > 0 {
        fmt.Println(b.Cap(), b.Len(), "取到token，可以继续业务。b.GetToken()==", token)
      } else {
        fmt.Println(b.Len(), "no token，can't continue。b.GetToken()==", token)
      }
    default:

    }
  }
}

func CallLimiter2() {
  config.InitConfig("local", "", "", "", "")

  err := zgo.Engine(&zgo.Options{
    Env:      config.Conf.Env,
    Project:  config.Conf.Project,
    Loglevel: config.Conf.Loglevel,
  })

  if err != nil {
    panic(err)
  }

  //###############################################################开始使用limiter
  b := zgo.Limiter.NewBucket(300) //生成3个长度的bucket

  //用法：
  //b.Get() //如果返回>0，表示可用。
  //b.Release() //搞定业务后，释放，让其它请求用

  for i := 0; i < 30; i++ {
    fmt.Println(i, "连续取30次", b.Get(10)) //连续取3次
  }

  b.Resize(1000, 0) //扩充到10

  for i := 0; i < 70; i++ {
    fmt.Println(i, "扩充到1000,还可以再连续取70次", b.Get(10)) //还可以再连续取7次
  }

  b.Resize(600, 0) //缩小到6
  go func() {
    b.Resize(300, 0)
  }()

  for i := 0; i < 100; i++ {
    fmt.Println(i, "只有60次可以释放", b.Release(10)) //还可以再连续取10次
  }

  for i := 0; i < 100; i++ {
    fmt.Println(i, "最新的只能取到60次", b.Get(10)) //连续取10次

  }

}

func CallLimiter3() {
  config.InitConfig("local", "", "", "", "")

  err := zgo.Engine(&zgo.Options{
    Env:      config.Conf.Env,
    Project:  config.Conf.Project,
    Loglevel: config.Conf.Loglevel,
  })

  if err != nil {
    panic(err)
  }

  //###############################################################开始使用limiter
  b := zgo.Limiter.NewBucket(3) //生成3个长度的bucket

  //用法：
  //b.Get() //如果返回>0，表示可用。
  //b.Release() //搞定业务后，释放，让其它请求用

  go func() {
    for i := 0; i < 10000; i++ {
      time.Sleep(1000 * time.Millisecond)

      token := b.Get(1)
      if token > 0 {
        fmt.Println(b.Cap(), b.Len(), "取到token，可以继续业务。b.GetToken()==", token)
      } else {
        fmt.Println(b.Len(), "no token，can't continue。b.GetToken()==", token)
      }
    }
  }()

  go func() {
    for i := 0; i < 10000; i++ {
      time.Sleep(1000 * time.Millisecond)
      fmt.Println(b.Release(1)) //每1秒释放一次token
    }
  }()

  for {
    if false {
      fmt.Println(111)
    }
    select {
    case <-time.NewTicker(500 * time.Millisecond).C:
      b.Resize(10, 0)
      fmt.Println("============================重置长度10", b.Cap(), b.Len())
    default:

    }
  }

}
