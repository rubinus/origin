package handlers

import (
  "github.com/gitcpu-io/origin/models"
  "github.com/gitcpu-io/origin/services"
  "github.com/gitcpu-io/zgo"
  "github.com/kataras/iris/v12"
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
      _, err := zgo.Http.JsonpErr(ctx, "json body is error，"+err.Error())
      if err != nil {
        zgo.Log.Error(err)
      }
      return
    }
  } else {
    _, err := zgo.Http.JsonpErr(ctx, "pls send application/json")
    if err != nil {
      zgo.Log.Error(err)
    }
    return
  }

  if request.Bid == "" || request.Aid == 0 {
    _, err := zgo.Http.JsonpErr(ctx, "业务线和事件名不能为空")
    if err != nil {
      zgo.Log.Error(err)
    }
    return
  }

  pay := services.NewPay()
  tcb, err := pay.Insert(request)
  if err != nil {
    zgo.Log.Error(err)
    _, err := zgo.Http.JsonpErr(ctx, err.Error())
    if err != nil {
      zgo.Log.Error(err)
    }
    return
  }

  r := make(map[string]interface{})
  r["status"] = "done"
  r["id"] = tcb.Id

  _, err = zgo.Http.JsonpOK(ctx, r)
  if err != nil {
    zgo.Log.Error(err)
  }
}
