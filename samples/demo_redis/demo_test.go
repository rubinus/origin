package demo_redis

import (
	"context"
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"sync"
	"testing"
	"time"
)

const (
	label_bj = "redis_label_bj"
)

func TestGet(t *testing.T) {

	err := zgo.Engine(&zgo.Options{
		Env:     "dev",
		Project: "1553240759",
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			_, err := zgo.Redis.Set(ctx, fmt.Sprintf("%s%d", "zgo:start:niubi:", i), i)
			if err != nil {
				panic(err)
			}
			wg.Done()
			select {

			case <-ctx.Done():
				fmt.Println("超时")
			default:
				//fmt.Print(r)
			}
		}(i)
	}
	wg.Wait()

}

type result struct {
	Abc string `json:"abc" redis:"abc"`
	Def string `json:"def" redis:"def"`
}

func TestSubscribe(t *testing.T) {

	err := zgo.Engine(&zgo.Options{
		Env:     "dev",
		Project: "1553240759",
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(3 * time.Second)

	go func() {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				ch, err := zgo.Redis.Publish(context.TODO(), "mychan", "lalala")
				fmt.Println(ch, err)

				r, _ := zgo.Redis.Hgetall(context.TODO(), "aaa")
				fmt.Printf("%+v====\n", r)
				bytes, err := zgo.Utils.Marshal(r)
				re := result{}
				zgo.Utils.Unmarshal(bytes, &re)
				fmt.Printf("%+v---------\n", re)

				get, err := zgo.Redis.Get(context.TODO(), "zgo:start:niubi:6")
				fmt.Println("get==", get.(string))

				zrangebyscore, err := zgo.Redis.Zrangebyscore(context.TODO(), "za", 0, 1000, true, 0, 101)
				fmt.Println("zrange--", zrangebyscore)

			}
		}
	}()

	Subscribe()

}

func TestXadd(t *testing.T) {
	err := zgo.Engine(&zgo.Options{
		Env:     "dev",
		Project: "1553240759",
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)

	Xadd()

	Xdel()

	Xrange()

	//GroupCreate()

	//Xack()

	go func() {
		ReadNew()
	}()

	Read()

}
