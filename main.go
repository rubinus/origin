package main

import (
	"flag"
	"fmt"
	"git.zhugefang.com/gobase/origin/config"
	"git.zhugefang.com/gobase/origin/engine"
	"git.zhugefang.com/gobase/origin/routes"
	"git.zhugefang.com/gobase/origin/server"
	"github.com/kataras/iris"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
)

var wg = new(sync.WaitGroup)

func init() {
	var (
		env       string
		project   string
		etcdHosts string
		port      string
		rpcPort   string
	)
	//默认读取config/local.json
	flag.StringVar(&env, "env", "dev", "start local/dev/qa/pro env config")
	flag.StringVar(&project, "project", "", "create project id by zgo engine admin")
	flag.StringVar(&etcdHosts, "etcdHosts", "", "etcd hosts host:port,host:port")
	flag.StringVar(&port, "port", "", "port")
	flag.StringVar(&rpcPort, "rpcPort", "", "rpcPort")

	flag.Parse()
	if os.Getenv("ENV") != "" { //从os的env取得ENV，用来在yaml文件中的配置
		env = os.Getenv("ENV")
	}
	//load config from dev/qa/pro
	config.InitConfig(env, project, etcdHosts, port, rpcPort)

	if os.Getenv("PROJECT") != "" {
		config.Conf.Project = os.Getenv("PROJECT") //从os的env取得PROJECT，用来在yaml文件中的配置
	}
	if os.Getenv("ETCDHOSTS") != "" {
		config.Conf.EtcdHosts = os.Getenv("ETCDHOSTS") //从os的env取得ETCDHOSTS，用来在yaml文件中的配置
	}
	if os.Getenv("PORT") != "" {
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		config.Conf.ServerPort = port //从os的env取得PORT，用来在yaml文件中的配置
	}
	if os.Getenv("RPCPORT") != "" {
		config.Conf.RpcPort = os.Getenv("RPCPORT") //从os的env取得RPCPORT，用来在yaml文件中的配置
	}
}

func main() {
	//优雅的退出
	WatchSignal()

	err := engine.Run() //start zgo engine
	if err != nil {
		panic(err)
	}

	app := iris.New() //start web

	var pre string
	if config.Conf.UsePreAbsPath == 1 {
		prefix, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "/")
		pre = prefix + "/views"
	} else {
		pre = "./views"
	}

	app.StaticWeb("/", "./public") //static

	app.RegisterView(iris.HTML(pre, ".html")) // select the html engine to serve templates

	//集中调用路由
	routes.Index(app)

	//测试消费nsq
	//demo_nsq.Consumer()
	//测试消费kafka
	//demo_kafka.Consumer()

	go func() { //start grpc server on the default port 50051
		server.Start()
	}()

	app.Run(iris.Addr(":" + strconv.Itoa(config.Conf.ServerPort)))

	wg.Wait()
}

func WatchSignal() {
	//创建监听退出chan
	signalChan := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range signalChan {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				go func() {
					wg.Add(1)
					defer wg.Done()

					fmt.Println("---i am out over from killed cmd, but not kill -9, you can do something ...---")

				}()
			case syscall.SIGUSR1:
				//todo something
				fmt.Println("usr1", s)
			case syscall.SIGUSR2:
				//todo something

				fmt.Println("usr2", s)
			default:
				//todo something

				fmt.Println("other", s)
			}
		}
	}()
}
