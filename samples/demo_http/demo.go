package demo_http

import (
  "bytes"
  "encoding/base64"
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

var (
  url = "http://example.com"
)

func DemoGet() {
  config.InitConfig("","local", "", "", 0, 0)

  err := zgo.Engine(&zgo.Options{
    Env:      config.Conf.Env,
    Project:  config.Conf.Project,
    Loglevel: config.Conf.Loglevel,
  })
  if err != nil {
    panic(err)
  }
  time.Sleep(1 * time.Second)

  bytes, err := zgo.Http.Get(url)
  if err != nil {
    zgo.Log.Error(err)
  }
  fmt.Printf("%s", bytes)

}

func DemoPostJson() {
  config.InitConfig("","local", "", "", 0, 0)

  err := zgo.Engine(&zgo.Options{
    Env:      config.Conf.Env,
    Project:  config.Conf.Project,
    Loglevel: config.Conf.Loglevel,
  })
  if err != nil {
    panic(err)
  }
  time.Sleep(1 * time.Second)

  buf := new(bytes.Buffer)

  _, err = zgo.File.Get("cp.jpeg", buf)
  if err != nil {
    zgo.Log.Debug(err)
  }
  encodeToString := base64.StdEncoding.EncodeToString(buf.Bytes())

  data := fmt.Sprintf("{\"image\":\"%s\"}", encodeToString)
  b1, err := zgo.Http.PostForm("https://aip.baidubce.com/rest/2.0/ocr/v1/lottery?access_token=24.96f79401c2c76f1cd3ab5ad78f91bee9.2592000.1574395795.282335-17598500", []byte(data))

  if err != nil {
    zgo.Log.Error(err)
  }
  fmt.Printf("%s\n", b1)

  //byte, err = zgo.Http.PostForm("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=9tQWYFwI8Ai1LfpvO3ounpf6&client_secret=7FiC4vGd6wrFwZafGk0HjLU7LjeGO1KI", []byte(jsonStr))
  //if err != nil {
  //	zgo.Log.Error(err)
  //}
  //fmt.Printf("%s\n", byte)

  //byte, err := zgo.Http.PostJson("http://localhost/demopost", []byte(jsonStr))
  //if err != nil {
  //	zgo.Log.Error(err)
  //}
  //fmt.Printf("%s", byte)

}
