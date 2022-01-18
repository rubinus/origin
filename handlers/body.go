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
      zgo.Http.JsonpErr(ctx, "json body is error，"+err.Error())
      return
    }
  } else {
    zgo.Http.JsonpErr(ctx, "pls send application/json")
    return
  }

  zgo.Http.JsonpOK(ctx, "OK")

}
