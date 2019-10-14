package services

import "git.zhugefang.com/gobase/origin/models"

/*
@Time : 2019-03-06 19:02
@Author : rubinus.chu
@File : index
@project: origin
*/

type Payer interface {
	Search(pay *models.PayRequest) (*models.Trade, error)
}

func NewPay() Payer {
	return &A{}
}

type A struct {
}

//Search 这个func 由http与grpc共用
func (a *A) Search(pay *models.PayRequest) (*models.Trade, error) {
	trade := &models.Trade{}

	//todo 构造Trade
	//trade.Channel = 1

	err := trade.Insert()
	if err != nil {
		return nil, err
	}
	return trade, nil
}
