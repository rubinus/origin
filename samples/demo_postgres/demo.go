package demo_postgres

import (
  "fmt"
  "github.com/gitcpu-io/zgo"
  "sync"
  "time"
)

/*
@Time : 2019-06-04 17:36
@Author : rubinus.chu
@File : demo
@project: origin
*/

type Account struct {
  //tableName struct{} `sql:"accounts"` // "_" means no name
  Bid     string  `json:"bid"`
  Aid     int32   `json:"aid"`
  Balance float64 `json:"balance"`
  Ctime   int64   `json:"ctime"`
  Utime   int64   `json:"utime"`
}

func InsertAccount(db *zgo.PostgresDB, name string) (*Account, error) {

  var err error

  account := &Account{
    Bid:     "b101",
    Aid:     1234,
    Balance: 0.30,
    Ctime:   zgo.Utils.GetTimestamp(13),
    Utime:   zgo.Utils.GetTimestamp(13),
  }
  err = db.Insert(account)
  if err != nil {
    return nil, err
  }
  return account, nil
}

func SelectAccount(db *zgo.PostgresDB, bid string, aid int32) (*Account, error) {
  var err error

  t1 := time.Now()

  var account []*Account
  err = db.Model(&account).Where("bid=? and aid=?", bid, aid).Select()

  //account := &Account{
  //	Bid: bid,
  //	Aid: aid,
  //}
  //err = db.Select(account)

  if err != nil {
    return nil, err
  }

  t2 := time.Now()
  fmt.Println(t2.Sub(t1))

  return account[0], nil
}

type Wxuser struct {
  //tableName  struct{} `sql:"wxusers"` // "_" means no name
  Id         int64  `json:"id"`
  UnionId    string `json:"union_id"`
  Nickname   string `json:"nickname"`
  HeadImgUri string `json:"head_img_uri"`
  CreateTime int64  `json:"create_time"`
  UpdateTime int64  `json:"update_time"`
}

type User struct {
  Id     int64    `json:"id"`
  Name   string   `json:"name"`
  Emails []string `json:"emails"`
}

type Share struct {
  //tableName    struct{} `sql:"shares"` // "_" means no name
  Id           int64  `json:"id"`
  Bid          string `json:"bid"`
  Aid          uint32 `json:"aid"`
  UnionId      string `json:"union_id"`
  ShareType    uint8  `json:"share_type"`
  PlatformType uint8  `json:"platform_type"`
  CreateTime   int64  `json:"create_time"`
  UpdateTime   int64  `json:"update_time"`
  WxUserId     int64
  WxUser       *Wxuser
}

func (u User) String() string {
  return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
  Id       int64
  Title    string
  AuthorId int64
  Author   *User
}

func (s Story) String() string {
  return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func CreateTable(db *zgo.PostgresDB) {
  err := createSchema(db)
  if err != nil {
    panic(err)
  }
}

func SelectUser1000(db *zgo.PostgresDB, id int64) {
  var err error

  t1 := time.Now()
  wg := sync.WaitGroup{}
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(i int) {
      user := &User{Id: id}
      err = db.Select(user)
      if err != nil {
        panic(err)
      }
      fmt.Println(i, user)
      wg.Done()
    }(i)

  }
  wg.Wait()
  t2 := time.Now()
  fmt.Println(t2.Sub(t1))
}

func SelectUser(db *zgo.PostgresDB, name string) ([]*User, error) {
  var err error

  t1 := time.Now()

  var users []*User

  err = db.Model(&users).Where("name=?", name).Select()
  if err != nil {
    return nil, err
  }

  t2 := time.Now()
  fmt.Println(t2.Sub(t1))

  return users, nil
}

func InsertUser(db *zgo.PostgresDB, name string) (*User, error) {

  var err error

  user1 := &User{
    Name:   name,
    Emails: []string{name + "@test.com", name + "@test.net"},
  }
  err = db.Insert(user1)
  if err != nil {
    return nil, err
  }
  return user1, nil
}

func InsertStory(db *zgo.PostgresDB, user1 *User) (*Story, error) {
  story1 := &Story{
    Title:    user1.Name + "的故事",
    AuthorId: user1.Id,
  }
  err := db.Insert(story1)
  if err != nil {
    return nil, err
  }
  return story1, nil
}

func SelectRelation(db *zgo.PostgresDB) (*Story, error) {
  story := new(Story)
  err := db.Model(story).
    Relation("Author").
    Where("story.id = ?", 1).
    Select()
  if err != nil {
    return nil, err
  }
  return story, nil
}

func SelectRelation2(db *zgo.PostgresDB) (*Share, error) {
  share := new(Share)
  err := db.Model(share).
    Relation("WxUser").
    Where("share.id = ?", 27).
    Select()
  if err != nil {
    return nil, err
  }
  return share, nil
}

func ExampleDB_Model() {

  //var err error

  //user1 := &User{
  //	Name:   "admin",
  //	Emails: []string{"admin1@admin", "admin2@admin"},
  //}
  //err = db.Insert(user1)
  //if err != nil {
  //	panic(err)
  //}

  //
  //err = db.Insert(&User{
  //	Name:   "root",
  //	Emails: []string{"root1@root", "root2@root"},
  //})
  //if err != nil {
  //	panic(err)
  //}
  //
  //story1 := &Story{
  //	Title:    "Cool story",
  //	AuthorId: user1.Id,
  //}
  //err = db.Insert(story1)
  //if err != nil {
  //	panic(err)
  //}
  //
  //// Select user by primary key.
  //user := &User{Id: user1.Id}

  //user := &User{Id: 5}
  //err = db.Select(user)
  //if err != nil {
  //	panic(err)
  //}
  //fmt.Println(user)

  //
  // Select all users.
  //var users []User
  //err = db.Model(&users).Where("name=?", "root").Select()
  //if err != nil {
  //	panic(err)
  //}

  // Select story and associated author in one query.
  //story := new(Story)
  //err = db.Model(story).
  //	Relation("Author").
  //	Where("story.id = ?", 1).
  //	Select()
  //if err != nil {
  //	panic(err)
  //}

  //fmt.Println(user)
  //fmt.Println(users)
  //fmt.Println(story)
  // Output: User<1 admin [admin1@admin admin2@admin]>
  // [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
  // Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}

func createShareAndWxuserSchema(db *zgo.PostgresDB) error {
  for _, model := range []interface{}{(*Wxuser)(nil), (*Share)(nil)} {
    fmt.Println(model)
    err := db.CreateTable(model, &zgo.PostgresCreateTableOptions{
      //Temp: true,
    })
    if err != nil {
      return err
    }
  }
  return nil
}

func createSchema(db *zgo.PostgresDB) error {
  for _, model := range []interface{}{(*User)(nil), (*Story)(nil)} {
    fmt.Println(model)
    err := db.CreateTable(model, &zgo.PostgresCreateTableOptions{
      //Temp: true,
    })
    if err != nil {
      return err
    }
  }
  return nil
}

func createAccountSchema(db *zgo.PostgresDB) error {
  for _, model := range []interface{}{(*Account)(nil)} {
    err := db.CreateTable(model, &zgo.PostgresCreateTableOptions{
      //Temp: true,
    })
    if err != nil {
      return err
    }
  }
  return nil
}

func ExampleDB_Insert_dynamicTableName(db *zgo.PostgresDB) {
  type NamelessModel struct {
    tableName struct{} `sql:"_"` // "_" means no name
    Id        int
  }

  err := db.Model((*NamelessModel)(nil)).Table("dynamic_name").CreateTable(nil)
  panicIf(err)

  row123 := &NamelessModel{
    Id: 123,
  }
  _, err = db.Model(row123).Table("dynamic_name").Insert()
  panicIf(err)

  row := new(NamelessModel)
  err = db.Model(row).Table("dynamic_name").First()
  panicIf(err)
  fmt.Println("id is", row.Id)

  err = db.Model((*NamelessModel)(nil)).Table("dynamic_name").DropTable(nil)

  panicIf(err)

  // Output: id is 123
}

func panicIf(err error) {
  if err != nil {
    panic(err)
  }
}
