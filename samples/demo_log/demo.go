package demo_log

import (
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/origin/engine"
  "github.com/gitcpu-io/zgo"
)

const (
  debug = iota
  info
  warn
  error
)

func init() {
  config.InitConfig("","dev", "1553240759", "", "", "")
}

func Call() {
  err := engine.Run()
  if err != nil {
    panic(err)
  }

  zgo.Log.Debug(debug)
  zgo.Log.Info(info)
  zgo.Log.Warn(warn)
  zgo.Log.Error(error)

  //l.Info(222)

  //fmt.Println(debug)
  //fmt.Println(info)
  //fmt.Println(warn)
  //fmt.Println(error)
}
