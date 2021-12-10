package demo_etcd

import (
	"fmt"
	"github.com/gitcpu-io/zgo"
	"testing"
	"time"
)

/*
@Time : 2019-06-05 10:46
@Author : rubinus.chu
@File : demo_test
@project: origin
*/

const (
	label_bj = "etcd_label"
)

var project = "1553240759"

func TestGet(t *testing.T) {
	err := zgo.Engine(&zgo.Options{
		Env:     "dev",
		Project: project,

		Etcd: []string{
			label_bj,
		},
	}) //测试时表示使用neo4j，在origin中使用一次

	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	etcdCh, err := zgo.Etcd.GetConnChan() //有多个时，需要指定label
	if err != nil {
		panic(err)
	}

	if etcd, ok := <-etcdCh; ok {
		r, err := Get(etcd)
		fmt.Println(err, r)
	}

}
