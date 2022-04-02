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

//前端请求接口-->main.go(Run)-->routes(视图 V 层)-->handlers(控制 C 层)-->services(逻辑 L 层)-->models(库 M 层)-->zgo.engine组件(mysql/mongo/redis/kafka)

func Index(app *iris.Application) {
	app.OnErrorCode(iris.StatusNotFound, handlers.FourZeroFourPage)
	app.OnErrorCode(iris.StatusInternalServerError, handlers.FiveZeroZeroPage)

	app.Get("/", handlers.IndexPage)

	// 不要删除这个路由，这是专门为容器运行在k8s时，提供的探针路由，判断微服务是否健康的
	app.Get("/health", handlers.Health)

	apis := app.Party("/apis/") //apis接口
	{
		// demo: trace
		traceV1 := apis.Party("/trace") //资源
		{
			v1 := traceV1.Party("/v1") //版本
			{
				v1.Get("/do", handlers.TraceGet) //REST api
				//todo add other route

			}
		}
		// demo: redis
		redisV1 := apis.Party("/redis") //资源
		{
			v1 := redisV1.Party("/v1") //版本
			{
				//这是一个redis get的例子，可以直接copy或是更改
				v1.Get("/get", handlers.RedisGet) //REST api
				//todo add other route

			}

		}
		// demo: mongo
		mongoV1 := apis.Party("/mongo") //资源
		{
			v1 := mongoV1.Party("/v1") //版本
			{
				//这是一个mongo get的例子，可以直接copy或是更改
				v1.Get("/get", handlers.MongoGet) //REST api
				//这是一个mongo list的例子，可以直接copy或是更改
				v1.Get("/list", handlers.MongoList) //REST api
				//todo add other route

			}

		}

		// weather 请使用以下路由，更改对应的其它路由及内部详细代码，实现 MVC 模式的面向接口编程
		weatherV1 := apis.Party("/weather") //资源
		{
			v1 := weatherV1.Party("/v1") //版本
			{
				//这是一个典型的 MVC 模式的Post例子，请严格按照结构，更改结构体与请求参数（route->handler->service->models）
				v1.Post("/put", handlers.SaveWeather) //REST api

				//这是一个典型的 MVC 模式的Get 列表的例子
				v1.Get("/list", handlers.ListWeather) //REST api
				//todo add other route

			}
		}

	}

}
