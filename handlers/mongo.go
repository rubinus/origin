package handlers

import (
	"context"
	//"fmt"

	"git.zhugefang.com/gocore/zgo"
	"github.com/kataras/iris"
)

//var uid int64 = 1

type User struct {
	UserId int64  `bson:"user_id" json:"user_id"`
	Name   string `bson:"name" json:"name"`
}

func MongoGetV1(ctx iris.Context) {
	city := ctx.Params().GetStringDefault("city", "")
	if city == "" {
		zgo.Http.JsonpErr(ctx, "city is empty")
		return
	}

	userId := ctx.Params().GetIntDefault("userId", -1)
	userId = 10
	if userId == -1 {
		zgo.Http.JsonpErr(ctx, "userId is null")
		return
	}

	var args = make(map[string]interface{}, 0)
	//var obj User
	var obj interface{}

	args["db"] = "test"
	args["table"] = "user"

	args["query"] = map[string]interface{}{
		"user_id": 100,
	}
	args["obj"] = &obj
	//fmt.Println(args)

	err := zgo.Mongo.FindOne(context.TODO(), args)
	if err != nil {
		zgo.Http.JsonpErr(ctx, "mongo query error:"+err.Error())
		return
	}

	zgo.Http.JsonpOK(ctx, obj)
	return
}

func MongoSaveV1(ctx iris.Context) {
	city := ctx.Params().GetStringDefault("city", "")
	if city == "" {
		zgo.Http.JsonpErr(ctx, "city is null")
		return
	}

	//var obj = map[string]interface{}{
	//	"name":    "Jim King",
	//	"user_id": 1,
	//}
	var obj = &User{
		Name:   "Jim King",
		UserId: 100,
	}

	var args = make(map[string]interface{}, 0)
	args["db"] = "test"
	args["table"] = "user"
	args["items"] = obj

	err := zgo.Mongo.Insert(context.TODO(), args)

	if err != nil {
		zgo.Http.JsonpErr(ctx, "mongo insert error:"+err.Error())
		return
	}

	zgo.Http.JsonpOK(ctx, "ok")
	return
}
