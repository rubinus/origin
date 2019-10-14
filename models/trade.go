package models

import (
	"git.zhugefang.com/gocore/zgo"
)

/*
@Time : 2019-10-14 14:18
@Author : rubinus.chu
@File : trade
@project: origin
*/

type Trader interface {
	Insert(trade *Trade) error
}

func NewTradeRepo() Trader {
	return &Trade{}
}

// 映射pg表结构
type Trade struct {
	//tableName  struct{}               `sql:"trades"` // "_" means no name
	Id  int64  `json:"id" sql:"id,pk"`
	Bid string `json:"bid"`
	Aid uint32 `json:"aid"`

	PayType uint8  `json:"pay_type"`
	Appid   string `json:"appid"`
	BuyerId string `json:"buyer_id"`

	OrderNo string `json:"order_no"`
	TradeNO string `json:"trade_no"`

	TradeType uint8 `json:"trade_type"`
	Channel   uint8 `json:"channel"`

	Amount  int         `json:"amount"`
	FeeType string      `json:"fee_type"`
	Body    string      `json:"body"`
	Detail  interface{} `json:"detail"`

	Status     uint8  `json:"status"`
	StatusDesc string `json:"status_desc"`
	Attach     string `json:"attach"`

	TradeTime  int64 `json:"trade_time"`
	ExpireTime int64 `json:"expire_time"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

// Insert保存方法
func (repo *Trade) Insert(trade *Trade) error {

	if trade.CreateTime == 0 {
		trade.CreateTime = zgo.Utils.GetTimestamp(10)
	}
	if trade.UpdateTime == 0 {
		trade.UpdateTime = trade.CreateTime
	}

	//取db连接
	dbCh, err := zgo.Postgres.GetConnChan()
	if err != nil {
		zgo.Log.Error("db get Conn Error:" + err.Error())
		return err
	}

	if db, ok := <-dbCh; !ok {
		zgo.Log.Error("db get ConnChan Error:" + err.Error())
		return err
	} else {
		err := db.Insert(trade)
		if err != nil {
			zgo.Log.Error(err)
		}
		return nil
	}
}
