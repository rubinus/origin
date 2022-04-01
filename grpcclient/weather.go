package grpcclient

import (
	"context"
	"errors"
	pb_weather "github.com/gitcpu-io/origin/pb/weather"
)

/*
@Time : 2022-04-01 11:09
@Author : rubinus.chu
@File : weather
@project: origin
*/

func RpcWeather(ctx context.Context, request *pb_weather.ListRequest) (*pb_weather.ListResponse, error) {
	out := make(chan *pb_weather.ListResponse)
	errCh := make(chan error)
	go func() {
		if WeatherClient == nil {
			errCh <- errors.New("WeatherClient not ready")
			return
		}
		response, err := WeatherClient.List(ctx, request)
		if err != nil {
			errCh <- err
			return
		}
		out <- response
	}()

	select {
	case <-ctx.Done():
		errStr := "RpcWeather timeout"
		return nil, errors.New(errStr)

	case err := <-errCh:
		return nil, err

	case r := <-out:
		return r, nil
	}

}
