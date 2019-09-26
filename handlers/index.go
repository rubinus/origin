package handlers

import (
	"fmt"
	"git.zhugefang.com/gobase/origin/config"
	"github.com/kataras/iris"
	"os"
)

/*
@Time : 2019-03-06 19:46
@Author : rubinus.chu
@File : demoHandler
@project: origin
*/

type indexPage struct {
	Title    string
	Message  string
	Version  string
	HostName string
}

func IndexPage(ctx iris.Context) {
	//ctx.WriteString("追踪 ...")
	project := config.Conf.Project
	hostName, _ := os.Hostname()
	ctx.ViewData("", indexPage{
		Title:    project,
		Message:  fmt.Sprintf("%s welcome origin by zgo engine ...", project),
		Version:  config.Conf.Version,
		HostName: hostName,
	})
	ctx.View("index.html")
}
