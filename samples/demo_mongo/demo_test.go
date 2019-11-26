package demo_mongo

import (
	"context"
	"fmt"
	"git.zhugefang.com/gobase/origin/config"
	"git.zhugefang.com/gocore/zgo"
	"git.zhugefang.com/gocore/zgo/zgomongo"
	"testing"
	"time"
)

const (
	label_bj = "mongo_label_bj"
	label_sh = "mongo_label_sh"
)

//func TestGetUser(t *testing.T) {
//	//test samples for mongo query
//	config.InitConfig("local","","")
//	err := zgo.Engine(&zgo.Options{
//		Env:     "local",
//		Project: "origin",
//		Mongo: []string{
//			//label_bj,
//			label_sh,
//		},
//	})
//	//测试时表示使用mongo，在origin中使用一次
//
//	if err != nil {
//		panic(err)
//	}
//
//	GetUser()
//
//}

func TestUpdateNameById(t *testing.T) {
	config.InitConfig("local", "", "", "", "")
	err := zgo.Engine(&zgo.Options{
		Env:     "local",
		Project: "origin",
		Mongo: []string{
			//label_bj,
			label_sh,
		},
	})
	//测试时表示使用mongo，在origin中使用一次

	if err != nil {
		panic(err)
	}

	err = UpdateNameById()
	//err = DeleteById()
	//err = Delete()

	fmt.Println(err)

}

func TestGet(t *testing.T) {
	//test samples for mongo query
	config.InitConfig("local", "", "", "", "")
	err := zgo.Engine(&zgo.Options{
		Env:     "local",
		Project: "origin",
		Mongo: []string{
			label_bj,
			label_sh,
		},
	})
	//测试时表示使用mongo，在origin中使用一次

	if err != nil {
		panic(err)
	}

	clientBj, err := zgo.Mongo.New(label_bj)
	clientSh, err := zgo.Mongo.New(label_sh)
	if err != nil {
		panic(err)
	}

	//测试读取nsq数据，wait for sdk init connection
	time.Sleep(2 * time.Second)

	var replyChan = make(chan int)
	var countChan = make(chan int)
	l := 10000 //暴力测试50000个消息，时间10秒，本本的并发每秒5000

	count := []int{}
	total := []int{}
	stime := time.Now()

	for i := 0; i < l; i++ {
		go func(i int) {
			countChan <- i //统计开出去的goroutine
			if i%2 == 0 {
				ch := createMongo(label_bj, clientBj, i)
				reply := <-ch
				replyChan <- reply
				<-getMongo(label_bj, clientBj, i)

			} else {
				ch := createMongo("user", clientSh, i)
				reply := <-ch
				replyChan <- reply
				<-getMongo("user", clientSh, i)

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

type user struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Age   int    `json:"age"`
}

func getMongo(label string, client zgomongo.Mongoer, i int) chan int {

	//还需要一个上下文用来控制开出去的goroutine是否超时
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	//输入参数：上下文ctx，mongoChan里面是client的连接，args具体的查询操作参数
	args := make(map[string]interface{})
	args["db"] = "test"
	args["table"] = label
	args["query"] = make(map[string]interface{})
	result := &user{}
	args["obj"] = result
	err := client.Get(ctx, args)
	if err != nil {
		panic(err)
	}
	//fmt.Println(result.Age, result.Label, "======")

	out := make(chan int, 1)
	select {
	case <-ctx.Done():
		fmt.Println("超时")
		out <- 10001
		return out
	default:
		bytes, err := zgo.Utils.Marshal(result)
		nu := &user{}
		zgo.Utils.Unmarshal(bytes, nu)
		fmt.Println(nu)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(bytes), err, "---from mongo successful---",nu)
		out <- 1
	}

	return out

}

func createMongo(label string, client zgomongo.Mongoer, i int) chan int {

	//还需要一个上下文用来控制开出去的goroutine是否超时
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	//输入参数：上下文ctx，mongoChan里面是client的连接，args具体的查询操作参数
	args := make(map[string]interface{})
	args["db"] = "test"
	args["table"] = label
	args["items"] = &user{
		Name:  fmt.Sprintf("%s-%d", label, i),
		Label: label,
		Age:   i,
	}
	result := &user{}
	args["obj"] = result
	err := client.Create(ctx, args)
	if err != nil {
		panic(err)
	}

	out := make(chan int, 1)
	select {
	case <-ctx.Done():
		fmt.Println("超时")
		out <- 10001
		return out
	default:
		bytes, err := zgo.Utils.Marshal(result)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes), err, "---from mongo successful---", label)
		out <- 1
	}

	return out

}
