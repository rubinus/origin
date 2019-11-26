package demo_lb

import (
	"fmt"
	"git.zhugefang.com/gobase/origin/config"
	"git.zhugefang.com/gocore/zgo"
)

/*
@Time : 2019-04-29 14:34
@Author : rubinus.chu
@File : demo
@project: origin
*/

func CallLb() {
	config.InitConfig("local", "", "", "", "")

	err := zgo.Engine(&zgo.Options{
		Env:      config.Conf.Env,
		Project:  config.Conf.Project,
		Loglevel: config.Conf.Loglevel,
	})

	if err != nil {
		panic(err)
	}

	//lb := zgo.LB.WR2("127.0.0.1:8009", "127.0.0.1:8008", "127.0.0.1:8007")
	lb := zgo.LB.WR2()
	lb.AddWeight("127.0.0.1:9000", 2)
	lb.AddWeight("127.0.0.1:9001", 3)
	lb.AddWeight("127.0.0.1:9002", 4)
	for i := 0; i < 9; i++ {
		host, err := lb.Balance()

		if err != nil {
			fmt.Println(err, "=======")
			continue
		}

		fmt.Printf("Send request #%d to host %s\n", i, host)
	}

}
