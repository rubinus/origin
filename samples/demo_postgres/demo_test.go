package demo_postgres

import (
  "fmt"
  "github.com/gitcpu-io/zgo"
  "strconv"
  "testing"
  "time"
)

/*
@Time : 2019-06-04 17:36
@Author : rubinus.chu
@File : demo_test
@project: origin
*/

const (
  label_bj = "postgres_label_bj"
  label_sh = "postgres_label_sh"
)

var project = "1553240759"

func TestSelectAccount(t *testing.T) {
  err := zgo.Engine(&zgo.Options{
    Env:     "local",
    Project: project,

    Postgres: []string{
      label_sh,
    },
  }) //测试时表示使用postgres，在origin中使用一次

  if err != nil {
    panic(err)
  }

  time.Sleep(2 * time.Second)

  dbch, err := zgo.Postgres.GetConnChan() //有多个时，需要指定label
  if err != nil {
    panic(err)
    return
  }
  db := <-dbch

  //createAccountSchema(db)

  account, err := SelectAccount(db, "b101", 1234)
  //account, err := InsertAccount(db,"")
  if err != nil {
    fmt.Println(err)
    return
  }

  dst := strconv.FormatFloat(account.Balance, 'f', 2, 64)
  fmt.Printf("账号%+v, 余额：%v, 转换：%s：", account, account.Balance, dst)
}

//创建表
func TestCreateTable(t *testing.T) {
  err := zgo.Engine(&zgo.Options{
    Env:     "dev",
    Project: "1560588569",

    Postgres: []string{
      label_bj,
    },
  }) //测试时表示使用postgres，在origin中使用一次

  if err != nil {
    panic(err)
  }

  time.Sleep(2 * time.Second)

  //db,_ := zgo.Postgres.New(label_sh)

  dbch, err := zgo.Postgres.GetConnChan() //有多个时，需要指定label
  if err != nil {
    panic(err)
  }
  db := <-dbch

  //创建表
  //CreateTable(db)

  //createShareAndWxuserSchema(db)

  //SelectUser1000(db, 11)
  //for {
  //	select {
  //	case <-time.Tick(1 * time.Minute):
  //		fmt.Println("begin use pool 1 分钟 。。。")
  //		SelectUser1000(db, 11)
  //
  //	}
  //
  //}

  //***************************************插入一条
  r, err := InsertUser(db, "朱大仙儿")
  if err != nil {
    fmt.Println("---error-InsertUser--", err)
    return
  }
  //
  //s, err := InsertStory(db, r)
  //if err != nil {
  //	fmt.Println("---error InsertStory---", err)
  //	return
  //}
  //
  fmt.Println(r, "---插入一条-用户--")
  //fmt.Println(s, "---插入一条-故事--")

  //***************************************通过name查询，返回多条
  //u, err := SelectUser(db, "朱大仙儿")
  //if err != nil {
  //	fmt.Println("---error---", err)
  //	return
  //}
  //fmt.Println(u, "----查询----")

  //***************************************查询关连表
  //sr, err := SelectRelation(db)
  //sr, err := SelectRelation2(db)
  //if err != nil {
  //	fmt.Println("---error-SelectRelation--", err)
  //	return
  //}
  //fmt.Printf("%+v%s", sr, "----查询--SelectRelation--")

}
