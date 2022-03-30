package demo_limiter

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/zgo"
)

/*
@Time : 2022-01-19 14:23
@Author : rubinus.chu
@File : demo
@project: origin
*/

func CallK8s() {
  config.InitConfig("local", "", "", "", "")

  err := zgo.Engine(&zgo.Options{
    CPath:    config.Conf.CPath,
    Env:      config.Conf.Env,
    Project:  config.Conf.Project,
    Loglevel: config.Conf.Loglevel,
  })

  if err != nil {
    panic(err)
  }

  //###############################################################开始使用k8s
  //1.通过配置选项Func
  cof := zgo.K8s.ConfigOption().GetFunc()

  //2.为参数赋值
  kco, err := zgo.K8s.ConfigOption().Build(
    cof.WithMasterUrl("https://kubernetes.docker.internal:6443"),
    cof.WithKubeConfig("/Users/rubinus/.kube/config"),
  )
  if err != nil {
    panic(err)
  }
  //2-1.为另一个赋值
  kco_1, err := zgo.K8s.ConfigOption().Build(
   cof.WithMasterUrl("https://install-prow-dns-809743c1.hcp.eastasia.azmk8s.io:443"),
   cof.WithKubeConfig("/Users/rubinus/app/kubeconfig/config-prow"),
  )

  if err != nil {
    zgo.Log.Infof("----%+v", err)
    return
  }

  //3.调用生成config
  config, err := zgo.K8s.Builder().BuildConfig(kco)
  if err != nil {
    return
  }

  //3-1.调用生成另外一个config
  config_1, err := zgo.K8s.Builder().BuildConfig(kco_1)
  if err != nil {
   return
  }

  //4. 打印config
  fmt.Println("config: ",zgo.K8s.GetContext(config.Host))
  fmt.Println("config-1: ",zgo.K8s.GetContext(config_1.Host))

  //5.生成clientset
  _, err = zgo.K8s.Builder().BuildClientSet(config.Host,config)
  if err != nil {
    return
  }

  //5-1生成另外一个clientset
  _, err = zgo.K8s.Builder().BuildClientSet(config_1.Host,config_1)
  if err != nil {
   return
  }

  //6.打印clientset
  fmt.Println("clientset: ",zgo.K8s.GetClientSet(config.Host))
  fmt.Println("clientset-1: ",zgo.K8s.GetClientSet(config_1.Host))

  //7. 打印version
  info, err := zgo.K8s.ServerVersion(config.Host)
  if err != nil {
   fmt.Println("==ServerVersion==err: ",err)
   return
  }
  zgo.Log.Infof("%s,%s", info.Platform,info.GitVersion)

  //7-1. 打印另一个version
  info_1, err := zgo.K8s.ServerVersion(config_1.Host)
  if err != nil {
  fmt.Println("err: ",err)
  return
  }
  fmt.Println("\n----")
  zgo.Log.Infof("%s,%s", info_1.Platform,info_1.GitVersion)

  //8.
  zgo.K8s.UseContext(config.Host) //使用指定的config.host的context
  dlist, err := zgo.K8s.Deployment().List(context.TODO(),"default","","",-1,false)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("总量：%v\n", len(dlist.Items))
  for _, item := range dlist.Items {
    fmt.Println(item.Namespace,item.Name)
  }
}
