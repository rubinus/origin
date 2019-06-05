package handlers

import (
	"context"
	//"fmt"
	"time"

	"git.zhugefang.com/gocore/zgo"
	"github.com/kataras/iris"
)

type AdminCallEnd struct {
	//Id            int       `json:"id"`
	CallId        string    `gorm:"call_id"`
	TaskId        string    `gorm:"task_id"`
	UserId        string    `gorm:"user_id"`
	CompanyId     string    `gorm:"company_id"`
	Phone         string    `gorm:"phone"`
	StatusCode    string    `gorm:"status_code"`
	Duration      int       `gorm:"duration"`
	Caller        string    `gorm:"caller"`
	StatusMsg     string    `gorm:"status_msg"`
	StartTime     string    `gorm:"start_time"`
	EndTime       string    `gorm:"end_time"`
	CreateAt      time.Time `gorm:"create_at"`
	Userlevel     string    `gorm:"column:userlevel"`
	Recordingurl  string    `gorm:"column:recordingurl"`
	QiNiuUrl      string    `gorm:"column:qiniuurl"`
	NewUserLevel  string    `gorm:"column:newuserlevel"`
	MatchTags     string    `gorm:"column:match_tags"`
	Read          int       `gorm:"read"`
	CDate         int       `gorm:"c_date"`
	IsLongTrip    int       `gorm:"column:is_longtrip"`
	PhoneCity     string    `gorm:"phone_city"`
	PerimeterName string    `gorm:"perimeter_name"`
	TollType      string    `gorm:"toll_type"`
}

func MysqlGetV1(ctx iris.Context) {
	//fmt.Println("mysql get start")
	city := ctx.Params().GetStringDefault("city", "")
	if city == "" {
		zgo.Http.JsonpErr(ctx, "city is Null")
		return
	}

	var dbname = city
	var result = &AdminCallEnd{}
	var args = make(map[string]interface{}, 0)
	var err error

	args["table"] = dbname + ".admin_call_end"
	args["query"] = " id = ?"
	args["args"] = []interface{}{int(1)}
	args["obj"] = result

	err = zgo.Mysql.Get(context.TODO(), args)

	if err != nil {
		zgo.Http.JsonpErr(ctx, "execute sql error: "+err.Error())
		return
	}

	zgo.Http.JsonpOK(ctx, result)

	//fmt.Println("mysql get end")
	return
}

func MysqlAddV1(ctx iris.Context) {
	//fmt.Println("mysql add start")
	city := ctx.Params().GetStringDefault("city", "")
	if city == "" {
		zgo.Http.JsonpErr(ctx, "city is empty")
		return
	}
	var err error

	govhs := &AdminCallEnd{
		CallId:     "114850357335^101657167335",
		TaskId:     "11",
		UserId:     "94",
		CompanyId:  "1",
		Phone:      "15313500146",
		StatusCode: "200005",
		Duration:   10,
		StatusMsg:  "用户拒绝",
		Caller:     "01086488021",
		StartTime:  "",
		EndTime:    "2018-10-19 17:21:34",
		CreateAt:   time.Now(),
		//UserLevel:  "F",
		Read:       1,
		CDate:      20190325,
		IsLongTrip: 1,
	}

	var dbname = city
	var args = make(map[string]interface{}, 0)
	args["table"] = dbname + ".admin_call_end"
	args["obj"] = govhs

	//err = zgo.Mysql.Create(context.TODO(), args)

	if err != nil {
		zgo.Http.JsonpErr(ctx, "insert data error:"+err.Error())
		return
	}

	zgo.Http.JsonOK(ctx, "ok")
	//fmt.Println("mysql add end")
	return
}
