package demo_mysql

import (
	"context"
	"fmt"
	"github.com/gitcpu-io/zgo"
	"sync"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	b := &MysqlDemo{}
	a, err := b.Create("新增的数据")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

	a, err = b.Get(a.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

	i, err := b.UpdateByObj(a.Id, "第一次修改")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	i, err = b.UpdateByObj(a.Id, "第二次修改")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	i, err = b.UpdateByObj(a.Id, "第三次修改")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	c, err := b.List()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}

func TestMysqlDemo_Get(t *testing.T) {

	//查询参数
	zgo.Engine(&zgo.Options{
		Env:     "dev",
		Project: "1553240759",

		Mysql: []string{
			"mysql_sell_2",
		},
	})
	time.Sleep(1 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 输入参数：上下文ctx，args具体的查询操作参数

	t1 := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			args := make(map[string]interface{})
			obj := &House{
				Name: "mysql",
			}
			// 拼接dbname   如果不随城市变化，固定城市只用写表名即可
			args["table"] = "user"

			args["query"] = " id = ? "      //  用问号作为占位符
			args["args"] = []interface{}{1} // 参数 放入interface的slice里面
			args["obj"] = obj               // 输出对象
			zgo.Mysql.Get(ctx, args)

			fmt.Println(obj)

			if i%1000 == 0 {
				fmt.Println(i, obj)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

}
