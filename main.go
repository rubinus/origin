package main

import (
	"context"
	"flag"
	"fmt"
	"git.zhugefang.com/gobase/origin/backend"
	"git.zhugefang.com/gobase/origin/config"
	"git.zhugefang.com/gobase/origin/engine"
	"git.zhugefang.com/gobase/origin/routes"
	"git.zhugefang.com/gobase/origin/server"
	"git.zhugefang.com/gocore/zgo"
	"github.com/kataras/iris"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//var wg = new(sync.WaitGroup)

func init() {
	var (
		env          string
		project      string
		etcdHosts    string
		port         string
		rpcPort      string
		svcName      string
		svcHost      string
		svcHttpPort  string
		svcGrpcPort  string
		svcEtcdHosts string
	)
	//默认读取config/local.json
	flag.StringVar(&env, "env", "dev", "start local/dev/qa/pro env config")

	flag.StringVar(&project, "project", "", "create project id by zgo engine admin")

	flag.StringVar(&etcdHosts, "etcdHosts", "", "etcd hosts host:port,host:port")

	flag.StringVar(&port, "port", "", "http port")

	flag.StringVar(&rpcPort, "rpcPort", "", "grpc port")

	//暴露服务信息
	flag.StringVar(&svcName, "svc_name", "", "让服务对外可访问的名称")

	flag.StringVar(&svcHost, "svc_host", "", "让服务对外可访问的主机地址，默认是宿主机的内网IP")

	flag.StringVar(&svcHttpPort, "svc_http_port", "", "让服务http对外可访问的端口号，默认是80")

	flag.StringVar(&svcGrpcPort, "svc_grpc_port", "", "让服务grpc对外可访问的端口号，默认是50051")

	flag.StringVar(&svcEtcdHosts, "svc_etcd_hosts", "", "让服务可以用外部的注册中心地址，默认与zgo engine相同")

	flag.Parse()
	if os.Getenv("ENV") != "" { //从os的env取得ENV，用来在yaml文件中的配置
		env = os.Getenv("ENV")
	}

	//load config from dev/qa/pro
	config.InitConfig(env, project, etcdHosts, port, rpcPort)

	//输入覆盖配置中的.json中的
	if svcName != "" {
		config.Conf.ServiceInfo.SvcName = svcName
	}
	if svcHost != "" {
		config.Conf.ServiceInfo.SvcHost = svcHost
	}
	if svcHttpPort != "" {
		config.Conf.ServiceInfo.SvcHttpPort = svcHttpPort
	}
	if svcGrpcPort != "" {
		config.Conf.ServiceInfo.SvcGrpcPort = svcGrpcPort
	}
	if svcEtcdHosts != "" {
		config.Conf.ServiceInfo.SvcEtcdHosts = svcEtcdHosts
	}

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

	//用docker配置覆盖服务注册与发现的配置
	if os.Getenv("SVC_NAME") != "" {
		config.Conf.ServiceInfo.SvcName = os.Getenv("SVC_NAME") //来os的env，用来在yaml文件中的配置
	}
	if os.Getenv("SVC_HOST") != "" {
		config.Conf.ServiceInfo.SvcHost = os.Getenv("SVC_HOST") //来os的env，用来在yaml文件中的配置
	}
	if os.Getenv("SVC_HTTP_PORT") != "" {
		config.Conf.ServiceInfo.SvcHttpPort = os.Getenv("SVC_HTTP_PORT") //来os的env，用来在yaml文件中的配置
	}
	if os.Getenv("SVC_GRPC_PORT") != "" {
		config.Conf.ServiceInfo.SvcGrpcPort = os.Getenv("SVC_GRPC_PORT") //来os的env，用来在yaml文件中的配置
	}
	if os.Getenv("SVC_ETCD_HOSTS") != "" {
		config.Conf.ServiceInfo.SvcEtcdHosts = os.Getenv("SVC_ETCD_HOSTS") //来os的env，用来在yaml文件中的配置
	}

	if config.Conf.ServiceInfo.SvcHttpPort == "" {
		config.Conf.ServiceInfo.SvcHttpPort = port
	}
	if config.Conf.ServiceInfo.SvcGrpcPort == "" {
		config.Conf.ServiceInfo.SvcGrpcPort = rpcPort
	}

}

func main() {
	//优雅的退出 -- 使用iris框架中的退出
	//WatchSignal()

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

	//消费nsq 需要先配置上nsq
	//queue.NsqConsumer()
	//消费kafka 需要先配置上kafka
	//queue.KafkaConsumer()
	//消费Rabbitmq
	//queue.RabbitmqConsumer() //需要先配置上rabbitmq

	go func() { //start grpc server on the default port 50051 如果作为rpc服务端，让其它client连接进来
		server.Start()
	}()

	if config.Conf.StartService == false { //不使用服务发现，原来标准模式

		//###################### ################
		//todo 启动GRPC 客户端 如果作为客户端要连其它rpc server需要开启下面，否则可注释掉
		//###################### ################
		backend.RPCClientsRun(nil) //start grpc client

		//run起自己
		//****change four*****
		iris.RegisterOnInterrupt(func() {
			timeout := 5 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			fmt.Println("######origin, this server shutdown by Iris, you can do something from here ...######")
			// 关闭所有主机
			_ = app.Shutdown(ctx)
		})
		_ = app.Run(iris.Addr(":"+strconv.Itoa(config.Conf.ServerPort), func(h *iris.Supervisor) {
			h.RegisterOnShutdown(func() {

			})

		}), iris.WithoutInterruptHandler)

	} else {
		//使用服务注册与服务发现模式
		useServiceRegistryDiscover(app)
	}

	//wg.Wait()
}

func useServiceRegistryDiscover(app *iris.Application) {
	//注册服务到 注册中心etcd中，然后监听当前服务使用的其它服务的名字

	//***********************************************************
	//第一步 必须
	//***********************************************************
	registryAndDiscover, err := zgo.Service.New(5,
		config.Conf.ServiceInfo.SvcEtcdHosts)
	//与注册中心保持心跳间隔5秒，默认与zgo engine的etcd相同
	if err != nil {
		zgo.Log.Errorf("创建微服务实例化失败 %v", err)
		return
	}

	//***********************************************************
	//第二步 配置文件决定是否开启使用，这一步很重要，一定要注册外部可以访问当前服务的，尤其用docker时要注意
	//***********************************************************
	//todo 请确认下面三项 host httpport grpcport 使其它服务可访问到
	if config.Conf.StartServiceRegistry == true {
		var host string
		if config.Conf.ServiceInfo.SvcHost == "" { //默认为空使用宿主机内部IP
			host = zgo.Utils.GetIntranetIP()
		} else {
			host = config.Conf.ServiceInfo.SvcHost
		}
		err = registryAndDiscover.Registry(config.Conf.ServiceInfo.SvcName, host,
			config.Conf.ServiceInfo.SvcHttpPort, config.Conf.ServiceInfo.SvcGrpcPort)
		//注册当前服务(自己)到注册中心
		if err != nil {
			zgo.Log.Errorf("%s 注册微服务失败 %v", "test", err)
			return
		}
	}

	//***********************************************************
	//第三步 配置文件决定是否开启使用，这是服务发现的监听，必须与第四步同时使用
	//***********************************************************
	if config.Conf.StartServiceDiscover == true {
		watch := zgo.Service.Watch()
		httpChan := make(chan string, 1000)
		grpcChan := make(chan string, 1000)
		go func() {
			for value := range watch {
				go func(value string) {
					config.WatchHttpConfigByService(httpChan) //http再次初始化负载的host,port
					httpChan <- value
				}(value)

				go func(value string) {
					//###################### ################
					//todo 启动GRPC 客户端 如果作为客户端要连其它rpc server需要开启下面，否则可注释掉
					//###################### ################
					backend.RPCClientsRun(grpcChan) //grpc再次初始化host,port
					grpcChan <- value
				}(value)

			}
		}()

		//***********************************************************
		//第四步 通过服务发现要使用的其它服务，并自动watch其状态的改变，先初始化
		//这是服务发现，必须与第三步同时使用
		//***********************************************************
		otherService := config.Conf.OtherServices
		err = registryAndDiscover.Discovery(otherService) //err=nil表示并发执行服务发现成功并监听ing...
		for _, value := range otherService {              //初始化
			watch <- value
		}

	}

	//这是测试可取消
	//TestLB()

	//***********************************************************
	//第五步 run起服务自己并监听当服务down之前，执行某些操作
	//***********************************************************
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		fmt.Println("######origin, this server shutdown by Iris, you can do something from here ...######")
		// 关闭所有主机
		_ = app.Shutdown(ctx)
	})
	_ = app.Run(iris.Addr(":"+strconv.Itoa(config.Conf.ServerPort), func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			//注销掉当前服务 unregistry
			err = registryAndDiscover.UnRegistry()
			if err != nil {
				zgo.Log.Error(err)
			}
		})

	}), iris.WithoutInterruptHandler)
}

//func WatchSignal() {
//	//创建监听退出chan
//	signalChan := make(chan os.Signal)
//	//监听指定信号 ctrl+c kill
//	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
//	go func() {
//		for s := range signalChan {
//			switch s {
//			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
//				go func() {
//					wg.Add(1)
//					defer wg.Done()
//
//					fmt.Println("---i am out over from killed cmd, but not kill -9, you can do something ...---")
//
//				}()
//			case syscall.SIGUSR1:
//				//todo something
//				fmt.Println("usr1", s)
//			case syscall.SIGUSR2:
//				//todo something
//
//				fmt.Println("usr2", s)
//			default:
//				//todo something
//
//				fmt.Println("other", s)
//			}
//		}
//	}()
//}

func TestLB() {
	//以下为测试使用，通过内部负载均衡使用其它服务
	go func() {
		for {
			//每次LB会动态改变config.Conf中的host变量
			fmt.Println("DemoHostForPayCanChangeAnyName: ", config.Conf.DemoHostForPayCanChangeAnyName)

			time.Sleep(2 * time.Second)
			lbRes, err := zgo.Service.LB(config.Conf.ServiceInfo.SvcName)
			if err != nil {
				zgo.Log.Error(err)
				return
			}
			//fmt.Printf("host: %s, HttpPort: %s, GrpcPort: %s\n", lbRes.SvcHost, lbRes.SvcHttpPort, lbRes.SvcGrpcPort)
			ul := fmt.Sprintf("http://%s:%s", lbRes.SvcHost, lbRes.SvcHttpPort)
			_, err = zgo.Http.Get(ul)
			if err != nil {
				zgo.Log.Error(err)
				continue
			}
			zgo.Log.Infof("请求http: %s, 200\n", ul)

			//测试call rpc
			//backend.CallRpcHelloworld()

		}
	}()
}
