package demo_mgo

import (
	"context"
	"fmt"
	"github.com/gitcpu-io/zgo"
)

/*
@Time : 2019-09-17 10:52
@Author : rubinus.chu
@File : demo.go
@project: origin
*/

const (
	label_bj = "mongo_label_bj"
	label_sh = "mongo_label_sh"
)

type User struct {
	Id       zgo.MongoObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Username string            `json:"username" bson:"username" `
	Age      int               `json:"age" bson:"age"`
	Address  string            `json:"address" bson:"address"`
}

////////*******************FindById 按id
func FindById(ctx context.Context, id string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	r := &User{}
	err := zgo.Mongo.FindById(ctx, collection, r, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.Id.IsZero() {
		fmt.Println("没有这个id:", id)
		return
	}
	fmt.Println("成功", r.Id.Hex(), r)
}

////////*******************FindOne 一条
func FindOne(ctx context.Context, filter map[string]interface{}) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	sort := make(map[string]interface{})
	sort["_id"] = -1 //-1降序，1升序

	//返回错误：Projection cannot have a mix of inclusion and exclusion; 要么是1，要么是0
	fields := make(map[string]interface{})
	fields["age"] = 1 //要么全是1，要么全是0
	fields["address"] = 1
	fields["username"] = 1

	r := &User{}

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
		Fields: fields, //对查询出的结果项，筛选字段
		Sort:   sort,   //排序
		Result: r,      //传入 &User{} ,结果
	}

	c, err := zgo.Mongo.FindOne(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	if c > 0 {
		fmt.Println(c, err, r)
	}
}

////////*******************Find  多条
func Find(ctx context.Context, username string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	filter := make(map[string]interface{}) //查询username是且age >= 30的
	filter["username"] = username
	filter["age"] = map[string]interface{}{
		"$gte": 10,
	}

	sort := make(map[string]interface{})
	sort["_id"] = -1

	//返回错误：Projection cannot have a mix of inclusion and exclusion; 要么是1，要么是0
	fields := make(map[string]interface{})
	fields["age"] = 1
	fields["address"] = 1 //要么全是1，要么全是0
	fields["username"] = 1

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
		Fields: fields, //对查询出的结果项，筛选字段
		Sort:   sort,   //排序
		Limit:  10,     //查询结果数量
		Skip:   0,      //从哪一条开始跳过 开区间，不包括skip的值
	}

	results, err := zgo.Mongo.Find(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	var users []*User
	for _, v := range results {
		u := User{}
		err := zgo.Utils.BsonUnmarshal(v, &u) //对每一条数据进行 bsonUnmarshal 转为go结构体
		if err != nil {
			continue
		}
		users = append(users, &u)
	}

	for _, v := range users {
		fmt.Println(v, "====", v.Id.Hex(), v.Username, v.Age)
	}
}

////////*******************Count  数量
func Count(ctx context.Context, username string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	filter := make(map[string]interface{}) //查询username是且age >= 30的
	filter["username"] = username
	filter["age"] = map[string]interface{}{
		"$gte": 30,
	}

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
		Limit:  0,      //查询结果数量 0表示所有数据 可选
		Skip:   0,      //从哪一条开始跳过 开区间，不包括skip的值 可选
	}

	results, err := zgo.Mongo.Count(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	fmt.Println("count=", username, results)

}

////////*******************Insert  保存
func Insert(ctx context.Context) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	document := &User{
		Username: "zhangsan",
		Age:      34,
		Address:  "上海",
	}

	result, err := zgo.Mongo.Insert(ctx, collection, document)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	fmt.Println("insert id=", result)

}

////////*******************InsertMany  保存多条
func InsertMany(ctx context.Context) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	document1 := &User{
		Username: "zhangsan",
		Age:      28,
		Address:  "上海",
	}
	document2 := &User{
		Username: "lisi",
		Age:      30,
		Address:  "北京",
	}

	var users []interface{}
	users = append(users, document1)
	users = append(users, document2)

	results, err := zgo.Mongo.InsertMany(ctx, collection, users)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	fmt.Println("insertMany ids=", results)

}

////////*******************UpdateById  通过Id更新一条
func UpdateById(ctx context.Context, id string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	update := make(map[string]interface{})
	update["$set"] = map[string]interface{}{
		"address": "UpdateById更新到大连",
		//可以有多个字段k,v
	}
	update["$inc"] = map[string]interface{}{
		"age": 1000,
		//可以有多个字段k,v
	}

	byId, err := zgo.Mongo.UpdateById(ctx, collection, update, id)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	fmt.Println("UpdateById=", id, byId)

}

////////*******************UpdateOne  更新一条
func UpdateOne(ctx context.Context, filter map[string]interface{}) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	//filter := make(map[string]interface{}) //查询username是且age >= 25的
	//filter["username"] = username
	//filter["age"] = map[string]interface{}{
	//	"$gte": 25,
	//}

	update := make(map[string]interface{})
	update["$set"] = map[string]interface{}{
		"address": "UpdateOne更新到上海--11$set",
		//可以有多个字段k,v
	}
	update["$inc"] = map[string]interface{}{
		"age": 110,
		//可以有多个字段k,v
	}

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
		Update: update, //更新字段
		Upsert: true,   //查询不到时，是否插入
	}

	u, i, s, err := zgo.Mongo.UpdateOne(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	fmt.Println("updateOne=", filter, u, i, s)

}

////////*******************ReplaceOne  替换一条
func ReplaceOne(ctx context.Context, username string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	filter := make(map[string]interface{}) //查询username是且age >= 25的
	filter["username"] = username

	update := make(map[string]interface{})
	update["address"] = "ReplaceOne替换到南京"
	update["username"] = username

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
		Update: update, //替换字段
		Upsert: true,   //查询不到时，是否插入
	}

	u, i, s, err := zgo.Mongo.ReplaceOne(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	fmt.Println("ReplaceOne=", username, u, i, s)

}

////////*******************UpdateMany  更新多条
func UpdateMany(ctx context.Context, username string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	filter := make(map[string]interface{}) //查询username是且age >= 25的
	filter["username"] = username
	filter["age"] = map[string]interface{}{
		"$gte": 25,
	}

	update := make(map[string]interface{})
	update["$set"] = map[string]interface{}{
		"address": "UpdateMany更新到深圳--11$set",
		//可以有多个字段k,v
	}
	update["$inc"] = map[string]interface{}{
		"age": 110,
		//可以有多个字段k,v
	}

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
		Update: update, //更新字段
		Upsert: true,   //查询不到时，是否插入
	}

	u, i, i2, err := zgo.Mongo.UpdateMany(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	fmt.Println("updateMany=", username, u, i, i2)

}

////////*******************DeleteById  删除多条
func DeleteById(ctx context.Context, id string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	deleteById, err := zgo.Mongo.DeleteById(ctx, collection, id)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	fmt.Println("deleteById=", id, deleteById)

}

////////*******************DeleteOne  删除多条
func DeleteOne(ctx context.Context, username string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	filter := make(map[string]interface{}) //查询username是且age >= 29的
	filter["username"] = username
	filter["age"] = map[string]interface{}{
		"$gte": 30,
	}

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
	}

	deleteOne, err := zgo.Mongo.DeleteOne(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	fmt.Println("deleteOne=", username, deleteOne)

}

////////*******************DeleteMany  删除多条
func DeleteMany(ctx context.Context, username string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	filter := make(map[string]interface{}) //查询username是且age >= 25的
	filter["username"] = username
	filter["age"] = map[string]interface{}{
		"$gte": 30,
	}

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
	}

	deleteMany, err := zgo.Mongo.DeleteMany(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	fmt.Println("deleteMany=", username, deleteMany)

}

////////*******************FindOneAndUpdate 一条
func FindOneAndUpdate(ctx context.Context, filter, update map[string]interface{}, arrayFilters []map[string]interface{}) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	sort := make(map[string]interface{})
	sort["_id"] = -1 //-1降序，1升序

	//返回错误：Projection cannot have a mix of inclusion and exclusion; 要么是1，要么是0
	fields := make(map[string]interface{})
	fields["age"] = 1 //要么全是1，要么全是0
	fields["address"] = 1
	fields["username"] = 1

	r := &User{}

	//组织args
	args := &zgo.MongoArgs{
		Filter:       filter, //查询条件
		ArrayFilters: arrayFilters,
		Fields:       fields, //对查询出的结果项，筛选字段
		Update:       update, //更新的
		Sort:         sort,   //排序
		Result:       r,      //传入 &User{} ,结果
		Upsert:       true,   //查询不到时，是否插入
	}

	err := zgo.Mongo.FindOneAndUpdate(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	fmt.Println("FindOneAndUpdate=", filter, r)
}

////////*******************FindOneAndReplace 一条
func FindOneAndReplace(ctx context.Context, filter map[string]interface{}) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	sort := make(map[string]interface{})
	sort["_id"] = 1 //-1降序，1升序

	//返回错误：Projection cannot have a mix of inclusion and exclusion; 要么是1，要么是0
	fields := make(map[string]interface{})
	fields["age"] = 1 //要么全是1，要么全是0
	fields["address"] = 1
	fields["username"] = 1

	update := make(map[string]interface{}) //替换后，没有age字段
	update["address"] = "FindOneAndReplace替换到石家庄"

	r := &User{}

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
		Fields: fields, //对查询出的结果项，筛选字段
		Update: update, //替换项
		Sort:   sort,   //排序
		Result: r,      //传入 &User{} ,结果
		Upsert: true,   //查询不到时，是否插入
	}

	err := zgo.Mongo.FindOneAndReplace(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	fmt.Println("FindOneAndReplace=", filter, r)
}

////////*******************FindOneAndDelete 一条
func FindOneAndDelete(ctx context.Context, username string) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	filter := make(map[string]interface{})
	filter["username"] = username

	sort := make(map[string]interface{})
	sort["_id"] = 1 //-1降序，1升序

	//返回错误：Projection cannot have a mix of inclusion and exclusion; 要么是1，要么是0
	fields := make(map[string]interface{})
	fields["age"] = 1 //要么全是1，要么全是0
	fields["address"] = 1
	fields["username"] = 1

	r := &User{}

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
		Fields: fields, //对查询出的结果项，筛选字段
		Sort:   sort,   //排序
		Result: r,      //传入 &User{} ,结果
	}

	err := zgo.Mongo.FindOneAndDelete(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	fmt.Println("FindOneAndDelete=", username, r)
}

////////*******************BulkWrite并行多条操作
func BulkWrite(ctx context.Context) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	var bulkWrites []*zgo.MongoBulkWriteOperation

	//******************************************************insertOne operation
	document := &User{
		Username: "朱大仙儿",
		Age:      zgo.Utils.RandRangeInt(20, 30),
		Address:  "北京",
	}
	bwInsertOne := &zgo.MongoBulkWriteOperation{
		Operation: zgo.MongoBWOInsertOne,
		MongoArgs: &zgo.MongoArgs{
			Document: document,
		},
	}
	bulkWrites = append(bulkWrites, bwInsertOne)

	//******************************************************insertOne operation again
	document2 := &User{
		Username: "朱大仙儿",
		Age:      zgo.Utils.RandRangeInt(20, 30),
		Address:  "北京 朝阳",
	}
	bwInsertOne2 := &zgo.MongoBulkWriteOperation{
		Operation: zgo.MongoBWOInsertOne,
		MongoArgs: &zgo.MongoArgs{
			Document: document2,
		},
	}
	bulkWrites = append(bulkWrites, bwInsertOne2)

	//******************************************************updateOne operation
	bwUpdateOne := &zgo.MongoBulkWriteOperation{
		Operation: zgo.MongoBWOUpdateOne,
		MongoArgs: &zgo.MongoArgs{
			Filter: map[string]interface{}{
				"_id": "5d81ec82ada5f1088cb1d77c",
				//"username": "zhudaxian",
			},
			Update: map[string]interface{}{
				"$set": map[string]interface{}{
					"address": "bulkUpdate",
					"grades":  []int{20, 30, 50},
				},
				"$inc": map[string]interface{}{
					"age": zgo.Utils.RandRangeInt(20, 30),
				},
			},
		},
	}
	bulkWrites = append(bulkWrites, bwUpdateOne)

	//******************************************************updateOne operation again
	bwUpdateOne2 := &zgo.MongoBulkWriteOperation{
		Operation: zgo.MongoBWOUpdateOne,
		MongoArgs: &zgo.MongoArgs{
			Filter: map[string]interface{}{
				//"_id": "5d81ec82ada5f1088cb1d77c",
				"username": "没有这个用户时直接插入:" + zgo.Utils.RandomString(4),
			},
			Update: map[string]interface{}{
				"$set": map[string]interface{}{
					"address": "更新时upsert=true",
					"grades":  []int{20, 30, 50},
				},
				"$inc": map[string]interface{}{
					"age": zgo.Utils.RandRangeInt(20, 30),
				},
			},
			Upsert: true,
		},
	}
	bulkWrites = append(bulkWrites, bwUpdateOne2)

	//******************************************************replaceOne operation
	bwReplaceOne := &zgo.MongoBulkWriteOperation{
		Operation: zgo.MongoBWOReplaceOne,
		MongoArgs: &zgo.MongoArgs{
			Filter: map[string]interface{}{
				"_id": "5d832d133c4da8b1f1580f6f",
				//"username": "没有这个用户时直接插入",
			},
			Update: map[string]interface{}{
				"$set": map[string]interface{}{
					"address": "replaceOne",
				},
				"$inc": map[string]interface{}{
					"age": zgo.Utils.RandRangeInt(20, 30),
				},
			},
		},
	}
	bulkWrites = append(bulkWrites, bwReplaceOne)

	//******************************************************deleteOne operation
	bwDeleteOne := &zgo.MongoBulkWriteOperation{
		Operation: zgo.MongoBWODeleteOne,
		MongoArgs: &zgo.MongoArgs{
			Filter: map[string]interface{}{
				"_id": "5d8332725bcdd1230bf4d007",
				//"username": "没有这个用户时直接插入",
			},
		},
	}
	bulkWrites = append(bulkWrites, bwDeleteOne)

	//******************************************************updateMany operation
	bwUpdateMany := &zgo.MongoBulkWriteOperation{
		Operation: zgo.MongoBWOUpdateMany,
		MongoArgs: &zgo.MongoArgs{
			Filter: map[string]interface{}{
				//"_id": "5d8332725bcdd1230bf4d007",
				"username": "朱大仙儿",
			},
			Update: map[string]interface{}{
				"$set": map[string]interface{}{
					"address": "beijing",
					"tag":     "bj",
				},
			},
			Upsert: true,
		},
	}
	bulkWrites = append(bulkWrites, bwUpdateMany)

	//******************************************************deleteMany operation
	/**
	  bwDeleteMany := &zgo.MongoBulkWriteOperation{
	  	Operation: zgo.MongoBWODeleteMany,
	  	MongoArgs: &zgo.MongoArgs{
	  		Filter: map[string]interface{}{
	  			//"_id": "5d8332725bcdd1230bf4d007",
	  			"username": "zhudaxian",
	  		},
	  	},
	  }
	  bulkWrites = append(bulkWrites, bwDeleteMany)
	*/

	bulkWriteResult, err := zgo.Mongo.BulkWrite(ctx, collection, bulkWrites, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v-----------BulkWrite------\n", bulkWriteResult)

}

////////*******************Aggregate 聚合查询
func Aggregate(ctx context.Context) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	var pipeline []interface{}

	//match 用于过滤数据，只输出符合条件的文档，相当于Filter入参
	match := make(map[string]interface{})
	match["$match"] = map[string]interface{}{
		"username": "朱大仙儿",
	}

	//group 将集合中的文档分组，可用于统计结果
	group := make(map[string]interface{})
	group["$group"] = map[string]interface{}{
		"_id": "$username",
		"total": map[string]interface{}{
			"$sum": "$age",
		},
	}

	//limit 用来限制聚合管道返回的文档数
	limit := make(map[string]interface{})
	limit["$limit"] = 50

	//sort 将输入文档排序后输出
	sort := make(map[string]interface{})
	sort["$sort"] = map[string]interface{}{
		"total": -1,
	}

	pipeline = append(pipeline, match)
	pipeline = append(pipeline, group)
	pipeline = append(pipeline, sort)
	pipeline = append(pipeline, limit)

	results, err := zgo.Mongo.Aggregate(ctx, collection, pipeline)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	type Total struct {
		Id    string `json:"_id" bson:"_id"`
		Total int    `json:"total"`
	}
	var total []*Total //指定接受的结构体
	for _, v := range results {
		t := Total{}
		err := zgo.Utils.BsonUnmarshal(v, &t) //对每一条数据进行 bsonUnmarshal 转为go结构体
		if err != nil {
			continue
		}
		total = append(total, &t)
	}

	for _, v := range total {
		fmt.Println(v, "====")
	}
}

////////*******************Distinct 去重查询
func Distinct(ctx context.Context) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	filter := make(map[string]interface{})
	//filter["age"] = map[string]interface{}{
	//	"$gte": 30,
	//}
	filter["username"] = "朱大仙儿"

	results, err := zgo.Mongo.Distinct(ctx, collection, "age", filter)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	for _, v := range results {
		fmt.Println(v)
	}

}

////////*******************Watch 监听
func Watch(ctx context.Context) {
	var collection = zgo.Mongo.GetCollection("profiles", "bj", label_bj)

	var pipeline []interface{}
	match := make(map[string]interface{})
	match["$match"] = map[string]interface{}{
		"username": "朱大仙儿",
	}

	pipeline = append(pipeline, match)

	changeStream, err := zgo.Mongo.Watch(ctx, collection, pipeline)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	u := User{}
	for changeStream.Next(ctx) {
		err := changeStream.Decode(&u)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(u)
	}

}
