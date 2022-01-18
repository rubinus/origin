package models

/*
@Time : 2019-10-14 15:23
@Author : rubinus.chu
@File : pay
@project: origin
*/

type PayRequest struct {
  Bid     string `json:"bid"`
  Aid     uint32 `json:"aid"`
  PayType uint8  `json:"pay_type"` //1微信 2支付宝
  OrderNo string `json:"order_no"` //定单id 32位
  Channel uint8  `json:"channel"`  //对应微信或支付宝的通道

  Amount int         `json:"amount"` //金额，分
  Body   string      `json:"body"`
  Detail interface{} `json:"detail"`

  BuyerId    string `json:"buyer_id"`
  ExpireTime int64  `json:"expire_time"` //过期时间

}
