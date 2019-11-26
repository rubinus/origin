package demo_clickhouse

import (
	"git.zhugefang.com/gocore/zgo"
	"testing"
	"time"
)

/*
@Time : 2019-09-26 16:34
@Author : rubinus.chu
@File : demo.test
@project: origin
*/

func TestGet(t *testing.T) {
	err := zgo.Engine(&zgo.Options{
		Env:     "dev",
		Project: "1553240759",
		ClickHouse: []string{
			label_bj,
		},
	})
	//测试时表示使用clickhouse，在origin中使用一次
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)

	Get()

}
