package demo_cache

import (
	"context"
	"fmt"
	"github.com/gitcpu-io/origin/configs"
	"github.com/gitcpu-io/zgo"
	"time"
)

type CacheDemo struct {
}

func init() {
	configs.InitConfig("", "local", "", "", 0, 0)
}

//QueryMysql 测试读取Mysqldb数据，wait for sdk init connection
func (m CacheDemo) run() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	//查询参数
	err := zgo.Engine(&zgo.Options{
		CPath:   configs.Conf.CPath,
		Env:     configs.Conf.Env,
		Project: configs.Conf.Project,
		Pika: []string{
			"pika_label_rw", // 需要一个pika的配置
		},
	})
	if err != nil {
		zgo.Log.Error(err)
		return
	}
	param := make(map[string]interface{})
	param["ceshi1"] = 1
	param["ceshi2"] = 2
	param["ceshi3"] = 2
	param["ceshi4"] = 2
	param["ceshi5"] = 2
	param["ceshi6"] = 2
	param["ceshi7"] = 2
	h := make(map[string]interface{})
	// 无缓存
	start := time.Now().UnixNano()
	err = m.GetData(ctx, param, &h)
	if err != nil {
		zgo.Log.Error(err)
	}
	fmt.Println(h)
	fmt.Println("正常用时", (time.Now().UnixNano()-start-2000000000)/1000)
	fmt.Println("")

	// 正常缓存
	h1 := make(map[string]interface{})
	start = time.Now().UnixNano()
	err = zgo.Cache.Decorate(m.GetData, 1)(ctx, param, &h1)
	if err != nil {
		zgo.Log.Error(err)
		return
	}
	fmt.Println(h1)
	fmt.Println("第一次请求用时", (time.Now().UnixNano()-start-2000000000)/1000)

	// 正常缓存第二次请求
	fmt.Println("")
	fmt.Println("-------第二次请求开始-----")
	start = time.Now().UnixNano()
	h2 := make(map[string]interface{})
	err = zgo.Cache.Decorate(m.GetData, 10000000)(ctx, param, &h2)
	if err != nil {
		zgo.Log.Error(err)
		return
	}
	fmt.Println(h2)
	fmt.Println("第二次请求用时", (time.Now().UnixNano()-start)/1000)

	start = time.Now().UnixNano()
	fmt.Println("")
	fmt.Println("")
	fmt.Println(start)
	fmt.Println("降级缓存测试：")
	// 降级缓存正常情况
	h3 := make(map[string]interface{})
	err = zgo.Cache.TimeOutDecorate(m.GetData1, 10)(ctx, param, &h3)
	if err != nil {
		zgo.Log.Error(err)
	}
	fmt.Println(h3)
	fmt.Println("正常降级缓存用时", (time.Now().UnixNano()-start-2000000000)/1000)
	fmt.Println("")
	fmt.Println("")
	start = time.Now().UnixNano()
	// 超时情况
	h4 := make(map[string]interface{})
	err = zgo.Cache.TimeOutDecorate(m.GetData1, 1)(ctx, param, &h4)
	if err != nil {
		zgo.Log.Error(err)
	}
	fmt.Println(h4)
	fmt.Println("超时降级缓存用时", (time.Now().UnixNano()-start-1000000000)/1000)
}

func (m CacheDemo) GetData(ctx context.Context, param map[string]interface{}, obj interface{}) error {
	time.Sleep(2 * time.Second)
	data := (*obj.(*map[string]interface{}))
	data["test"] = "测试数据"
	return nil
}

func (m CacheDemo) GetData1(ctx context.Context, param map[string]interface{}, obj interface{}) error {
	time.Sleep(2 * time.Second)
	data := (*obj.(*map[string]interface{}))
	data["test"] = "测试数据"
	return nil
}
