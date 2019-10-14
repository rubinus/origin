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

	v1 := app.Party("/v1")
	{
		//这是一个get的例子，可以直接copy或是更改
		v1.Get("/redis/get", handlers.RedisGet)

		//这是一个Post的例子，请按照结构，更改结构体与请求参数
		v1.Post("/pay/do", handlers.DoPay)

	}

}
