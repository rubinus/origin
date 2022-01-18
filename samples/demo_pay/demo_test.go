package demo_pay

import (
  "testing"
)

/*
@Time : 2019-10-11 17:43
@Author : rubinus.chu
@File : demo_test
@project: origin
*/

func TestWechat(t *testing.T) {
  //WechatTradeOrder()	//统一下单，生成pre order id
  //WechatMicropay()
  WechatTradeQuery()

  //AliPayOrderPay()
  //ZhimaCreditScoreGet()
  //AliPayTradePrecreate()
  //fmt.Println("------")
  //AliPayTradeCreate()
  //fmt.Println("------")
  //AliPayTradeAppPay()
  //fmt.Println("------")

  AliPayTradeQuery()
  //AliPayTradeWapPay()
  //AliPayTradePagePay()
  //AliPayTradeAppPay()
  //AliPayOpenAuthTokenApp()
  //
  //AliPayFundTransToaccountTransfer()
  //AliPayFundTransOrderQuery()
  //AliPayFundAccountQuery()

  //AliPaySystemOauthToken()
}
