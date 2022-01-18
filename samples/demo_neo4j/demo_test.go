package demo_neo4j

import (
  "fmt"
  "github.com/gitcpu-io/zgo"
  "testing"
  "time"
)

/*
@Time : 2019-06-04 20:17
@Author : rubinus.chu
@File : demo_test
@project: origin
*/

const (
  label_bj = "neo4j_label"
)

var project = "1553240759"

func TestGetPeople(t *testing.T) {
  err := zgo.Engine(&zgo.Options{
    Env:     "dev",
    Project: project,

    Neo4j: []string{
      label_bj,
    },
  }) //测试时表示使用neo4j，在origin中使用一次

  if err != nil {
    panic(err)
  }

  time.Sleep(2 * time.Second)

  dbch, err := zgo.Neo4j.GetConnChan() //有多个时，需要指定label
  if err != nil {
    panic(err)
  }

  db := <-dbch

  p, err := GetPeople(db)
  if err != nil {
    panic(err)
  }
  fmt.Println(p)
}
