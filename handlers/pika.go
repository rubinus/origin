package handlers

import (
	"context"
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"github.com/kataras/iris"
	"time"
)

/*
@Time : 2019-03-22 11:50
@Author : rubinus.chu
@File : redis
@project: integral
*/

func PikaGet(ctx iris.Context) {
	name := ctx.URLParam("name")

	var errStr string
	cotx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond) //you can change this time number
	defer cancel()

	key := fmt.Sprintf("%s:%s:%s", "zgo", "start", name)

	//get String key return a map
	val, err := zgo.Pika.Get(cotx, key)

	result := zgo.Utils.StringToMap(val.(string))

	if err != nil {
		zgo.Log.Error(err)
		return
	}
	select {
	case <-cotx.Done():
		errStr = "call pika get string timeout"
		zgo.Log.Error(errStr) //通过zgo.Log统计日志
		//ctx.JSONP(iris.Map{"status": 201, "msg": errStr}) //返回jsonp格式
		zgo.Http.JsonpErr(ctx, errStr)
	default:
		//ctx.JSONP(iris.Map{"status": 200, "data": result})
		zgo.Http.JsonpOK(ctx, result)

	}

}

func PikaPut(ctx iris.Context) {
	name := ctx.URLParam("name")

	var errStr string
	cotx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond) //you can change this time number
	defer cancel()

	key := fmt.Sprintf("%s:%s:%s", "zgo", "start", name)

	s := `{"branch":"beta","change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`

	//set stander json to redis String
	result, err := zgo.Pika.Set(cotx, key, s)
	if err != nil {
		zgo.Log.Error(err)
		return
	}

	select {
	case <-cotx.Done():
		errStr = "call pika put string timeout"
		zgo.Log.Error(errStr) //通过zgo.Log统计日志
		//ctx.JSONP(iris.Map{"status": 201, "msg": errStr}) //返回jsonp格式
		zgo.Http.JsonpErr(ctx, errStr)
	default:
		//ctx.JSONP(iris.Map{"status": 200, "data": result})
		zgo.Http.JsonpOK(ctx, result)

	}

}

func PikaHashGet(ctx iris.Context) {
	name := ctx.URLParam("name")

	var errStr string
	cotx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond) //you can change this time number
	defer cancel()

	key := fmt.Sprintf("%s_%s_%s", "zgo", "start", name)

	//get String key return a map
	val, err := zgo.Pika.Hget(cotx, key, "test")

	result := zgo.Utils.StringToMap(val.(string))

	if err != nil {
		zgo.Log.Error(err)
		return
	}
	select {
	case <-cotx.Done():
		errStr = "call pika get string timeout"
		zgo.Log.Error(errStr) //通过zgo.Log统计日志
		//ctx.JSONP(iris.Map{"status": 201, "msg": errStr}) //返回jsonp格式
		zgo.Http.JsonpErr(ctx, errStr)
	default:
		//ctx.JSONP(iris.Map{"status": 200, "data": result})
		zgo.Http.JsonpOK(ctx, result)

	}

}

func PikaHashPut(ctx iris.Context) {
	name := ctx.URLParam("name")

	var errStr string
	cotx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond) //you can change this time number
	defer cancel()

	key := fmt.Sprintf("%s_%s_%s", "zgo", "start", name)

	s := `{"branch":"beta","change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`

	//set stander json to redis String
	result, err := zgo.Pika.Hset(cotx, key, "test", s)
	if err != nil {
		zgo.Log.Error(err)
		return
	}

	select {
	case <-cotx.Done():
		errStr = "call pika put string timeout"
		zgo.Log.Error(errStr) //通过zgo.Log统计日志
		//ctx.JSONP(iris.Map{"status": 201, "msg": errStr}) //返回jsonp格式
		zgo.Http.JsonpErr(ctx, errStr)
	default:
		//ctx.JSONP(iris.Map{"status": 200, "data": result})
		zgo.Http.JsonpOK(ctx, result)

	}

}
