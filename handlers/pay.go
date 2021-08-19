package handlers

import (
	"github.com/kataras/iris/v12"
	"github.com/rubinus/origin/models"
	"github.com/rubinus/origin/services"
	"github.com/rubinus/zgo"
	"strings"
)

/*
@Time : 2019-03-22 11:50
@Author : rubinus.chu
@File : redis
@project: origin
*/

// DoPay 使用MVC模式
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
