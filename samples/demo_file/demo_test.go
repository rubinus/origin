package demo_file

import (
	"bufio"
	"fmt"
	"github.com/rubinus/origin/config"
	"github.com/rubinus/zgo"
	"os"
	"strconv"
	"testing"
	"time"
)

/*
@Time : 2019-02-28 20:38
@Author : rubinus.chu
@File : demo_test
@project: origin.git
*/

func TestF(t *testing.T) {
	config.InitConfig("local", "", "", "", "")

	err := zgo.Engine(&zgo.Options{
		Env:      config.Conf.Env,
		Project:  config.Conf.Project,
		Loglevel: config.Conf.Loglevel,
	})

	if err != nil {
		panic(err)
	}

	ch := make(chan string)

	go func() {
		c := 0
		for v := range ch {
			c++
			fmt.Println(c, v)
			//zgo.Mysql.Create(context.TODO())
		}
	}()

	//F()
	fd, err := os.OpenFile("shoujihao.txt", os.O_CREATE|os.O_WRONLY, 0600)

	output := bufio.NewWriter(fd)

	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		s := strconv.Itoa(i)
		if len(s) <= 6 {
			pre := ""
			for ii := 0; ii < 6-len(s); ii++ {
				pre = "0" + pre
			}
			s = pre + s
		}
		o := "10100" + s
		ch <- o
		output.WriteString(o)
		output.WriteString("\n")
	}
	t2 := time.Now()
	output.Flush()
	fd.Close()
	t3 := time.Now()

	fmt.Println("读到内存：", t2.Sub(t1))
	fmt.Println("由内存刷到硬盘：", t3.Sub(t2))
	fmt.Println("全部耗时：", t3.Sub(t1))

	time.Sleep(10 * time.Second)
}
