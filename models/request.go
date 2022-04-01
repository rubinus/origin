package models

/*
@Time : 2019-10-14 15:23
@Author : rubinus.chu
@File : pay
@project: origin
*/

type WeatherRequest struct {
	Appid     string `json:"appid"`
	Appsecret string `json:"appsecret"`
	Version   string `json:"version"`
	Query     string `json:"query"`
}
