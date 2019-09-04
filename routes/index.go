package routes

/*
@Time : 2019-03-05 18:27
@Author : rubinus.chu
@File : index
@project: origin
*/

import (
	"git.zhugefang.com/gobase/origin/handlers"
	"github.com/kataras/iris"
)

//前端ajax-->main.go(Run)-->routes-->(实际业务处理handler)-->services-->zgo.组件(mysql/mongo/redis/pika)-->models(库)

func Index(app *iris.Application) {

	app.Get("/", handlers.IndexPage)

	//http://localhost:8080/hello?name=niubi
	app.Get("/hello", handlers.Hello)

	//http://localhost:8080/v1/bj/addone
	v1 := app.Party("/v1")
	{
		v1.Get("/nsq/put", handlers.NsqPut)
		v1.Get("/kafka/put", handlers.KafkaPut)
		v1.Get("/redis/put", handlers.RedisPut)
		v1.Get("/redis/get", handlers.RedisGet)
		v1.Get("/pika/put", handlers.PikaPut)
		v1.Get("/pika/get", handlers.PikaGet)
		v1.Get("/es/{city:string}/esget", handlers.CityareaAggsV1) //es
		v1.Post("/es/{city:string}/esput", handlers.AddDataV1)     //es
		v1.Get("/es/{city:string}/esquery", handlers.EsQueryV1)    //es

		v1.Get("/mysql/{city:string}/get", handlers.MysqlGetV1)  //mysql get
		v1.Post("/mysql/{city:string}/put", handlers.MysqlAddV1) //mysql put

		v1.Get("/mongo/{city:string}/mongoget", handlers.MongoGetV1)   //mongo get
		v1.Post("/mongo/{city:string}/mongoput", handlers.MongoSaveV1) //mongo put

		v1.Get("/redis/hashput", handlers.RedisHashPut) //redis pika hash 操作
		v1.Get("/redis/hashget", handlers.RedisHashGet)
		v1.Get("/pika/hashput", handlers.PikaHashPut)
		v1.Get("/pika/hashget", handlers.PikaHashGet)

	}

}
