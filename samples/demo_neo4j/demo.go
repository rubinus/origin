package demo_neo4j

import (
  "fmt"
  "github.com/gitcpu-io/zgo"
)

/*
@Time : 2019-06-04 20:16
@Author : rubinus.chu
@File : demo
@project: origin
*/

func Trans(tx zgo.Neo4jTransaction) (interface{}, error) {
  var list []string
  var result zgo.Neo4jResult
  var err error

  if result, err = tx.Run("MATCH (a:Person) RETURN a.name ORDER BY a.name", nil); err != nil {
    return nil, err
  }

  if err := result.Err(); err != nil {
    return nil, err
  }

  for result.Next() {
    // 输出结果集中的记录
    fmt.Println(result.Record())
    list = append(list, result.Record().GetByIndex(0).(string))
  }

  return list, nil
}

func GetPeople(session zgo.Neo4jSession) ([]string, error) {
  var people interface{}
  var err error

  people, err = session.ReadTransaction(Trans)
  if err != nil {
    return nil, err
  }

  return people.([]string), nil
}
