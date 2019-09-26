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
		v1.Get("/mongo/get", handlers.RedisGet)
		v1.Post("/mongo/put", handlers.RedisSet)

	}

}
