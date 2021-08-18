package handlers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/rubinus/origin/config"
	"github.com/rubinus/zgo"
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
	IP       string
}

func IndexPage(ctx iris.Context) {
	project := config.Conf.Project
	hostName, _ := os.Hostname()
	ctx.ViewData("", indexPage{
		Title:    project,
		Message:  fmt.Sprintf("%s welcome by zgo engine %s ...", project, zgo.Version),
		Version:  config.Conf.Version,
		HostName: hostName,
		IP:       zgo.Utils.GetIntranetIP(),
	})
	ctx.View("index.html")
}

func FourZeroFourPage(ctx iris.Context) {
	project := config.Conf.Project
	hostName, _ := os.Hostname()
	ctx.ViewData("", indexPage{
		Title:    project,
		Message:  fmt.Sprintf("%d -- %s by zgo engine %s ...", 404, project, zgo.Version),
		Version:  config.Conf.Version,
		HostName: hostName,
		IP:       zgo.Utils.GetIntranetIP(),
	})
	ctx.View("404.html")
}

func FiveZeroZeroPage(ctx iris.Context) {
	project := config.Conf.Project
	hostName, _ := os.Hostname()
	ctx.ViewData("", indexPage{
		Title:    project,
		Message:  fmt.Sprintf("%d -- %s by zgo engine %s ...", 500, project, zgo.Version),
		Version:  config.Conf.Version,
		HostName: hostName,
		IP:       zgo.Utils.GetIntranetIP(),
	})
	ctx.View("500.html")
}

func Health(ctx iris.Context) {
	ctx.JSONP(map[string]string{
		"health": "true",
	})
}
