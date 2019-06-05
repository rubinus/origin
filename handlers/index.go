package handlers

import (
	"context"
	"fmt"
	"git.zhugefang.com/gobase/integral/config"
	"git.zhugefang.com/gobase/integral/models/mysql_models"
	"git.zhugefang.com/gobase/integral/services"
	"git.zhugefang.com/gocore/zgo"
	"github.com/kataras/iris"
	"os"
	"time"
)

/*
@Time : 2019-03-06 19:46
@Author : rubinus.chu
@File : demoHandler
@project: integral
*/

type indexPage struct {
	Title    string
	Message  string
	Version  string
	HostName string
}

func IndexPage(ctx iris.Context) {
	//ctx.WriteString("追踪 ...")
	project := config.Conf.Project
	hostName, _ := os.Hostname()
	ctx.ViewData("", indexPage{
		Title:    project,
		Message:  fmt.Sprintf("%s welcome by zgo engine ...", project),
		Version:  config.Conf.Version,
		HostName: hostName,
	})
	ctx.View("index.html")
}

func Hello(ctx iris.Context) {
	name := ctx.URLParam("name")

	var errStr string
	cotx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond) //you can change this time number
	defer cancel()

	//key := fmt.Sprintf("%s:%s:%s", "zgo", "start", name)

	s := `{"branch":"beta","change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`

	//set stander json to redis String
	_, err := zgo.Redis.Set(cotx, "integral_"+name, s)
	if err != nil {
		zgo.Log.Error(err)
		return
	}

	//get String key return a map
	val, err := zgo.Redis.Get(cotx, "integral_"+name)

	result := zgo.Utils.StringToMap(val.(string))

	fmt.Println(result, "-------")

	if err != nil {
		zgo.Log.Error(err)
		return
	}

	//result, err := zgo.Redis.Hgetall(cotx, key)
	//if err != nil {
	//	zgo.Log.Error(err)
	//ctx.JSONP(iris.Map{"status": 201, "msg": errStr}) //返回jsonp格式
	//zgo.Http.JsonpErr(ctx, errStr)
	//return
	//}

	//发送到nsq
	//n, err := zgo.Nsq.New("nsq_label_bj")
	//if err != nil {
	//	panic(err)
	//}
	//n.Producer(cotx, "integral", []byte("nsq888888"))
	//
	////发送到kafka
	//k, err := zgo.Kafka.New("kafka_label_bj")
	//if err != nil {
	//	panic(err)
	//}
	//k.Producer(cotx, "integral", []byte("kafka9999999999"))

	select {
	case <-cotx.Done():
		errStr = "call redis hgetall timeout"
		zgo.Log.Error(errStr) //通过zgo.Log统计日志
		//ctx.JSONP(iris.Map{"status": 201, "msg": errStr}) //返回jsonp格式
		zgo.Http.JsonpErr(ctx, errStr)
	default:
		//ctx.JSONP(iris.Map{"status": 200, "data": result})
		zgo.Http.JsonpOK(ctx, result)

	}

}

// AddHouse 保存房源信息的handler方法
func AddHouse(ctx iris.Context) {
	zctx := context.Background()
	city := ctx.Params().GetString("city")
	h := &mysql_models.House{}
	err := ctx.ReadJSON(h)
	if err != nil {
		zgo.Http.JsonpErr(ctx, err.Error())
	}
	fmt.Println(h)
	// 验证数据是否有Id
	if h.Id > 0 {
		zgo.Http.JsonpErr(ctx, "Id不为空")
		return
	}

	hs := services.HouseService{}
	err = hs.AddHouse(zctx, h, city)
	if err != nil {
		zgo.Http.JsonpErr(ctx, err.Error())
		return
	}

	zgo.Http.JsonpOK(ctx, h)
}
