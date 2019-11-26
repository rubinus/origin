package demo_file

import (
	"git.zhugefang.com/gocore/zgo"
	"os"
	"strings"
)

/*
@Time : 2019-02-28 20:37
@Author : rubinus.chu
@File : demo
@project: origin.git
*/

func F() {
	input := strings.NewReader("hello world")
	pn, err := zgo.File.Put("/x/a.txt", input)
	if err != nil {
		zgo.Log.Debug(err)
	}
	zgo.Log.Info("put bytes num:", pn)

	n, err := zgo.File.Get("/tmp/x/a.txt", os.Stdout)
	if err != nil {
		zgo.Log.Debug(err)
	}
	zgo.Log.Info("----", n, "====zgo.file.get=======")

	size, err := zgo.File.Size("/x/a.txt")
	if err != nil {
		zgo.Log.Debug(err)
	}
	zgo.Log.Info("file size:", size)
	if n != size {
		zgo.Log.Error("error file size")
	}
}
