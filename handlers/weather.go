package handlers

import (
  "context"
  "github.com/gitcpu-io/origin/models"
  "github.com/gitcpu-io/origin/services"
  "github.com/gitcpu-io/zgo"
  "github.com/kataras/iris/v12"
  "strings"
  "time"
)

/*
@Time : 2022-03-31 17:50
@Author : rubinus.chu
@File : weather
@project: origin
*/

// SaveWeather 使用MVC模式
func SaveWeather(ctx iris.Context) {
  // 第一：定义错误返回变量，请求上下文，通过defer来最后响应
  var errStr string

  cotx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //you can change this time number
  defer cancel()

  defer func() {
    if errStr != "" {
      _, err := zgo.Http.JsonpErr(ctx, errStr)
      if err != nil {
        zgo.Log.Error(err)
      }
    }
  }()

  //第二步：解析请求参数
  request := &models.WeatherRequest{}
  if strings.Contains(ctx.GetContentTypeRequested(), "json") {
    if err := ctx.ReadJSON(request); err != nil {
      _, err := zgo.Http.JsonpErr(ctx, "json body is error，"+err.Error())
      if err != nil {
        errStr = err.Error()
        zgo.Log.Error(err)
      }
      return
    }
  } else {
    _, err := zgo.Http.JsonpErr(ctx, "pls send application/json")
    if err != nil {
      errStr = err.Error()
      zgo.Log.Error(err)
    }
    return
  }

  if request.Query == "" {
    _, err := zgo.Http.JsonpErr(ctx, "查询城市不能为空")
    if err != nil {
      errStr = err.Error()
      zgo.Log.Error(err)
    }
    return
  }

  // 第三：调用zgo engine来处理业务逻辑
  wea := services.NewWeather()
  res, err := wea.Insert(cotx, request)
  if err != nil {
    errStr = err.Error()
    zgo.Log.Error(err)
    _, err = zgo.Http.JsonpErr(ctx, err.Error())
    if err != nil {
      errStr = err.Error()
      zgo.Log.Error(err)
    }
    return
  }

  // 第四：使用select来响应处理结果与超时
  select {
  case <-cotx.Done():
    errStr = "call weather save timeout"
    zgo.Log.Error(errStr) //通过zgo.Log统计日志
  default:
    _, err := zgo.Http.JsonpOK(ctx, res)
    if err != nil {
      zgo.Log.Error(err)
    }
  }

}

// ListWeather 使用MVC模式
func ListWeather(ctx iris.Context) {
  // 第一：定义错误返回变量，请求上下文，通过defer来最后响应
  var errStr string

  cotx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //you can change this time number
  defer cancel()

  defer func() {
    if errStr != "" {
      _, err := zgo.Http.JsonpErr(ctx, errStr)
      if err != nil {
        zgo.Log.Error(err)
      }
    }
  }()

  // 第二：解析请求参数
  city := ctx.URLParam("city")
  if city == "" {
    errStr = "必须输入查询城市参数"
    return
  }

  // 第三：调用zgo engine来处理业务逻辑
  wea := services.NewWeather()
  res, err := wea.List(cotx, city)
  if err != nil {
    errStr = err.Error()
    zgo.Log.Error(err)
    return
  }

  // 第四：使用select来响应处理结果与超时
  select {
  case <-cotx.Done():
    errStr = "call weather list timeout"
    zgo.Log.Error(errStr) //通过zgo.Log统计日志
  default:
    _, err := zgo.Http.JsonpOK(ctx, res)
    if err != nil {
      zgo.Log.Error(err)
    }
  }

}
