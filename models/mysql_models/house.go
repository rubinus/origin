/*
@Time : 2019-03-08 16:37
@Author : lucas
@File : house
@project: base-to-base-wait-copy
*/
package mysql_models

type House struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (h *House) TableName() string {
	return "house"
}
