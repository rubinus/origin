package routes

/*
@Time : 2019-03-05 18:27
@Author : rubinus.chu
@File : index
@project: origin
*/

import (
  "github.com/gitcpu-io/origin/handlers"
  "github.com/kataras/iris/v12"
)

//前端ajax-->main.go(Run)-->routes-->(实际业务处理handler)-->services-->zgo.组件(mysql/mongo/redis/pika)-->models(库)

func Index(app *iris.Application) {
  app.OnErrorCode(iris.StatusNotFound, handlers.FourZeroFourPage)
  app.OnErrorCode(iris.StatusInternalServerError, handlers.FiveZeroZeroPage)

  app.Get("/", handlers.IndexPage)

  // 不要删除这个路由，这是专门为容器运行在k8s时，提供的探针路由，判断微服务是否健康的
  app.Get("/health", handlers.Health)

  v1 := app.Party("/v1")
  {
    v1.Get("/trace", handlers.TraceGet)

    //这是一个redis get的例子，可以直接copy或是更改
    v1.Get("/redis/get", handlers.RedisGet)

    //这是一个mongo get的例子，可以直接copy或是更改
    v1.Get("/mongo/get", handlers.MongoGet)

    //这是一个mongo list的例子，可以直接copy或是更改
    v1.Get("/mongo/list", handlers.MongoList)

    //这是一个Post的例子，请按照结构，更改结构体与请求参数
    v1.Post("/pay/do", handlers.DoPay)

  }

}
