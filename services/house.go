/*
@Time : 2019-03-08 17:40
@Author : lucas
@File : House
@project: visource
*/
package services

import (
	"context"
	"errors"
	"git.zhugefang.com/goymd/visource/models/mysql_models"
)

type HouseService struct {
}

func (hs *HouseService) AddHouse(ctx context.Context, h *mysql_models.House, city string) error {
	//mysqlService, err := zgo.Mysql.MysqlServiceByCityBiz(city, "sell")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return err
	//}
	//dbname, err1 := mysqlService.GetDbByCityBiz(city, "sell")
	//if err1 != nil {
	//	fmt.Println(err.Error())
	//	return err1
	//}
	//
	//args := make(map[string]interface{})
	//args["obj"] = h
	//args["table"] = dbname + "." + h.TableName()
	//err2 := mysqlService.Create(ctx, args)
	//if err2 != nil {
	//	fmt.Println(err2.Error())
	//	return err2
	//}
	//if h.Id > 0 {
	//	return nil
	//}
	return errors.New("创建失败")
}
