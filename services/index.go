package services

import "git.zhugefang.com/gobase/origin/models"

/*
@Time : 2019-03-06 19:02
@Author : rubinus.chu
@File : index
@project: origin
*/

type Payer interface {
	Insert(payReq *models.PayRequest) (*models.Trade, error)
	//请在此处添加其它方法
}

func NewPay() Payer {
	return &svc{
		repo: models.NewTradeRepo(),
	}
}

type svc struct {
	repo models.Trader
}

// Insert保存方法
func (svc *svc) Insert(payReq *models.PayRequest) (*models.Trade, error) {
	trade := &models.Trade{}
	//todo 通过传入的参数payReq 来构造Trade

	err := svc.repo.Insert(trade)
	if err != nil {
		return nil, err
	}
	return trade, nil
}

//todo add other func
