package demo_log

import (
	"github.com/rubinus/origin/config"
	"github.com/rubinus/origin/engine"
	"github.com/rubinus/zgo"
)

const (
	debug = iota
	info
	warn
	error
)

func init() {
	config.InitConfig("dev", "1553240759", "", "", "")
}

func Call() {
	engine.Run()

	zgo.Log.Debug("debug")
	zgo.Log.Info("info")
	zgo.Log.Warn("warn")
	zgo.Log.Error("error")

	//l.Info(222)

	//fmt.Println(debug)
	//fmt.Println(info)
	//fmt.Println(warn)
	//fmt.Println(error)
}
