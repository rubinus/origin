package demo_mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/gitcpu-io/zgo"
	"time"
)

type User struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Age   int    `json:"age"`
}

func GetUser() {
	u := User{}
	//输入参数：上下文ctx，args具体的查询操作参数
	args := make(map[string]interface{})
	query := make(map[string]interface{})
	query["name"] = "abc"

	args["db"] = "test"
	args["table"] = "user"
	args["query"] = query
	args["obj"] = &u

	res, err := getUser(args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func getUser(args map[string]interface{}) (*User, error) {
	//还需要一个上下文用来控制开出去的goroutine是否超时
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	result := zgo.Mongo.FindOne(ctx, args)

	//if err != nil {
	//	zgo.Log.Error("取mongo中的数据失败" + err.Error())
	//	return nil,err
	//}

	select {
	case <-ctx.Done():
		return nil, errors.New("超时")
	default:
		//zgo.Utils.MapToStruct(result, u)
		fmt.Println(result)
	}
	return nil, nil
}

func UpdateNameById() error {
	//还需要一个上下文用来控制开出去的goroutine是否超时
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var args = make(map[string]interface{})
	var update = make(map[string]interface{})
	update["name"] = "dddddddd"
	update["age"] = 100
	args["db"] = "test"
	args["table"] = "user"
	args["update"] = update

	err := zgo.Mongo.UpdateById(ctx, "5cf4f2de8896b853ef722d52", args)

	if err != nil {
		zgo.Log.Error("更新failed" + err.Error())
		return err
	}

	select {
	case <-ctx.Done():
		return errors.New("超时")
	default:
		//zgo.Utils.MapToStruct(result, u)
	}
	return nil
}

func DeleteById() error {
	//还需要一个上下文用来控制开出去的goroutine是否超时
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var args = make(map[string]interface{})
	args["db"] = "test"
	args["table"] = "user"
	err := zgo.Mongo.DeleteById(ctx, "5c93706fd42527dbfafc1121", args)

	if err != nil {
		zgo.Log.Error("更新failed" + err.Error())
		return err
	}

	select {
	case <-ctx.Done():
		return errors.New("超时")
	default:
		//zgo.Utils.MapToStruct(result, u)
	}
	return nil
}

func Delete() error {
	//还需要一个上下文用来控制开出去的goroutine是否超时
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var query = make(map[string]interface{})
	query["name"] = "测试更新3333"
	var args = make(map[string]interface{})
	args["db"] = "test"
	args["table"] = "user"
	args["query"] = query
	err := zgo.Mongo.Delete(ctx, args)

	if err != nil {
		zgo.Log.Error("更新failed" + err.Error())
		return err
	}

	select {
	case <-ctx.Done():
		return errors.New("超时")
	default:
		//zgo.Utils.MapToStruct(result, u)
	}
	return nil
}
