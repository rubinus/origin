package main

import (
	"context"
	goflag "flag"
	"fmt"
	"github.com/gitcpu-io/origin/configs"
	"github.com/gitcpu-io/origin/engine"
	"github.com/gitcpu-io/origin/grpcclient"
	"github.com/gitcpu-io/origin/grpcserver"
	"github.com/gitcpu-io/origin/routes"
	"github.com/gitcpu-io/zgo"
	"github.com/kataras/iris/v12"
	flag "github.com/spf13/pflag"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

const (
	DefaultGrpcPort int = 50051
)

var (
	cpath          string
	env            string
	project        string
	etcdAddress    string
	port           int
	rpcPort        int
	svcName        string
	svcHost        string
	svcHttpPort    string
	svcGrpcPort    string
	svcEtcdAddress string
)

func init() {
	flag.StringVar(&cpath, "cpath", "", "当不使用etcd作为配置中心时，通过配置文件cpath和env=local或env=container一起使用")

	flag.StringVar(&env, "env", "local", "start local/dev/qa/pro env config，本机开发使用local，打包容器后使用container")

	flag.StringVar(&project, "project", "", "create project id by zgo engine admin")

	flag.StringVar(&etcdAddress, "etcdAddress", "", "输入IP:PORT,IP:PORT指定配置中心etcd的host")

	flag.IntVar(&port, "port", 0, "http port")

	flag.IntVar(&rpcPort, "rpcPort", DefaultGrpcPort, "grpc port")

	//暴露服务信息
	flag.StringVar(&svcName, "svc_name", "", "让服务对外可访问的名称")

	flag.StringVar(&svcHost, "svc_host", "", "让服务对外可访问的主机地址，默认是宿主机的内网IP")

	flag.StringVar(&svcHttpPort, "svc_http_port", "", "让服务http对外可访问的端口号，默认是80")

	flag.StringVar(&svcGrpcPort, "svc_grpc_port", "", "让服务grpc对外可访问的端口号，默认是50051")

	flag.StringVar(&svcEtcdAddress, "svc_etcd_address", "", "让服务可以用外部的注册中心地址，默认与zgo engine相同")

	//====解析入参，并打印出来====
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()
	err := goflag.CommandLine.Parse([]string{})
	if err != nil {
		panic(err)
	}
	var inParams = make(map[string]string)
	flag.VisitAll(func(f *flag.Flag) {
		inParams[f.Name] = f.Value.String()
	})
	fmt.Println("Input args:", inParams)
	fmt.Println()
	//====结束入参处理====

	if os.Getenv("ENV") != "" { //从os的env取得ENV，用来在yaml文件中的配置，决定使用哪个*.json配置
		env = os.Getenv("ENV")
	}

	// 显示runtime信息
	if runtime.GOOS != "windows" && env != "local" {
		c := exec.Command("sh", "-c", "sh ./entrypoint.sh")
		output, err := c.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(output))
	}

	//load configs from dev/qa/pro
	if cpath == "" {
		pwd, err := os.Getwd()
		if err == nil {
			cpath = fmt.Sprintf("%s/%s", pwd, "configs")
		}
	}
	configs.InitConfig(cpath, env, project, etcdAddress, port, rpcPort)

	if os.Getenv("PROJECT") != "" {
		configs.Conf.Project = os.Getenv("PROJECT") //从os的env取得PROJECT，用来在yaml文件中的配置
	}
	if os.Getenv("ETCDADDRESS") != "" {
		configs.Conf.EtcdAddress = os.Getenv("ETCDADDRESS") //从os的env取得ETCDADDRESS，用来在yaml文件中的配置
	}
	if os.Getenv("PORT") != "" {
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		configs.Conf.HttpPort = port //从os的env取得PORT，用来在yaml文件中的配置
	}
	if os.Getenv("RPCPORT") != "" {
		atoi, err := strconv.Atoi(os.Getenv("RPCPORT"))
		if err != nil {
			atoi = DefaultGrpcPort
		}
		configs.Conf.RpcPort = atoi //从os的env取得RPCPORT，用来在yaml文件中的配置
	}

	//用docker配置覆盖服务注册与发现的配置
	if os.Getenv("SVC_NAME") != "" {
		configs.Conf.ServiceInfo.SvcName = os.Getenv("SVC_NAME") //来os的env，用来在yaml文件中的配置
	}
	if os.Getenv("SVC_HOST") != "" {
		configs.Conf.ServiceInfo.SvcHost = os.Getenv("SVC_HOST") //来os的env，用来在yaml文件中的配置
	}
	if os.Getenv("SVC_HTTP_PORT") != "" {
		configs.Conf.ServiceInfo.SvcHttpPort = os.Getenv("SVC_HTTP_PORT") //来os的env，用来在yaml文件中的配置
	}
	if os.Getenv("SVC_GRPC_PORT") != "" {
		configs.Conf.ServiceInfo.SvcGrpcPort = os.Getenv("SVC_GRPC_PORT") //来os的env，用来在yaml文件中的配置
	}
	if os.Getenv("SVC_ETCD_ADDRESS") != "" {
		configs.Conf.ServiceInfo.SvcEtcdAddress = os.Getenv("SVC_ETCD_ADDRESS") //来os的env，用来在yaml文件中的配置
	}

	//Args输入会覆盖ENV及配置中的xxxx.json中的变量
	if project != "" {
		configs.Conf.Project = project
	}
	if etcdAddress != "" {
		configs.Conf.EtcdAddress = etcdAddress
	}
	if port != 0 {
		configs.Conf.HttpPort = port
	}
	if rpcPort != 0 {
		configs.Conf.RpcPort = rpcPort
	}
	if project != "" {
		configs.Conf.Project = project
	}
	if svcName != "" {
		configs.Conf.ServiceInfo.SvcName = svcName
	}
	if svcHost != "" {
		configs.Conf.ServiceInfo.SvcHost = svcHost
	}
	if svcHttpPort != "" {
		configs.Conf.ServiceInfo.SvcHttpPort = svcHttpPort
	}
	if svcGrpcPort != "" {
		configs.Conf.ServiceInfo.SvcGrpcPort = svcGrpcPort
	}
	if svcEtcdAddress != "" {
		configs.Conf.ServiceInfo.SvcEtcdAddress = svcEtcdAddress
	}

	if configs.Conf.ServiceInfo.SvcHttpPort == "" {
		configs.Conf.ServiceInfo.SvcHttpPort = fmt.Sprintf("%d", configs.Conf.HttpPort)
	}
	if configs.Conf.ServiceInfo.SvcGrpcPort == "" {
		configs.Conf.ServiceInfo.SvcGrpcPort = fmt.Sprintf("%d", configs.Conf.RpcPort)
	}

	fmt.Println()
	structToMap := zgo.Utils.StructToMap(&configs.Conf)
	fmt.Printf("应用到进程的配置项总共: %d 个\n", len(structToMap))
	for idx, val := range structToMap {
		fmt.Println(idx, ": ", val)
	}
	fmt.Println()
	fmt.Println()
}

func main() {
	err := engine.Run() //start zgo engine
	if err != nil {
		panic(err)
	}

	app := iris.New() //start web http server
	app.Logger().SetLevel(configs.Conf.Loglevel)

	var pre string
	if configs.Conf.UsePreAbsPath == 1 {
		prefix, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "/")
		pre = prefix + "/web"
	} else {
		pre = "./web"
	}

	app.HandleDir("/", "./web/static") //static

	app.RegisterView(iris.HTML(pre, ".html").Reload(configs.Conf.IrisMod)) // select the html engine to serve templates

	//集中调用路由
	routes.Index(app)

	//消费nsq 需要先配置上nsq
	//queue_pop.NsqConsumer()
	//消费kafka 需要先配置上kafka
	//queue_pop.KafkaConsumer()
	//消费Rabbitmq
	//queue_pop.RabbitmqConsumer() //需要先配置上rabbitmq

	go func() { //start grpc grpcserver on the default port 50051 如果作为rpc服务端，让其它client连接进来
		grpcserver.Start()
	}()

	//用于pprof server分析性能
	go func() {
		fmt.Printf("Now listening pprof Server on: http://%s:%d/debug/pprof\n", zgo.Utils.GetIntranetIP(), configs.Conf.PprofPort)
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", configs.Conf.PprofPort), nil)
		if err != nil {
			zgo.Log.Error(err)
			return
		}
	}()

	if !configs.Conf.StartService {
		//正常启动服务，不使用服务注册与发现，标准模式
		normalStart(app)
	} else {
		//使用服务注册与服务发现模式
		useServiceRegistryDiscover(app)
	}

}

func normalStart(app *iris.Application) {
	//###################### ################
	//todo 启动GRPC 客户端 如果作为客户端要连其它rpc server需要开启下面，否则可注释掉
	//###################### ################
	grpcclient.RPCClientsRun(nil) //start grpc client

	//run起自己
	iris.RegisterOnInterrupt(func() { //优雅的退出 -- 使用iris框架中的退出
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		fmt.Println()
		fmt.Println("######origin, this grpcserver is normal shutdown by Iris, you can do something from here ...######")
		fmt.Println()
		// 关闭所有主机
		configs.Goodbye()
		_ = app.Shutdown(ctx)
	})
	_ = app.Run(iris.Addr(":"+strconv.Itoa(configs.Conf.HttpPort), func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {

		})
	}), iris.WithoutInterruptHandler)
}

func useServiceRegistryDiscover(app *iris.Application) {
	//注册服务到 注册中心etcd中，然后监听当前服务使用的其它服务的名字

	//***********************************************************
	//第一步 必须
	//***********************************************************
	registryAndDiscover, err := zgo.Service.New(6,
		configs.Conf.ServiceInfo.SvcEtcdAddress)
	if err != nil {
		zgo.Log.Errorf("创建微服务实例化失败 %v", err)
		return
	}

	//***********************************************************
	//第二步 配置文件决定是否开启使用，这一步很重要，一定要注册外部可以访问当前服务的，尤其用docker时要注意
	//***********************************************************
	//todo 请确认下面三项 host httpport grpcport 使其它服务可访问到
	if configs.Conf.StartServiceRegistry {
		var host string
		if configs.Conf.ServiceInfo.SvcHost == "" { //默认为空使用宿主机内部IP
			host = zgo.Utils.GetIntranetIP()
		} else {
			host = configs.Conf.ServiceInfo.SvcHost
		}
		shport, err := strconv.Atoi(configs.Conf.ServiceInfo.SvcHttpPort)
		if err != nil {
			shport = configs.Conf.HttpPort
		}
		sgport, err := strconv.Atoi(configs.Conf.ServiceInfo.SvcGrpcPort)
		if err != nil {
			sgport = configs.Conf.RpcPort
		}
		err = registryAndDiscover.Registry(configs.Conf.ServiceInfo.SvcName, host,
			shport, sgport)
		//注册当前服务(自己)到注册中心
		if err != nil {
			zgo.Log.Errorf("%s 注册微服务失败 %v", "test", err)
			return
		}
	}

	//***********************************************************
	//第三步 配置文件决定是否开启使用，这是服务发现的监听，必须与第四步同时使用
	//***********************************************************
	if configs.Conf.StartServiceDiscover {
		watch := zgo.Service.Watch()
		httpChan := make(chan string, 1000)
		grpcChan := make(chan string, 1000)
		go func() {
			for value := range watch {
				go func(value string) {
					configs.WatchHttpConfigByService(httpChan) //http再次初始化负载的host,port
					httpChan <- value
				}(value)

				go func(value string) {
					//###################### ################
					//todo 启动GRPC 客户端 如果作为客户端要连其它rpc server需要开启下面，否则可注释掉
					//###################### ################
					grpcclient.RPCClientsRun(grpcChan) //grpc再次初始化host,port
					grpcChan <- value
				}(value)

			}
		}()

		//***********************************************************
		//第四步 通过服务发现要使用的其它服务，并自动watch其状态的改变，先初始化
		//这是服务发现，必须与第三步同时使用
		//***********************************************************
		otherService := configs.Conf.OtherServices
		err = registryAndDiscover.Discovery(otherService) //err=nil表示并发执行服务发现成功并监听ing...
		for _, value := range otherService {              //初始化
			watch <- value
		}

	}

	//***********************************************************
	//第五步 run起服务自己并监听当服务down之前，执行某些操作
	//***********************************************************
	iris.RegisterOnInterrupt(func() { //优雅的退出 -- 使用iris框架中的退出
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		//注销掉当前服务 unregistry
		err = registryAndDiscover.UnRegistry()
		if err != nil {
			zgo.Log.Error(err)
		}
		fmt.Println()
		fmt.Println("######origin, this grpcserver use the register/discover shutdown by Iris, you can do something from here ...######")
		fmt.Println()
		// 关闭所有主机
		configs.Goodbye()
		_ = app.Shutdown(ctx)
	})
	_ = app.Run(iris.Addr(":"+strconv.Itoa(configs.Conf.HttpPort), func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			//注销掉当前服务 unregistry
			err = registryAndDiscover.UnRegistry()
			if err != nil {
				zgo.Log.Error(err)
			}
		})
	}), iris.WithoutInterruptHandler)
}
