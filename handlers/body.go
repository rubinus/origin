package handlers

import (
  "github.com/gitcpu-io/origin/models"
  "github.com/gitcpu-io/zgo"
  "github.com/kataras/iris/v12"
  "strings"
)

// DoPostBody 使用MVC模式
func DoPostBody(ctx iris.Context) {
  request := &models.PayRequest{}
  if strings.Contains(ctx.GetContentTypeRequested(), "json") {
    if err := ctx.ReadJSON(request); err != nil {
      _, err := zgo.Http.JsonpErr(ctx, "json body is error，"+err.Error())
      if err != nil {
        zgo.Log.Error(err)
        return
      }
      return
    }
  } else {
    _, err := zgo.Http.JsonpErr(ctx, "pls send application/json")
    if err != nil {
      zgo.Log.Error(err)
      return
    }
    return
  }

  _, err := zgo.Http.JsonpOK(ctx, "OK")
  if err != nil {
    zgo.Log.Error(err)
    return
  }

}
