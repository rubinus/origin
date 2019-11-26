package demo_log

import (
	"git.zhugefang.com/gobase/origin/config"
	"git.zhugefang.com/gobase/origin/engine"
	"git.zhugefang.com/gocore/zgo"
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
