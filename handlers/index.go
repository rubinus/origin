package handlers

import (
	"fmt"
	"github.com/rubinus/origin/config"
	"github.com/rubinus/zgo"
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
	IP       string
}

func IndexPage(ctx iris.Context) {
	project := config.Conf.Project
	hostName, _ := os.Hostname()
	ctx.ViewData("", indexPage{
		Title:    project,
		Message:  fmt.Sprintf("%s welcome origin by zgo engine ...", project),
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
		Message:  fmt.Sprintf("%d -- %s origin by zgo engine ...", 404, project),
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
		Message:  fmt.Sprintf("%d -- %s origin by zgo engine ...", 500, project),
		Version:  config.Conf.Version,
		HostName: hostName,
		IP:       zgo.Utils.GetIntranetIP(),
	})
	ctx.View("500.html")
}
