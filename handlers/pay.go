package handlers

import (
	"context"
	"fmt"
	"github.com/rubinus/origin/models"
	"github.com/rubinus/origin/services"
	"github.com/rubinus/zgo"
	"github.com/kataras/iris"
	"strings"
	"time"
)

/*
@Time : 2019-03-22 11:50
@Author : rubinus.chu
@File : redis
@project: origin
*/

func RedisGet(ctx iris.Context) {
	name := ctx.URLParam("name")

	var errStr string
	cotx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond) //you can change this time number
	defer cancel()

	key := fmt.Sprintf("%s:%s:%s", "zgo", "start", name)

	//get String key return a map
	val, err := zgo.Redis.Get(cotx, key)

	result := zgo.Utils.StringToMap(val.(string))

	if err != nil {
		zgo.Log.Error(err)
		return
	}

	select {
	case <-cotx.Done():
		errStr = "call redis get string timeout"
		zgo.Log.Error(errStr) //通过zgo.Log统计日志
		//ctx.JSONP(iris.Map{"status": 201, "msg": errStr}) //返回jsonp格式
		zgo.Http.JsonpErr(ctx, errStr)
	default:
		//ctx.JSONP(iris.Map{"status": 200, "data": result})
		zgo.Http.JsonpOK(ctx, result)

	}

}

func DoPay(ctx iris.Context) {

	request := &models.PayRequest{}
	if strings.Contains(ctx.GetContentTypeRequested(), "json") {
		if err := ctx.ReadJSON(request); err != nil {
			zgo.Http.JsonpErr(ctx, "json body is error，"+err.Error())
			return
		}
	} else {
		zgo.Http.JsonpErr(ctx, "pls send application/json")
		return
	}

	if request.Bid == "" || request.Aid == 0 {
		zgo.Http.JsonpErr(ctx, "业务线和事件名不能为空")
		return
	}

	pay := services.NewPay()
	tcb, err := pay.Insert(request)
	if err != nil {
		zgo.Log.Error(err)
		zgo.Http.JsonpErr(ctx, err.Error())
		return
	}

	r := make(map[string]interface{})
	r["status"] = "done"
	r["id"] = tcb.Id

	zgo.Http.JsonpOK(ctx, r)

}
