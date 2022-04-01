package grpchandlers

import (
  "context"
  pb_weather "github.com/gitcpu-io/origin/pb/weather"
  "github.com/gitcpu-io/origin/services"
  "github.com/gitcpu-io/zgo"
  "log"
)

// WeatherServer 可以起名为你的 xxxxServer
type WeatherServer struct{}

// List implements func
func (s *WeatherServer) List(ctx context.Context, request *pb_weather.ListRequest) (*pb_weather.ListResponse, error) {
  log.Printf("Received: Name %v", request.City)

  // 第一：定义返回变量
  var response = &pb_weather.ListResponse{}

  //第二：准备输入参数
  city := request.City

  // 第三：调用zgo engine来处理业务逻辑
  wea := services.NewWeather()
  res, err := wea.List(ctx, city)
  if err != nil {
    zgo.Log.Error(err)
    return response, nil
  }

  // 第四：构建返回值
  var listdata []*pb_weather.ListData
  for _, val := range res {
    r := &pb_weather.ListData{}
    day := &pb_weather.Day{
      Phrase:                 val.Day.Phrase,
      Narrative:              val.Day.Narrative,
      Temperature:            val.Day.Temperature,
      TemperatureMaxSince7Am: val.Day.TemperatureMaxSince7am,
      WindDirCompass:         val.Day.WindDirCompass,
    }

    r.Day = day
    r.Id = val.Id
    r.City = val.City
    r.Country = val.Country
    r.Latitude = val.Latitude
    r.Longitude = val.Longitude
    r.CreateTime = val.CreateTime

    listdata = append(listdata, r)
  }

  response.Data = listdata

  return response, nil

}
