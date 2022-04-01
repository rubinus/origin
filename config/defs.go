package config

import (
  "errors"
  "fmt"
  "github.com/gitcpu-io/zgo"
  "io/ioutil"
  "path/filepath"
  "runtime"
  "strconv"
)

//从.json文件中加载配置项

var Conf *allConfig

//服务信息
type ServiceInfo struct {
  SvcName      string `json:"svc_name"`
  SvcHost      string `json:"svc_host"`
  SvcHttpPort  string `json:"svc_http_port"`
  SvcGrpcPort  string `json:"svc_grpc_port"`
  SvcEtcdAddress string `json:"svc_etcd_hosts"`
}
type Service struct {
  StartService         bool        `json:"start_service"`          //开启使用服务注册与服务发现
  StartServiceRegistry bool        `json:"start_service_registry"` //开启服务注册
  StartServiceDiscover bool        `json:"start_service_discover"` //开启服务发现
  ServiceInfo          ServiceInfo `json:"service_info"`
  OtherServices        []string    `json:"other_services"`
}

type allConfig struct {
  CPath         string `json:"cpath"`
  Env           string `json:"env"`
  Version       string `json:"version"`
  Project       string `json:"project"`
  EtcdAddress     string `json:"etcdAddress"`
  Loglevel      string `json:"loglevel"`
  IrisMod       bool   `json:"irisMod"`
  RpcHost       string `json:"rpcHost"`
  RpcPort       string `json:"rpcPort"`
  PprofPort     int    `json:"pprofPort"`
  HttpPort      int    `json:"httpPort"`
  UsePreAbsPath int    `json:"usePreAbsPath"`

  Service //内嵌服务结构体

  //demo host
  DemoHostForPayCanChangeAnyName string `json:"demo_host_for_pay_can_change_any_name"`
}

func InitConfig(cpath, env, project, etcdAddress, port, rpcPort string) {
  initConfig(cpath, env, project, etcdAddress, port, rpcPort)
}

func initConfig(cpath, env, project, etcdAddress, port, rpcPort string) {
  if env == "local" {
    _, f, _, ok := runtime.Caller(1)
    if !ok {
      panic(errors.New("Can not get current file info"))
    }
    cpath = filepath.Dir(f)
  }

  var cf = fmt.Sprintf("%s/%s.json", cpath, env)
  bf, err := ioutil.ReadFile(cf)
  if err != nil {
    fmt.Println("报错: ", cf)
    panic(err)
  }

  //使用zgo.Utils中的反序列化
  err = zgo.Utils.Unmarshal(bf, &Conf)
  if err != nil {
    fmt.Println("反序列化: ",cf)
    panic(err)
  }

  if Conf.CPath == "" && cpath != "" { //cpath没有，就用默认的项目所在路径
    Conf.CPath = cpath
  }
  fmt.Println("配置文件地址config:",cf)
  if project != "" {
    Conf.Project = project
  }
  if len(etcdAddress) != 0 {
    Conf.EtcdAddress = etcdAddress
  }
  if port != "" {
    portInt, err := strconv.Atoi(port)
    if err != nil {
      zgo.Log.Error(err)
    } else {
      Conf.HttpPort = portInt
    }

  }
  if rpcPort != "" {
    Conf.RpcPort = rpcPort
  }

  fmt.Printf("origin %s is started on the ... %s\n", Conf.Version, Conf.Env)
}

func WatchHttpConfigByService(ch chan string) {
  go func() {
    for value := range ch {
      lbRes, err := zgo.Service.LB(value) //变化的服务
      if err != nil {
        zgo.Log.Error(fmt.Sprintf("%s 服务取Http负载,", value), err)
        continue
      }

      switch value {
      case "timer.bffp": //自己做为客户端连接自己的服务端测试
        Conf.DemoHostForPayCanChangeAnyName = fmt.Sprintf("%s:%s", lbRes.SvcHost, lbRes.SvcHttpPort)
        //其它变量如果已经存在，可以在不改变原代码前提下，对config.Conf.***中的变量再次赋值

      case "other":
        //继续通过服务名，来再次初始化host port
      }

      zgo.Log.Warnf("监听到Http服务：%s,正在使用负载节点 Host: %s, http_port: %s", value, lbRes.SvcHost, lbRes.SvcHttpPort)

    }
  }()
}
