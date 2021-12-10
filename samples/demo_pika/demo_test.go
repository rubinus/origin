package demo_pika

import (
	"context"
	"fmt"
	"github.com/gitcpu-io/origin/config"
	"github.com/gitcpu-io/zgo"
	"testing"
	"time"
)

const (
	sell_label_rw = "pika_label_rw"
	sell_label_r  = "pika_label_rw"
)

func TestGet(t *testing.T) {
	config.InitConfig("local", "", "", "", "")

	err := zgo.Engine(&zgo.Options{
		Env:     config.Conf.Env,
		Project: config.Conf.Project,
		Pika: []string{
			sell_label_rw,
			sell_label_r,
		},
	})
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	BJRW, _ := zgo.Pika.New(sell_label_rw)

	BJR, _ := zgo.Pika.New(sell_label_r)

	//fmt.Println(BJRW)

	result, err0 := BJRW.Hset(ctx, "china_online", "name", "pika_test9")
	if err0 != nil {
		panic(err0)
	}

	time.Sleep(2 * time.Millisecond)

	result1, err := BJR.Hget(ctx, "china_online", "name")

	fmt.Println(result, result1)
	if err != nil {
		panic(err)
	}

	select {

	case <-ctx.Done():
		fmt.Println("超时")
	default:
		fmt.Print(result)
	}
}
