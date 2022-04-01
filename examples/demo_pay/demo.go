package demo_pay

import (
	"fmt"
	"github.com/gitcpu-io/zgo"
	"strconv"
	"time"
)

/*
@Time : 2019-10-11 15:30
@Author : rubinus.chu
@File : demo
@project: origin
*/

const (
	appId  = "wx181888614b6a225e"
	mchId  = "1520244611"
	apiKey = "MK0FJ5WOVGNaiDHabmgbueDDSCHht7lb"

	aliAppId = "2016120603916259"
	////支付私匙，必须和添加到支付宝开放平台的公匙相对应
	privateKey = "MIIEogIBAAKCAQEAl9MMaBnxszRho8v6xD2VMqHEjQdt8NLXqQ6hoytkQp0FeZoYtSARAa/NhxgmeZAswdZQAOQ4e4bAqrmPwBPQZjekjMKhBIgQPfrfCDOrHIAH35amut7y54TENWbUMIorbI/TwIXDMmXkCigTAgY6IAWClCfhNT9ksOvnog++GDPFNim7U2I7iR7epy34WpXT1V2/2D506G0xPh1fnkmtgc1mqLrb5Sb1oaKTwXsUJRAqY3OEw+sUJOOaIRzYESEUR7PPgnNOSsrKOrpPiQAcp7T1EKHIq5ya2CnFHYkkV891XwkP7d+RX2OGky4QVC4eaQRdVZBvzQk3yqizPUarawIDAQABAoIBACl5ductRzCsO4wSWfOn2w0U0euwvuDNyCofnBpF7UKhQHCinuND2kF6tAuWllQZBZECKaLEtYVRH2rD/Df9ca2qv6HQPUpOnRlBYhIMg07qzrvOnMdpxjNmum7YI1kLNaeXdsIeCF/JI81+ewrPhqtetfghGM2B/tCx8Pn9kPFtL1i4caVS+vjmhQILLYZ/9W1TTUTuLGWlCMGbrXvG7Xsu+zQvEQYObUHG8E5p3Y4eTzqGWINLkQdi5ttRihnJlxQVmPn+ZaNbI/OETYcDZCNGo8qo0sFA9kMtRvkYbprbrqAqqc9oWMn3m9rlg2kXL8vFqkt0KL1HCyTqYJRysDkCgYEAxi0hOu28zA3GUhM5xzb6oW7kY7dMNbG81qIZ9bzrzIAbkJ9uPFXKGJkXYKWOzf2PjWoJU7CuwSn419Ej0EsS6E39wMJfEWAUdZ6MM6F8d/Hbhxz0Crm3PDBw6q4nXePU0RKevPQSAyrCvmOxVfqbJZ4PFTKje/C81xe3l2MeW3cCgYEAxB+ks57R8eSYiirooUdVTZRDpK4PPb5PtrOvze0Hb6PtfLTUYK8vvqT9RzHMYxiluql1lnKnUi4YWyAqJHQNrFg/kqrA5IlaAo0piFh3JMeKbM274MbfnTK/Cu7TW9G07QbcCEZ83KH0ZHSBK61m+X56nyGs0Ur9/1JMwpWTBK0CgYBkx66OeTf5zUd9lalcVek/D1W2IBDxfWG7BNSO03RWmw9pjKpSpI4R1Ei/LXJh7wCBudrkZIV6Vg3mMsUzt/n1iTPyQuZ1v5an+ejymLzjgmtRWrgAfFFimn/R2J7McIBZkk6HaNeAJM7EY45bWwZrKuYgsY/SB6sA91617gISIQKBgGU1Xy1CXF1T9lP26K2xvheW14F4QW0/dKt39tVx6FB4a+na924deaoMQpgm3q6U3ZGCOag16prCJtd/tb+yFAxITiZF79z/9G80oaY24OUeBBU55iaM+oCI24Ws0W2kvpVC4PF0TqMdAJ2GcdI0XduKpYTRP/mRlZnYwrv3rDLRAoGAWGftxlrEY2JSZzt0HNZmkMaupEzW3EmTA9vUXg1exJosBeRjwJOKKHLIboxZNz/Mec+fJLqwbbS3NxAyGQTGISrtBe+sY+/VKWSVuTCHzgOxcCLGLUPSS9EkPo9k2tv1MQsM3ZXs8JbFjsDkSGOJOuPNyT83E/egZArPYmSskBg="
	publicKey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAno07SCE4XlaP5bxyeLjK/GSx0qKLUQM50eYk3RDgcwzNe8fh1pALUk3Z04fWEHKrpkwWj9ihuD/D4hcB3S94tEawz8dZb5gvLSDcE+Jrdx8DjIzk+L2czHwlXBR7fYw7TVZb/4y8+gKwFMJLTp4AwQjvPftllMgJXykp30mGxQH20pCQ7/18oUs+oIAmFibAgBnGYBf6nxlw/mLxgXrlB3rPnBLq6KoMs7Xu0zTu2XzJYK/EPHxfIdWdVf+WzC9LvthjDLaM27RKb2YHUI5gl8HNyU6VOhJONrj9Ma0CMyBqyLuKaJuDX3bI/PkZ7yerK4ojGNXTRjmAUTABt4jysQIDAQAB"

	//aliAppId = "2019101968518085"
	//支付私匙，必须和添加到支付宝开放平台的公匙相对应
	//publicKey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlMB8A94jvUUgjNTKgWh1nlfZIDqzUp1WWgEZwAW5xWmRhNB3gG/rPqwk6v5Qwlq7txeep/P/NNL/BH9Idp7NqLVqzIrVR6FwcFWZpMwhAxihF/A6rSKAii6bOmHbBkOglBR6L0OCLmvg1VnaQbdYDHmxMT2GFMlsf5/nGVyVCIC+9hPwX71NhJ/q00VsT7Ya9s4dvLs6g4CT+gquKqzOZvMOmKXV6OBTrJZ9RXeDAU0EceICNudm6Dp7pglaJlvbQQrL7ik6b98KXpS7hBMaoMHVDXJJMUNCU1qED7LFlBTEMYCO5NI9RZPDWYROEYwuFNA9gPTjr+gh8mQAseY5xwIDAQAB"
	//privateKey = "MIIEowIBAAKCAQEAkcuAx19+22dCpB61YJx/U+e8LBU8KIPKw5Ciju4UZGnFDsuj4DwrqNrRDn9r3cDWP+JP4m2PwQtNOW8nTFy4C99LaFG1dF26pe7S/xMehbdMRibEIhtXEb1G7oYEkmsKodgDSURC6wShSeUaz/I33upO/VlvTv1GaZTc7RRLDciAe+HiJ8La+P57E6vd4/uyx6jBV1ezZntr8Y2wxgoj70xnvC//OchesJztx6saZq0IkSXtKcMWdsn0OzfjXwjMS3g5cm7FSo3L9QMRJiJDxbhtiU2iH8oeLDprw+50TsvLj9b2Hh4IJgZk0VBzF2/jK9wtre//DHpq0kdrvmi4jwIDAQABAoIBABTgyJNEeJ0gv2lTQHQSVSWF1OuXKuM+ZEX3K2A+dcsfXmnM/a250CLBGxjxZFgAKm5BkWECgYiKfsePxFfqGy0QD/NjaBG+7mCev7ZpXYCWjCSrnoCn26MdsM+tf3AcRgyBK49NaCvRoOs9FMbcj+WrNh6PtSHhoTizaaPFuS3C5iMy5mNASPiHBZi0zu4zLrf+H/5UjDF8/7QPdPBBZNRpmfW9bCSQmQrWiMcedlJXd4pfQSbw4Mkjb55fgYdQuhLe9oxybttbn2gfp7qEn6hc9zgzaDiPJlWwSEL7RF7KnZyyhH1u6PjIK1D0m/QupoShkiTGQ9UN1LyWxuOiMSECgYEA0lQFioJl0/hvqFGZ4Tb/suz9LAi/StDtnhfQ8NIBqco2j6gDPaXf7lH6VuhB+zZCHe2xSyIX8/XH662GpFjcSP4WLtaztKEewYE1TMUkWbnAVjGpmbkKUESS4U0DvjG29ztDuKLIUU4FNhuWottzlIxQXyCZJ7yDICCJYgEVbfECgYEAsXQgthOn0P9plgSyxVtI86EJeRbvMQTQWyyMMh4j/3Skaewz5/Uo2hzoYDfIjdUSIZtxC5a9EEg1wiDFBm2TgJzeamxZnw508G+82+opQOYePJJaLlFvco6FlOeWCD6pEm7QCWkFSfXJsnwCLuui+zNjXD/qEZ+EPZ3swGEEDn8CgYBg5GIVF9MvHkDZ6pWYAc0zsSdEVNdC8RK2BMS1XBl2DXirHzw29yY81LkdtitHPgyhWvGU6iTctzodIThol6MLYTn44+GvcZYIkKxsLFl9mCu5yXEXJv5QUfbUIbV6tc5TAJNHCH59rhKKhZUUe0I4iZcw64SCoL5LW5HOey9TEQKBgDii/NHAIj6lVljINRqiUP1ZN4HLXRpDFBEVfcV9MzYUT2lzNvngmGJM+anEBCGokLnjN8hgGwW4VlgYR2oOzRYuexpybIREg/Q9ZYS3DuWkzJ++gkPoP+7LKD1nUM5e2W2FqqZmO4boiLCLvdKl6IXOV/cYeyeWxwk3f1nDXR0XAoGBAK3Yqce+KiO8qlge2fqaFXQxVeHNUHXR3ISWJjnI8Ltqhhse7U3x+cW/cuC81dOdZLUvrMB9DmPYg3zPFfi4SQxdS2Ug6i9ksXf5QIXmOkch3OwTUB/SH9CAEL2f7nSfBWVrbSirO9LMGeMALmSVSFhFdfkbJKgzZDS0w/sCyiAV"
)

var aliPayer zgo.AliPayer
var wechatPayer zgo.WechatPayer

func init() {
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	aliPayer = zgo.AliPay.Pay(aliAppId, privateKey, true)
	aliPayer.SetNotifyUrl("http://testpaybffp.example.com/v1/pay/alipaynotify")

	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	wechatPayer = zgo.Wechat.Pay(appId, mchId, apiKey, true)
	wechatPayer.SetCountry(1)
	wechatPayer.SetNotifyUrl("http://testpaybffp.example.com/v1/pay/wechatnotify")
}

func WechatTradeOrder() {

	//设置国家
	//wechatPayer.SetCountry(1)

	number := zgo.Utils.RandomString(32)
	fmt.Println("out_trade_no:", number)
	//初始化参数Map
	body := make(zgo.BodyMap)
	body.Set("nonce_str", zgo.Utils.RandomString(32))
	body.Set("body", "JSAPI支付")
	body.Set("out_trade_no", number)
	body.Set("total_fee", 1)
	body.Set("spbill_create_ip", "127.0.0.1")
	body.Set("trade_type", "JSAPI")
	body.Set("device_info", "WEB")

	body.Set("openid", "oZOXcwUgI4vRzBd0Ieb6D9w8z07s")

	//请求支付下单，成功后得到结果
	wxRsp, err := wechatPayer.Order(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("wxRsp: %+v\n\n", wxRsp)
	//fmt.Println("wxRsp.MwebUrl:", wxRsp.MwebUrl)

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	//获取小程序支付需要的paySign
	paySign := wechatPayer.GetMiniPaySign(appId, wxRsp.NonceStr, wxRsp.PrepayId, "", timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	fmt.Println("minapp paySign:", paySign)

	//获取H5支付需要的paySign
	paySign_h5 := wechatPayer.GetH5PaySign(appId, wxRsp.NonceStr, wxRsp.PrepayId, "", timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	fmt.Println("h5 paySign:", paySign_h5)

	//获取app需要的paySign
	paySign_app := wechatPayer.GetAppPaySign(appId, mchId, wxRsp.NonceStr, wxRsp.PrepayId, "", timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	fmt.Println("app paySign:", paySign_app)
	fmt.Println("--------------*********-------------")
	fmt.Println("--------------*********-------------")
}

func WechatMicropay() {
	//初始化参数Map
	body := make(zgo.BodyMap)
	body.Set("nonce_str", zgo.Utils.RandomString(32))
	body.Set("body", "扫用户付款码支付")
	number := zgo.Utils.RandomString(32)
	fmt.Println("out_trade_no:", number)

	body.Set("out_trade_no", number)
	body.Set("total_fee", 1)
	body.Set("spbill_create_ip", zgo.Utils.GetIntranetIP())
	body.Set("auth_code", "134595229789828537")

	sign := wechatPayer.GetParamSign(appId, mchId, apiKey, body)

	body.Set("sign", sign)
	//请求支付，成功后得到结果
	wxRsp, err := wechatPayer.MicroPay(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", *wxRsp)

	ok, err := wechatPayer.VerifySign(apiKey, "MD5", wxRsp)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("同步验签结果：", ok)
}

func WechatTradeQuery() {

	//初始化参数结构体
	body := make(zgo.BodyMap)
	//body.Set("transaction_id", "97HiM5j6kGmM2fk7fYMc8MgKhPnEQ5Rk")
	//body.Set("out_trade_no", "xfymd-1189114761577500672")
	//body.Set("transaction_id", "4200000403201910160804352254")
	body.Set("out_trade_no", "xfymd-1186894365230895104")
	body.Set("nonce_str", zgo.Utils.RandomString(32))
	//body.Set("sign_type", "MD5")

	//请求订单查询，成功后得到结果
	wxRsp, err := wechatPayer.OrderQuery(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("wxRsp：%+v\n", wxRsp)
}

func AliPayOrderPay() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("subject", "条码支付")
	body.Set("scene", "bar_code")
	body.Set("auth_code", "286248566432274952")
	body.Set("out_trade_no", "GZ201901301040361014")
	body.Set("total_amount", "0.01")
	body.Set("timeout_express", "2m")
	//条码支付
	aliRsp, err := aliPayer.OrderPay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
	alipayPublicKey := publicKey
	ok, err := aliPayer.VerifySign(alipayPublicKey, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		fmt.Println("err:::", err)
	}
	fmt.Println("同步返回验签：", ok)
}

func ZhimaCreditScoreGet() {

	//配置公共参数
	aliPayer.SetAuthToken("") //必须设置此参数
	//请求参数
	body := make(zgo.BodyMap)
	transaction_id := zgo.Utils.RandomString(48)
	body.Set("transaction_id", transaction_id)
	body.Set("product_code", "w1010100100000000001")

	//创建订单
	aliRsp, err := aliPayer.ZhimaCreditScoreGet(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", aliRsp)
}

func AliPayTradeQuery() {

	//配置公共参数
	//aliPayer.SetAppAuthToken("201908BB03f542de8ecc42b985900f5080407abc")
	//请求参数
	body := make(zgo.BodyMap)
	//body.Set("out_trade_no", "xfymd-1189090633902460928")
	//body.Set("out_trade_no", "jjr-1234567819-1")
	body.Set("trade_no", "2019110122001464701408403883")

	//查询订单
	aliRsp, err := aliPayer.OrderQuery(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", aliRsp.SignData)
}

func AliPayOpenAuthTokenApp() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("grant_type", "authorization_code")
	body.Set("code", "866185490c4e40efa9f71efea6766X02")
	//发起请求
	aliRsp, err := aliPayer.OpenAuthTokenApp(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func AliPayFundTransToaccountTransfer() {

	//请求参数
	body := make(zgo.BodyMap)
	out_biz_no := zgo.Utils.RandomString(32)
	body.Set("out_biz_no", out_biz_no)

	//body.Set("payee_type", "ALIPAY_LOGONID")
	//body.Set("payee_account", "qlegjm8279@sandbox.com")

	body.Set("payee_type", "ALIPAY_USERID")
	body.Set("payee_account", "2088702756752065")
	body.Set("amount", "0.10")
	//body.Set("payer_show_name", "发钱人名字")
	//body.Set("payee_real_name", "沙箱环境")
	body.Set("remark", "转账测试")
	//创建订单
	aliRsp, err := aliPayer.FundTransToaccountTransfer(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("aliRsp: %+v", aliRsp)
}

func AliPayFundTransOrderQuery() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("out_biz_no", "2WygRwTbSYIyg1uGo6OUSkVoNKzmOmDJ")
	//body.Set("order_id", "20191018110070001506980024341699")

	//创建订单
	aliRsp, err := aliPayer.FundTransOrderQuery(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("aliRsp: %+v", aliRsp)
}

func AliPayFundAccountQuery() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("alipay_user_id", "2088702756752065")
	//body.Set("account_product_code", "DING_ACCOUNT")

	//创建订单
	aliRsp, err := aliPayer.FundAccountQuery(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("aliRsp: %+v", aliRsp)
}

func AliPayTradePrecreate() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("subject", "预创建创建订单")
	body.Set("out_trade_no", "GZ201901301040355709")
	body.Set("total_amount", "1")
	//创建订单
	aliRsp, err := aliPayer.Order(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
	fmt.Println("aliRsp.QrCode:", aliRsp.AlipayTradePrecreateResponse.QrCode)
	fmt.Println("aliRsp.OutTradeNo:", aliRsp.AlipayTradePrecreateResponse.OutTradeNo)

	alipayPublicKey := publicKey
	ok, err := aliPayer.VerifySign(alipayPublicKey, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		fmt.Println("err:::", err)
	}
	fmt.Println("同步返回验签：", ok)
}

func AliPayTradeCreate() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("subject", "创建订单")
	body.Set("buyer_logon_id", "qlegjm8279@sandbox.com")
	body.Set("out_trade_no", "GZ2019013010403557082")
	body.Set("total_amount", "0.01")
	//创建订单
	aliRsp, err := aliPayer.OrderCreate(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
	fmt.Println("aliRsp.OrderNo:", aliRsp.AlipayTradeCreateResponse.TradeNo)
}

func AliPayTradePagePay() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("subject", "网站测试支付")
	body.Set("out_trade_no", "GZ201901301040355706100468")
	body.Set("total_amount", "88.88")
	body.Set("product_code", "FAST_INSTANT_TRADE_PAY")

	//电脑网站支付请求
	payUrl, err := aliPayer.OrderPagePay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("page payUrl:", payUrl)
}

func AliPayTradeWapPay() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("subject", "手机网站测试支付")
	body.Set("out_trade_no", "GZ2019013010403557034")
	body.Set("quit_url", "https://testpaybffp.example.com")
	body.Set("total_amount", "100.00")
	body.Set("product_code", "QUICK_WAP_WAY")

	//电脑网站支付请求
	payUrl, err := aliPayer.OrderWapPay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("wap payUrl:", payUrl)
}

func AliPayTradeAppPay() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("subject", "测试APP支付")
	body.Set("out_trade_no", "GZ201901301040355709")
	body.Set("total_amount", "1.00")
	//手机APP支付参数请求
	payParam, err := aliPayer.OrderAppPay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("app payurl:", payParam)
}

func AliPaySystemOauthToken() {

	//请求参数
	body := make(zgo.BodyMap)
	body.Set("grant_type", "authorization_code")
	body.Set("code", "19387fd27ed14e229f8cc3ad3174VB26")
	//手机APP支付参数请求
	payParam, err := aliPayer.SystemOauthToken(body)
	//{"access_token":"authusrB7f62d5653a6648c69ff5ebc7534c6D26","alipay_user_id":"20880004621430076585717262612626","expires_in":1296000,"re_expires_in":2592000,"refresh_token":"authusrB8046f6db234c4009a7b217f6c3a90B26","user_id":"2088421970930265"}

	//aliPayer.SetAuthToken(payParam.AlipaySystemOauthTokenResponse.AccessToken)
	//tradeRes, err := aliPayer.UserInfoShare()
	//{"code":"10000","msg":"Success","avatar":"https:\/\/tfs.alipayobjects.com\/images\/partner\/https:\/\/tfsimg.alipay.com\/images\/uemprod\/TB1VjWXbr8CDuNkUuZmxKOU0XXa.jpeg","city":"石家庄市","is_certified":"T","is_student_certified":"F","province":"河北省","user_id":"2088421970930265","user_status":"T","user_type":"1"}
	//fmt.Println(tradeRes,"-----")

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("app payurl:", payParam)
}
