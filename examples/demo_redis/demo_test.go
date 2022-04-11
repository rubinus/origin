package demo_redis

import (
	"context"
	"fmt"
	"github.com/gitcpu-io/zgo"
  "os"
  "strings"
  "sync"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
  // 准备cpath
  var cpath string
  if cpath == "" {
    pwd, err := os.Getwd()
    if err == nil {
      arr := strings.Split(pwd,"/")
      cp := strings.Join(arr[:len(arr) - 2],"/")
      cpath = fmt.Sprintf("%s/%s", cp, "configs")
    }
  }
	err := zgo.Engine(&zgo.Options{
	  CPath: cpath,
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
			if false {
				fmt.Println(111)
			}
			select {
			case <-time.Tick(1 * time.Second):
				ch, err := zgo.Redis.Publish(context.TODO(), "mychan", "lalala")
				fmt.Println(ch, err)

				r, _ := zgo.Redis.Hgetall(context.TODO(), "aaa")
				fmt.Printf("%+v====\n", r)
				bytes, err := zgo.Utils.Marshal(r)
				if err != nil {
					zgo.Log.Error(err)
				}
				re := result{}
				err = zgo.Utils.Unmarshal(bytes, &re)
				if err != nil {
					zgo.Log.Error(err)
				}
				fmt.Printf("%+v---------\n", re)

				get, err := zgo.Redis.Get(context.TODO(), "zgo:start:niubi:6")
				if err != nil {
					zgo.Log.Error(err)
				}
				fmt.Println("get==", get.(string))

				zrangebyscore, err := zgo.Redis.Zrangebyscore(context.TODO(), "za", 0, 1000, true, 0, 101)
				if err != nil {
					zgo.Log.Error(err)
				}
				fmt.Println("zrange--", zrangebyscore)
			default:

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
