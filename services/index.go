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
	return &svc{
		repo: models.NewTradeRepo(),
	}
}

type svc struct {
	repo models.Trader
}

func (svc *svc) Search(pay *models.PayRequest) (*models.Trade, error) {
	trade := &models.Trade{}
	//todo 构造Trade

	err := svc.repo.Insert()
	if err != nil {
		return nil, err
	}
	return trade, nil
}
