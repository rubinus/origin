package services

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/origin/models"
  "github.com/gitcpu-io/origin/models/weather"
  "github.com/gitcpu-io/zgo"
)

/*
@Time : 2022-03-31 19:02
@Author : rubinus.chu
@File : index
@project: origin
*/

type Weatherer interface {
  Insert(ctx context.Context, req *models.WeatherRequest) (*weather.Weather, error)
  // 请在此处添加其它方法
  List(ctx context.Context, city string) (weathers []*weather.Weather, err error)
}

func NewWeather() Weatherer {
  return &svc{
    repo: weather.New(),
  }
}

type svc struct {
  repo weather.Weatherer
}

// List 查询天气列表
func (svc *svc) List(ctx context.Context, city string) (weathers []*weather.Weather,err error) {
  //todo something
  weathers, err = svc.repo.List(ctx, city)
  if err != nil {
    return
  }
  return
}

// Insert 保存方法
func (svc *svc) Insert(ctx context.Context, req *models.WeatherRequest) (*weather.Weather, error) {

  //todo something

  //请求天气的公共接口
  res, err:= svc.dealRequestWeather(ctx,req)
  if err != nil {
    return nil,err
  }
  //反序列化
  wea := &weather.Weather{}
  err = zgo.Utils.Unmarshal(res, wea)
  if err != nil {
    return nil, err
  }

  //保存
  id, err := svc.repo.Insert(ctx, wea)
  if err != nil {
    return nil, err
  }
  wea.Id = id
  return wea, nil
}

func (svc *svc) dealRequestWeather(ctx context.Context,req *models.WeatherRequest) ([]byte, error) {
  // 天气公共接口 "https://v0.tianqiapi.com/?version=today&query=深圳&appid=43656176&appsecret=I42og6Lm"

  //补充完appid和appsecret
  req.Appid = "43656176"
  req.Appsecret = "I42og6Lm"
  req.Version = "today"

  //转化请求参数
  structToMap := zgo.Utils.StructToMap(req)
  smap := make(map[string]string)
  for idx, val := range structToMap {
    smap[idx] = val.(string)
  }
  wquery := zgo.Utils.GetUrlFormedMap(smap)

  //构造请求地址
  reqHost := fmt.Sprintf("%s?%s", "https://v0.tianqiapi.com/", wquery)

  res, err := zgo.Http.Get(reqHost)
  if err != nil {
    return nil, err
  }
  return res, err
}

//todo add other func
