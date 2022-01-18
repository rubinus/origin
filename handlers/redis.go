package handlers

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/zgo"
  "github.com/kataras/iris/v12"
  "time"
)

// RedisGet RedisGet接口四板斧，这仅仅是一个例子
func RedisGet(ctx iris.Context) {
  // 第一：定义错误返回变量，请求上下文，通过defer来最后响应
  var errStr string

  cotx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //you can change this time number
  defer cancel()

  defer func() {
    if errStr != "" {
      zgo.Http.JsonpErr(ctx, errStr)
    }
  }()

  // 第二：解析请求参数
  name := ctx.URLParam("name")
  if name == "" {
    errStr = "必须输入query参数name"
    return
  }

  key := fmt.Sprintf("%s:%s:%s", "zgo", "start", name)

  // 第三：调用zgo engine来处理业务逻辑
  val, err := zgo.Redis.Get(cotx, key)
  if err != nil {
    errStr = err.Error()
    zgo.Log.Error(err)
    return
  }

  result := zgo.Utils.StringToMap(val.(string))

  // 第四：使用select来响应处理结果与超时
  select {
  case <-cotx.Done():
    errStr = "call redis get string timeout"
    zgo.Log.Error(errStr) //通过zgo.Log统计日志
  default:
    zgo.Http.JsonpOK(ctx, result)
  }

}
