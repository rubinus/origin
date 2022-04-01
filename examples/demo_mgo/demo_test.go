package demo_mgo

import (
	"context"
	"fmt"
	"github.com/gitcpu-io/zgo"
	"testing"
	"time"
)

/*
@Time : 2019-09-17 10:52
@Author : rubinus.chu
@File : demo_test.go
@project: origin
*/

func TestGet(t *testing.T) {
	err := zgo.Engine(&zgo.Options{
		Env:     "dev",
		Project: "1553240759",
		Mgo: []string{
			label_bj,
			label_sh,
		},
	})
	//测试时表示使用mgo，在origin中使用一次
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//++++++++++++++++++++++++++使用原生连接++++++++++++++++++++++++++++
	connChan, err := zgo.Mongo.GetConnChan(label_bj)
	if err != nil {
		zgo.Log.Error(err)
		return
	}
	client := <-connChan
	collection := client.Database("profiles").Collection("bj")
	cursor, err := collection.Indexes().List(context.TODO())
	if err != nil {
		zgo.Log.Error(err)
		return
	}
	var name interface{}
	for cursor.Next(ctx) {
		if cursor != nil {
			err := cursor.Decode(&name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(name, "====")
		}
	}
	fmt.Println("database name ----", collection.Database().Name())
	//++++++++++++++++++++++++++使用原生连接++++++++++++++++++++++++++++

	//###########################################################################
	// 开始使用 zgo Mgo 官方驱动
	//###########################################################################

	fmt.Println("\n-----------开始FindById")
	FindById(ctx, "5d8040af7956ebe482511adf")

	fmt.Println("\n-----------开始FindOne")
	filter := make(map[string]interface{})
	filter["_id"] = "5d821489ada5f1088cb1e82b"
	//filter["username"] = "zhangsan"
	FindOne(ctx, filter)

	fmt.Println("\n-----------开始Find")
	Find(ctx, "zhangsan")

	fmt.Println("\n-----------开始Count")
	Count(ctx, "zhangsan")

	fmt.Println("\n-----------开始Insert")
	Insert(ctx)

	fmt.Println("\n-----------开始InsertMany")
	InsertMany(ctx)

	fmt.Println("\n-----------开始UpdateById")
	UpdateById(ctx, "5d821489ada5f1088cb1e82b")
	FindById(ctx, "5d821489ada5f1088cb1e82b")

	fmt.Println("\n-----------开始UpdateOne")
	filter = make(map[string]interface{})
	filter["_id"] = "5d81e6ee8292b114061251a3"
	UpdateOne(ctx, filter)
	FindOne(ctx, filter)

	fmt.Println("\n-----------开始UpdateMany")
	UpdateMany(ctx, "lisi")
	Find(ctx, "lisi")

	fmt.Println("\n-----------开始DeleteById")
	DeleteById(ctx, "5d81b52118d6954f9809e22b")
	FindById(ctx, "5d81b52118d6954f9809e22b")

	fmt.Println("\n-----------开始DeleteOne")
	Count(ctx, "lisi")
	DeleteOne(ctx, "lisi")
	Count(ctx, "lisi")

	fmt.Println("\n-----------开始DeleteMany")
	Count(ctx, "lisi")
	DeleteMany(ctx, "lisi")
	Count(ctx, "lisi")

	fmt.Println("\n-----------开始ReplaceOne")
	ReplaceOne(ctx, "zhangsan")
	Find(ctx, "zhangsan")

	fmt.Println("\n-------------构造update filter 子文档查询条件-------开始FindOneAndUpdate--------")
	//查询条件Filter
	filter = make(map[string]interface{})
	//filter["_id"] = "5d81e00bada5f1088cb1d236"
	filter["username"] = "zhudaxian" //可以是某字段或_id
	filter["houses"] = map[string]interface{}{
		"$gte": 130, //可以是其它$or、$not、$lt
	}

	//更新项Update
	update := make(map[string]interface{})
	update["$inc"] = map[string]interface{}{
		"age":   100,
		"money": -100,
		//可以有多个字段k,v;
	}
	update["$set"] = map[string]interface{}{
		"address": "FindOneAndUpdate更新到重庆--11",
		"post":    "100002", //更新某字段
		//"houses.$[element]": 411001, //如果houses是纯数组:[xx,xx,xx]
		//子文档的$[element] 其中这个element可以自定义名字
		"grades.$[elem].mean": 100, //如果grades是对象数组:[{k:v,mean:v},{k:v,mean:v}]
		//子文档$[elem]
		//可以有多个字段k,v;但只能有一个顶级字段，意味着$[element]和$[ele]二选一
	}
	type Score struct {
		Grade int `json:"grade"`
		Mean  int `json:"mean"`
	}
	update["$push"] = map[string]interface{}{
		"scores": Score{ //已有一个数组，这里是一个个的push object进数组中
			Grade: 70,
			Mean:  65,
		},
	}

	//子文档ArrayFilters查询条件
	var arrayFilters []map[string]interface{}
	af := make(map[string]interface{})
	//af["element"] = map[string]interface{}{
	//	//这里的element对应上面的houses.$[element]，意思是数组中的每一项元素
	//	"$gte": 134,
	//}
	af["elem.grade"] = map[string]interface{}{ //elem.grade和element二选一
		"$gte": 70,
	}
	af["elem.mean"] = map[string]interface{}{ //但elem中可以有多个elem.xx或elem.yy
		"$gte": 60,
	}
	arrayFilters = append(arrayFilters, af)

	//begin update
	FindOneAndUpdate(ctx, filter, update, arrayFilters)

	fmt.Println("\n-----------开始FindOneAndReplace")
	filter = make(map[string]interface{})
	filter["_id"] = "5d821365ada5f1088cb1e78a"
	FindOneAndReplace(ctx, filter)

	fmt.Println("\n-----------开始FindOneAndDelete")
	FindOneAndDelete(ctx, "zhangsan")

	fmt.Println("\n-----------开始BulkWrite----------------")
	BulkWrite(ctx)
	filter = make(map[string]interface{})
	filter["username"] = "朱大仙儿"
	FindOne(ctx, filter)

	fmt.Println("\n-----------开始Aggregate聚合查询")
	Aggregate(ctx)

	fmt.Println("\n-----------开始Distinct去重查询")
	Distinct(ctx)

	//fmt.Println("\n-----------开始Watch去重查询")
	//Watch(ctx)	//only supported on replica sets

}
