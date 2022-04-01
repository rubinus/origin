package demo_etcd

import (
	"context"
	"github.com/gitcpu-io/zgo"
)

/*
@Time : 2019-06-05 10:45
@Author : rubinus.chu
@File : demo
@project: origin
*/

//var project1 = "origin"
//var label = "etcd_label"

var project1 = "1553240759"
var label = "1446177640400"

func Get(cli *zgo.EtcdClientV3) (*zgo.EtcdGetResponse, error) {
	key := "zgo/project/" + project1 + "/etcd/" + label
	res, err := cli.KV.Get(context.TODO(), key)
	if err != nil {
		return nil, err
	}
	return res, nil
}
