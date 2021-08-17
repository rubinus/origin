package demo_mysql

import (
	"context"
	"fmt"
	"github.com/rubinus/origin/config"
	"github.com/rubinus/zgo"
	"github.com/rubinus/zgo/zgomysql"
	"time"
)

type MysqlDemo struct {
}

func init() {

	config.InitConfig("local", "", "", "", "")

}

// 基类
type BaseModel struct {
	zgomysql.BaseModel // 继承zgo基类
	Ctime              int64
	Utime              int64
}

func (self *BaseModel) BeforeUpdate() error {
	self.Utime = time.Now().Unix()
	return nil
}

func (self *BaseModel) BeforeCreate() error {
	self.Ctime = time.Now().Unix()
	self.Utime = time.Now().Unix()
	return nil
}

func (self *BaseModel) GetId() (uint32, error) {
	return self.Id, nil
}

func (self *BaseModel) SetId(id uint32) error {
	self.Id = id
	return nil
}

var label = "mysql_sell_2"

// 实体类
type House struct {
	BaseModel
	Name string `json:"name"`
}

func (h *House) TableName() string {
	return "house"
}

func (h *House) DbName() string {
	return ""
}

//QueryMysql 测试读取Mysqldb数据，wait for sdk init connection
func (m MysqlDemo) Get(id uint32) (*House, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//查询参数
	zgo.Engine(&zgo.Options{
		Env:     "dev",
		Project: config.Conf.Project,
		Mysql: []string{
			"1730737039440", // 默认为 第一个label
			//"mysql_sell_2",
		},
	})

	args := make(map[string]interface{})
	obj := &House{}
	// 拼接dbname   如果不随城市变化，固定城市只用写表名即可
	args["table"] = obj.TableName()

	args["query"] = " id = ? "       //  用问号作为占位符
	args["args"] = []interface{}{id} // 参数 放入interface的slice里面
	args["obj"] = obj                // 输出对象
	// 通过城市和业务获取对应label
	c, err := zgo.Mysql.New(label)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 输入参数：上下文ctx，args具体的查询操作参数
	err1 := c.Get(ctx, args)

	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, err1
	}
	return obj, nil
}

func (m MysqlDemo) Create(name string) (*House, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//查询参数
	zgo.Engine(&zgo.Options{
		Env:     config.Conf.Env,
		Project: config.Conf.Project,
		Mysql: []string{
			//"mysql_sell_1", // 默认为 第一个label
			"mysql_sell_2",
		},
	})
	for {
		if zgo.Mysql != nil {
			fmt.Println(zgo.Mysql)
			break
		} else {
			fmt.Println(zgo.Cache)
		}
	}

	// id为自增id，禁止赋值
	obj := &House{Name: name}
	// 也可以直接通过label获取
	c, err := zgo.Mysql.New(label)
	if err != nil {
		// 获取失败
		fmt.Println(err.Error())
		return nil, err
	}
	// 输入参数：上下文ctx，args具体的查询操作参数
	err1 := c.Create(ctx, obj)

	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, err1
	}
	// 返回保存的对象
	return obj, nil
}

func (m MysqlDemo) List() (*[]House, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//查询参数
	zgo.Engine(&zgo.Options{
		Env:     config.Conf.Env,
		Project: config.Conf.Project,
		Mysql: []string{
			//"mysql_sell_1", // 默认为 第一个label
			"mysql_sell_2",
		},
	})

	args := make(map[string]interface{})
	obj := &[]House{}
	// 拼接dbname   如果不随城市变化，固定城市只用写表名即可。
	a := &House{}
	args["table"] = a.TableName()

	args["query"] = " id in (?) or id > ?" //  用问号作为占位符
	ids := []int{1, 2, 3}
	args["args"] = []interface{}{ids, 5} // 参数 放入interface的slice里面
	args["obj"] = obj                    // 查询结果输出对象
	args["offset"] = 10                  // 从50开始
	args["limit"] = 30                   // 查询30条
	args["order"] = " id desc "          // 排序
	// 通过城市和业务获取对应label
	// 也可以直接通过label获取
	c, err := zgo.Mysql.New(label)
	if err != nil {
		// 获取失败
		fmt.Println(err.Error())
		return nil, err
	}
	// 输入参数：上下文ctx，args具体的查询操作参数
	err1 := c.List(ctx, args)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, err1
	}
	return obj, nil
}

func (m MysqlDemo) Count() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//查询参数
	zgo.Engine(&zgo.Options{
		Env:     config.Conf.Env,
		Project: config.Conf.Project,
		Mysql: []string{
			//"mysql_sell_1", // 默认为 第一个label
			"mysql_sell_2",
		},
	})

	args := make(map[string]interface{})
	a := &House{}
	args["table"] = a.TableName()
	args["query"] = " id in (?) or id > ?" //  用问号作为占位符
	ids := []int{1, 2, 3}
	status := 1
	args["args"] = []interface{}{ids, status, 5} // 参数 放入interface的slice里面
	count := 0
	args["obj"] = &count // 输出对象 需要传入指针
	// 通过城市和业务获取对应label
	// 也可以直接通过label获取
	c, err := zgo.Mysql.New(label)
	if err != nil {
		// 获取失败
		fmt.Println(err.Error())
		return count, err
	}
	// 输入参数：上下文ctx，args具体的查询操作参数
	err1 := c.Count(ctx, args)

	if err1 != nil {
		fmt.Println(err1.Error())
		return count, err1
	}
	return count, nil
}

// 只更新非空值
func (m MysqlDemo) UpdateOne(id uint32, name string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//查询参数
	zgo.Engine(&zgo.Options{
		Env:     config.Conf.Env,
		Project: config.Conf.Project,
		Mysql: []string{
			//"mysql_sell_1", // 默认为 第一个label
			"mysql_sell_2",
		},
	})

	obj := House{}
	obj.Id = id
	obj.Name = name
	//data := map[string]interface{}{"name": "aaa"} // name 改成 aaa
	c, err := zgo.Mysql.New(label)
	if err != nil {
		// 获取失败
		fmt.Println(err.Error())
		return 0, err
	}
	// 输入参数：上下文ctx，args具体的查询操作参数
	cn, err1 := c.UpdateNotEmptyByObj(ctx, &obj)

	if err1 != nil {
		fmt.Println(err1.Error())
		return 0, err1
	}
	return cn, nil
}

// 更新Data中的字段全部数据(包括空值)
func (m MysqlDemo) UpdateByData(id uint32, name string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//查询参数
	zgo.Engine(&zgo.Options{
		Env:     config.Conf.Env,
		Project: config.Conf.Project,
		Mysql: []string{
			//"mysql_sell_1", // 默认为 第一个label
			"mysql_sell_2",
		},
	})

	obj := House{}
	obj.Id = id
	obj.Name = name
	data := map[string]interface{}{"name": "aaa"} // name 改成 aaa
	c, err := zgo.Mysql.New(label)
	if err != nil {
		// 获取失败
		fmt.Println(err.Error())
		return 0, err
	}
	// 输入参数：上下文ctx，args具体的查询操作参数
	cn, err1 := c.UpdateByData(ctx, &obj, data)

	if err1 != nil {
		fmt.Println(err1.Error())
		return 0, err1
	}
	return cn, nil
}

// 更新全部字段(包括空值)
func (m MysqlDemo) UpdateByObj(id uint32, name string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//查询参数
	zgo.Engine(&zgo.Options{
		Env:     config.Conf.Env,
		Project: config.Conf.Project,
		Mysql: []string{
			//"mysql_sell_1", // 默认为 第一个label
			"mysql_sell_2",
		},
	})

	obj := House{}
	obj.Id = id
	obj.Name = name
	c, err := zgo.Mysql.New(label)
	if err != nil {
		// 获取失败
		fmt.Println(err.Error())
		return 0, err
	}
	// 输入参数：上下文ctx，args具体的查询操作参数
	cn, err1 := c.UpdateByObj(ctx, &obj)

	if err1 != nil {
		fmt.Println(err1.Error())
		return 0, err1
	}
	return cn, nil
}

func (m MysqlDemo) DeleteOne(id uint32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//查询参数
	zgo.Engine(&zgo.Options{
		Env:     config.Conf.Env,
		Project: config.Conf.Project,
		Mysql: []string{
			//"mysql_sell_1", // 默认为 第一个label
			"mysql_sell_2",
		},
	})

	// 也可以直接通过label获取
	c, err := zgo.Mysql.New(label)
	if err != nil {
		// 获取失败
		fmt.Println(err.Error())
		return 0, err
	}
	// 输入参数：上下文ctx，args具体的查询操作参数
	a := &House{}
	cn, err1 := c.DeleteById(ctx, a.TableName(), id)

	if err1 != nil {
		fmt.Println(err1.Error())
		return 0, err1
	}
	return int64(cn), nil
}
