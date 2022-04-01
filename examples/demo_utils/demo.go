package demo_utils

import (
  "fmt"
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/zgo"
  "net"
)

var st = &struct {
  A string
}{
  A: "niubi",
}

func CallFun() {
  config.InitConfig("", "local", "", "", 0, 0)

  err := zgo.Engine(&zgo.Options{
    CPath:    config.Conf.CPath,
    Env:      config.Conf.Env,
    Project:  config.Conf.Project,
    Loglevel: config.Conf.Loglevel,
  })

  if err != nil {
    panic(err)
  }

  parseTime, _ := zgo.Utils.ParseTime("1573799617418")
  fmt.Println(parseTime.String())
  fmt.Println(parseTime.Minute(), parseTime.Second())

  fmt.Println(zgo.Utils.Snowflake(1).String())

  for i := 0; i < 2; i++ {
    fmt.Println(zgo.Utils.RandRangeInt(50, 150))

  }

  for i := 0; i < 2; i++ {
    fmt.Println(zgo.Utils.RandRangeInt(1000000, 10000000))

  }

  tm := make(map[string]interface{})
  tm["a"] = 2
  tm["b"] = "b"
  fmt.Println(tm)
  s, _ := zgo.Utils.ToString(tm)
  fmt.Println(s) //用于打印
  fmt.Println("==================")
  id, _ := zgo.Utils.IsChineseID("12222222")
  fmt.Println("IsChineseID:", id)

  //测试工具函数
  ip := zgo.Utils.GetIntranetIP()
  fmt.Println(ip)
  ts := zgo.Utils.GetTimestamp(13) //获取13位时间戳
  fmt.Printf("获取13位时间戳timestamp: %d\n", ts)

  m := map[string]string{}
  m["A"] = "123"
  m["B"] = "456"
  fmt.Println("GetUrlFormedMap:", zgo.Utils.GetUrlFormedMap(m), zgo.Utils.IPs())

  fmt.Println("==================")

  //测试log
  zgo.Log.Info(zgo.Utils.GetMD5Base64([]byte("lalalala")))
  zgo.Log.Info("333中国人")
  zgo.Log.Debug("44444")
  zgo.Log.Warn("55555")

  //test sjon
  str, err := zgo.Utils.Marshal(st)
  if err != nil {
    panic(err)
  }
  fmt.Print(string(str), err)

  //test json unmarshal
  st := &struct {
    B string
  }{}
  var str2 = `{"B":"456"}`
  err = zgo.Utils.Unmarshal([]byte(str), st)
  if err != nil {
    panic(err)
  }
  fmt.Println(st)
  fmt.Println(str2)
  fmt.Println("==================")

  timeString := zgo.Utils.FormatStringToStandTimeString("20190724151558")
  time, err := zgo.Utils.ParseTime(timeString)
  if err != nil {
    zgo.Log.Error(err)
  }
  pt := time
  short := zgo.Utils.FormatFromUnixTimeShort(pt.Unix())

  fmt.Println(short)

  byInterface, err := GetIPv4ByInterface("en0")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("byInterface:", byInterface)

  duration := zgo.Utils.NextDayDuration()
  fmt.Println("duration: ", duration)
}
func GetIPv4ByInterface(name string) (string, error) {
  var ips string

  iface, err := net.InterfaceByName(name)
  if err != nil {
    return "", err
  }

  addrs, err := iface.Addrs()
  if err != nil {
    return "", err
  }

  for _, a := range addrs {
    if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
      ips = ipnet.IP.String()
    }
  }
  return ips, nil
}
