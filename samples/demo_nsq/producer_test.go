package demo_nsq

import (
	"fmt"
	"github.com/rubinus/zgo"
	"testing"
	"time"
)

const (
	label_bj = "nsq_label_bj"
	label_sh = "nsq_label_sh"
)

var project = "1553240759"

func TestProducer(t *testing.T) {
	err := zgo.Engine(&zgo.Options{
		Env:     "dev",
		Project: project,

		Nsq: []string{
			label_bj,
			label_sh,
		},
	}) //测试时表示使用nsq，在origin中使用一次

	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	clientBj, err := zgo.Nsq.New()
	clientSh, err := zgo.Nsq.New()
	if err != nil {
		panic(err)
	}
	//map[string][]map[string]string

	//测试读取nsq数据，wait for sdk init connection
	time.Sleep(2 * time.Second)

	var replyChan = make(chan int)
	var countChan = make(chan int)
	l := 20 //暴力测试50000个消息，时间10秒，本本的并发每秒5000

	count := []int{}
	total := []int{}
	stime := time.Now()

	for i := 0; i < l; i++ {
		go func(i int) {
			countChan <- i //统计开出去的goroutine
			if i%2 == 0 {
				ch := Producer(project, clientBj, i, false)
				reply := <-ch
				replyChan <- reply

			} else {
				ch := Producer(project, clientSh, i, false)
				reply := <-ch
				replyChan <- reply
			}
		}(i)
	}

	go func() {
		for v := range replyChan {
			if v != 10001 { //10001表示超时
				count = append(count, v) //成功数
			}
		}
	}()

	go func() {
		for v := range countChan { //总共的goroutine
			total = append(total, v)
		}
	}()

	for _, v := range count {
		if v != 1 {
			fmt.Println("有不成功的")
		}
	}

	for {
		if len(count) == l {
			var timeLen time.Duration
			timeLen = time.Now().Sub(stime)

			fmt.Printf("总消耗时间：%s, 成功：%d, 总共开出来的goroutine：%d\n", timeLen, len(count), len(total))
			break
		}

		select {
		case <-time.Tick(time.Duration(1000 * time.Millisecond)):
			fmt.Println("处理进度每1000毫秒", len(count))

		}
	}
	time.Sleep(2 * time.Second)
}
