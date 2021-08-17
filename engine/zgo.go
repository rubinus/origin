package engine

import (
	"github.com/rubinus/origin/config"
	"github.com/rubinus/zgo"
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
		注意
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
			//"redis_label_bj",
		},
		Mysql: []string{
			//"mysql_sell_1",
			//"mysql_sell_2",
		},
	})

	time.Sleep(1 * time.Second) //wait 1 second for zgo engine start and init connection
	return err
}
