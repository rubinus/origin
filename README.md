# origin v1.1.1

# 本地运行docker-compose 启动中间件db/cache/queue等

执行 docker-compose up 或 docker-compose up -d

docker-compose ps

## Redis默认使用 6379 端口

## Mongo默认使用 27018 端口

docker exec -it mongo27018 sh

mongo

use admin

```javascript
db.createUser(
  {
    user: "admin",
    pwd: "admin",
    roles: [ { role: "root", db: "admin" } ]
  }
)
```

db.auth('admin','admin')

### 创建天气的db，准备测试
use weather

### 插入测试数据

use profile

```javascript
for(var i=100;i<=200;i++){ db.bj.insert({ username: 'zhangsan', age:Math.round(Math.random() * 100), address:Math.round(
Math.random() * 100), }); }
```

# 关于配置文件 _init/xxxx.json 和 config/xxxx.json 的说明

## _init目录，项目初始化（如果使用etcd作为配置中心的话运行下面的，如果是用本地配置文件的方式就不用）
```shell
cd origin/_init

go run main.go
```

- 其中_init目录下 local.json 只有redis和mongodb

- 如果使用docker运行，要确保_init/local.json中的 中间件 host为可访问到的地址，不能使用localhost

- _init/local.json这个配置文件会直接作用到zgo engine的配置上

### all_db_local_back.json 可以复制内容到 local.json 中，有全量的中间件可以使用，具体要看docker-compose运行了多少个

## config目录中的配置文件如何使用？

> local.json适合本地调试开发 8081端口，仍然使用原生的方式来连接各种db，以配置文件的方式在ECS机器上，在此以redis和mongodb为例

- 当使用local.json本地开发时，还需要在 engine/zgo.go中指定使用的中间件的key, 你可能使用一个或多个中间件

> 如果使用 dev.json/qa.json/pro.json/k8s.json就会使用etcd做为配置库 默认80端口，其中的数据存储格式就是_init/local.json中部分的实例，

- 当然真实的 etcd中还存有其它mysql/mongo/kafka等等的配置文件

> container.json仅供打包测试image时 默认是80端口

- 如果不使用etcd做为配置中心，又要docker build，那么最好的方式是使用container.json进行相应的参数变更，仅供测试image

- 最佳实践，如果部署到不同的k8s环境，可以使用dev/qa/pro/k8s相应的配置并修改

- 如果使用etcdAddress参数，支持etcd集群，多节点用,隔开(ip:port,ip:port,ip:port); args参数中的etcdAddress会覆盖dev/qa/pro/k8s.json中的etcdAddress

> local.json 和 container.json 经过zgo engine时不会读取配置中心etcd中的值，仅使用.json中的配置，意味着你可以把 中间件的信息放置在这2个配置文件中，模拟存储在etcd中的配置信息

# 安装gogo proto
go get github.com/gogo/protobuf/protoc-gen-gofast

# auto build image 将会使用container模式

通过makefile，运行dockerfile，制作包含git版本的image

```shell
make image
```

# 关于运行origin

## 在本地执行打包好的镜像origin

docker run --rm -p 8081:80 -p 8181:8181 -p 50051:50051 -d --name origin rubinus/origin:v1.0

## 如果使用了etcd为配置中心，需要带上参数 etcdAddress，同时env的取值需要是dev级别以上，不能再用local，如下所示

- 本机开发

go run main.go --env dev --etcdAddress localhost:3379

- 本机开发最佳实践，直接修改config/local.json中 env=dev 和 etcdAddress的值

- Docker运行时指定 env=dev etcdAddress参数为配置中心etcd的地址host:port（容器可访问到的IP）

docker run --rm -p 8081:80 -p 8181:8181 -p 50051:50051 --name origin rubinus/origin:v1.0 --env=dev --etcdAddress=192.168.110.173:3379

# 访问origin

## Post保存天气信息
```shell
curl -l -H "Content-type: application/json" -X POST -d '{"query":"深圳市"}' "http://localhost:8081/v1/weather/put"

```

## 查询天气列表

```shell
curl -l "http://localhost:8081/v1/weather/list?city=深圳市"

```

## 通过grpc来查询天气列表，当启动origin后，在另外一个terminal中执行
go run grpcclient-test/weather/main.go


# How to use the zgo engine

## Http

//前端ajax-->main.go(Run)-->routes-->(实际业务处理handler)-->services-->zgo.组件(mysql/mongo/redis)-->models(库)

请参照：routes对应的handlers中的regis.go来写接口

## Grpc

grpcserver是grpc服务端实现与启动

grpchandlers是类似与 handler的处理grpc的handler，服务端的实现

grpcclients是grp客户端的连接与封装

grpcclient-test是grpc模拟发送客户端

queue_push 是队列生产端

queue_pop 是队列消费端

git 复制这个项目后，改名成自己开发的项目名字（全局替换：origin），这是一个模板，内含有samples目录

安装docker,在本地一次性跑起redis,mongodb,mysql,nsq,kafka

# ========

## origin测试方法使用：建立xxx_test.go文件，生成相应的.out，并通过go tool pprof查看

### 查看测试代码覆盖率

go test -coverprofile=c.out

go tool cover -html=c.out

### 查看测试代码trace

go test -trace=t.out

go tool trace t.out

### 查看cpu使用

go test -bench . -cpuprofile cpu.out

go tool pprof -http=":8081" cpu.out

### 查看内存使用

go test -memprofile mem.out

go tool pprof -http=":8081" mem.out

#### 执行pprof后，然后输入web 或是quit 保证下载了svg

https://graphviz.gitlab.io/_pages/Download/Download_source.html

下载graphviz-2.40.1后进入目录

./configure

make

make install

## =====启动 go run main.go 查看web服务下的pprof输入web=====

### 图形报告

http://localhost:8181/debug/pprof/

### 使用pprof查看所有gorutines

go tool pprof http://localhost:8181/debug/pprof/goroutine?debug=1

### 使用pprof查看堆内存分配

go tool pprof http://localhost:8181/debug/pprof/heap

### 使用pprof查看10秒CPU使用

go tool pprof http://localhost:8181/debug/pprof/profile?seconds=10

### 使用go tool trace查看trace

wget -O trace.out http://localhost:8181/debug/pprof/trace?seconds=10

go tool trace trace.out

## ========

# 编译文件mac或linux

## 选项一：在当前目录下编译mac运行的二进制文件，仅适用于本机运行

go build -o origin

## 查看逃逸分析

go build -gcflags '-m -l' -o origin

## 使用godebug查看

GODEBUG=scheddetail=1,schedtrace=1000,gctrace=1 ./origin

## 使用godebug 直接运行main.go

GODEBUG=scheddetail=1,schedtrace=1000,gctrace=1 go run main.go

## 选项二：在当前目录下编译linux运行的二进制文件，适用于服务器linux环境

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o origin

# 用docker制作image

### 本机build

docker build -t cr.gitcpu-io/origin:v1.0 .

docker push cr.gitcpu-io/origin:v1.0

### 在服务器上执行

docker pull cr.gitcpu-io/origin:v1.0

docker rm -f origin

# ====本机运行====

#### 下面一行非服务注册模式

docker run -d --restart always -p 8081:80 -p 50051:50051 --name origin cr.gitcpu-io/origin:v1.0

### 作为服务注册(本地)

docker run -d --restart always -p 8081:8081 -p 50051:50051 -e SVC_HOST=192.168.100.19 -e SVC_HTTP_PORT=8081 -e
SVC_GRPC_PORT=50051 --name origin1 cr.gitcpu-io/origin:v1.0

### 再启动一个（仅更换端口号）模拟正式环境

docker run -d --restart always -p 8082:8082 -p 50052:50051 -e SVC_HOST=192.168.100.19 -e SVC_HTTP_PORT=8082 -e
SVC_GRPC_PORT=50052 --name origin2 cr.gitcpu-io/origin:v1.0

# ====服务器上运行====

## 正常运行

docker run -d --restart always -p 8081:80 -p 50051:50051 --name origin cr.gitcpu-io/origin:v1.0

## 在开发服务器上启动docker并指定 svc 服务的访问host及port(服务器上使用服务注册模式)

docker run -d --restart always -p 8281:8281 -p 50051:50051 -e SVC_HOST=localhost -e SVC_HTTP_PORT=8281 -e
SVC_GRPC_PORT=50051 --name origin1 cr.gitcpu-io/origin:v1.0

docker run -d --restart always -p 8282:8282 -p 50052:50051 -e SVC_HOST=localhost -e SVC_HTTP_PORT=8282 -e
SVC_GRPC_PORT=50052 --name origin2 cr.gitcpu-io/origin:v1.0

docker logs -f --tail=20 origin
