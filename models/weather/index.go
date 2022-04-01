package weather

import (
	"context"
	"errors"
	"fmt"
	"github.com/gitcpu-io/zgo"
)

/*
@Time : 2022-03-31 17:18
@Author : rubinus.chu
@File : weather
@project: origin
*/

type Weatherer interface {
	Insert(ctx context.Context, trade *Weather) (result string, err error)
	List(ctx context.Context, city string) (weathers []*Weather, err error)
}

func New() Weatherer {
	return &Weather{}
}

// Weather Weather映射mongo表结构
type Weather struct {
	Id         string `json:"id,omitempty" bson:"_id,omitempty"`
	City       string `json:"city" bson:"city"`
	Country    string `json:"country" bson:"country"`
	Latitude   string `json:"latitude" bson:"latitude"`
	Longitude  string `json:"longitude" bson:"longitude"`
	TimeZone   string `json:"timeZone" bson:"time_zone"`
	Day        Day    `json:"day" bson:"day"`
	CreateTime int64  `json:"createTime" bson:"createTime"`
	UpdateTime string `json:"updateTime" bson:"updateTime"`
}

type Day struct {
	Phrase                 string `json:"phrase" bson:"phrase"`
	Narrative              string `json:"narrative" bson:"narrative"`
	Temperature            string `json:"temperature" bson:"temperature"`
	TemperatureMaxSince7am string `json:"temperatureMaxSince7am" bson:"temperature_max_since_7_am"`
	WindDirCompass         string `json:"windDirCompass" bson:"wind_dir_compass"`
}

// Insert 保存方法
func (repo *Weather) Insert(ctx context.Context, wea *Weather) (result string, err error) {
	if wea.CreateTime == 0 {
		wea.CreateTime = zgo.Utils.GetTimestamp(10)
	}

	//取db连接
	var collection = zgo.Mongo.GetCollection("weather", "weather", "mongo_label_bj")
	if collection == nil {
		err := errors.New("没有这个db或document")
		return "", err
	}
	result, err = zgo.Mongo.Insert(ctx, collection, wea)
	if err != nil {
		zgo.Log.Error(err)
		return
	}
	fmt.Println("insert mongo done: ", result)
	return
}

// List 查询列表
func (repo *Weather) List(ctx context.Context, city string) (weathers []*Weather, err error) {
	//取db连接
	var collection = zgo.Mongo.GetCollection("weather", "weather", "mongo_label_bj")
	if collection == nil {
		err = errors.New("没有这个db或document")
		return
	}

	filter := make(map[string]interface{}) //查询query
	filter["city"] = city

	sort := make(map[string]interface{})
	sort["createTime"] = -1

	//fields := make(map[string]interface{})
	//fields["city"] = 1

	//组织args
	args := &zgo.MongoArgs{
		Filter: filter, //查询条件
		//Fields: fields, //对查询出的结果项，筛选字段
		Sort:  sort, //排序
		Limit: 10,   //查询结果数量
		Skip:  0,    //从哪一条开始跳过 开区间，不包括skip的值
	}

	results, err := zgo.Mongo.Find(ctx, collection, args)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	for _, v := range results {
		u := Weather{}
		err := zgo.Utils.BsonUnmarshal(v, &u) //对每一条数据进行 bsonUnmarshal 转为go结构体
		if err != nil {
			continue
		}
		weathers = append(weathers, &u)
	}
	return
}
