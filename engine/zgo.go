package engine

import (
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/zgo"
  "time"
)

func Run() error {
  err := zgo.Engine(&zgo.Options{
    CPath:     config.Conf.CPath,
    Env:       config.Conf.Env,
    Loglevel:  config.Conf.Loglevel,
    Project:   config.Conf.Project,
    EtcdHosts: config.Conf.EtcdHosts,

    /**
    注意local.json方式
    ********************
    如果是在本地开发可以对下面的组件选择是否使用，如果是非local，不需要填写，应用的配置是从etcd配置中心读取的
    ********************
    */
    Kafka: []string{
      //"kafka_label_bj",
      //"kafka_label_sh",
    },
    Nsq: []string{
      //"nsq_label_bj",
    },
    Redis: []string{
      //"redis_label_sh",	//测试时可以放开注释，通过配置文件来调试连接中间件redis
    },
    Mongo: []string{
      "mongo_label_bj",	//测试时可以放开注释，通过配置文件来调试连接中间件mongodb
    },
    Mysql: []string{
      //"mysql_sell_1",
      //"mysql_sell_2",
    },
  })

  time.Sleep(1 * time.Second) //wait 1 second for zgo engine start and _init connection
  return err
}
