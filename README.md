# origin

## config/local.json
local.json适合本地调试开发，仍然使用原生的方式来连接各种db，部署以配置文件的方式在ECS机器上，在此以redis为例
如果使用
dev.josn/qa.json/pro.json就会使用etcd做为库，其中的数据存储格式就是local.json中redis部分的实例，当然真实的
etcd中还存有其它mysql/mongo/kafka等等的配置文件

当使用local.json本地开发时，还需要在 engine/zgo.go中指定使用的中间件的key

## auto build image
make image
通过makefile，运行dockerfile，制作包含git版本的image

#### 在本地执行打包好的镜像origin并使用etcd
docker run --rm -p 8081:80 -p 8181:8181 -d --name origin rubinus/origin:1.0.0

##how to use the zgo engine

##deploy文件目录是用运维用来部署k8s和istio的，其中的yaml文件需要由开发人员编写

##Http
//前端ajax-->main.go(Run)-->routes-->(实际业务处理handler)-->services-->zgo.组件(mysql/mongo/redis/pika)-->models(库)
请参照：routes对应的handlers中的regis.go来写接口

##Grpc
grpcserver是grpc服务端实现与启动

grpchandlers是类似与 handler的处理grpc的handler，服务端的实现

grpcclients是grp客户端的连接与封装

grpcclient-test是grpc模拟发送客户端

queue_push 是队列生产端

queue_pop 是队列消费端

git 复制这个项目后，改名成自己开发的项目名字（全局替换：origin），这是一个模板，内含有samples目录

安装docker,在本地一次性跑起redis,mongodb,mysql,nsq,kafka

#========
##origin测试方法使用：建立xxx_test.go文件，生成相应的.out，并通过go tool pprof查看

###查看测试代码覆盖率

go test -coverprofile=c.out

go tool cover -html=c.out

###查看测试代码trace

go test -trace=t.out

go tool trace t.out

###查看cpu使用

go test -bench . -cpuprofile cpu.out

go tool pprof -http=":8081" cpu.out

###查看内存使用

go test -memprofile mem.out

go tool pprof -http=":8081" mem.out

####执行pprof后，然后输入web  或是quit 保证下载了svg

https://graphviz.gitlab.io/_pages/Download/Download_source.html

下载graphviz-2.40.1后进入目录

./configure

make

make install

##=====启动go run main.go 查看web服务下的pprof输入web=====
####图形报告
http://localhost:8181/debug/pprof/

####使用pprof查看所有gorutines
go tool pprof http://localhost:8181/debug/pprof/goroutine?debug=1

####使用pprof查看堆内存分配
go tool pprof http://localhost:8181/debug/pprof/heap

####使用pprof查看10秒CPU使用
go tool pprof http://localhost:8181/debug/pprof/profile?seconds=10

####使用go tool trace查看trace
wget -O trace.out http://localhost:8181/debug/pprof/trace?seconds=10

go tool trace trace.out

###========
#编译文件mac或linux
##选项一：在当前目录下编译mac运行的二进制文件，仅适用于本机运行
go build -o origin

####查看逃逸分析
go build -gcflags '-m -l' -o origin

####使用godebug查看
GODEBUG=scheddetail=1,schedtrace=1000,gctrace=1 ./origin

####使用godebug 直接运行main.go
GODEBUG=scheddetail=1,schedtrace=1000,gctrace=1 go run main.go

##选项二：在当前目录下编译linux运行的二进制文件，适用于服务器linux环境
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o origin

用docker制作image(dck.example.test是任意一个标识，如果愿意你可以改为你的名字，每一次v0.0.1需要递增)
###本机build
docker build -t dck.example.test/origin:v0.0.1 .

docker push dck.example.test/origin:v0.0.1

###在服务器上执行
docker pull dck.example.test/origin:v0.0.1

docker rm -f origin

#====本机运行====
####下面一行非服务注册模式
docker run -d --restart always -p 8080:80 -p 50051:50051 --name origin dck.example.test/origin:v0.0.1

###作为服务注册(本地)
docker run -d --restart always -p 8081:80 -p 51051:50051 -e SVC_HOST=192.168.100.19 -e SVC_HTTP_PORT=8081 -e SVC_GRPC_PORT=51051 --name origin1 dck.example.test/origin:v0.0.1

###再启动一个（仅更换端口号）模拟正式环境
docker run -d --restart always -p 8082:80 -p 51052:50051 -e SVC_HOST=192.168.100.19 -e SVC_HTTP_PORT=8082 -e SVC_GRPC_PORT=51052 --name origin2 dck.example.test/origin:v0.0.1

#====服务器上运行====
##正常运行
docker run -d --restart always -p 8080:80 -p 50051:50051 --name origin dck.example.test/origin:v0.0.1

##在开发服务器上启动docker并指定 svc 服务的访问host及port(服务器上使用服务注册模式)
docker run -d --restart always -p 8281:80 -p 52051:50051 -e SVC_HOST=localhost -e SVC_HTTP_PORT=8281 -e SVC_GRPC_PORT=52051 --name origin1 dck.example.test/origin:v0.0.1

docker run -d --restart always -p 8282:80 -p 52052:50051 -e SVC_HOST=localhost -e SVC_HTTP_PORT=8282 -e SVC_GRPC_PORT=52052 --name origin2 dck.example.test/origin:v0.0.1

docker logs -f --tail=20 origin


服务器build (exampleprod是生产，exampledev是qa，各自2个版本)
docker build -t registry.cn-beijing.aliyuncs.com/exampleprod/origin:v1.0.0.1 .
docker build -t registry.cn-beijing.aliyuncs.com/exampledev/origin:v1.0.0.1 .

docker build -t registry.cn-beijing.aliyuncs.com/exampleprod/origin:v1.0.0.2 .
docker build -t registry.cn-beijing.aliyuncs.com/exampledev/origin:v1.0.0.2 .

push到阿里云的私有镜像仓库
docker push registry.cn-beijing.aliyuncs.com/exampleprod/origin:v1.0.0.1
docker push registry.cn-beijing.aliyuncs.com/exampledev/origin:v1.0.0.1

docker push registry.cn-beijing.aliyuncs.com/exampleprod/origin:v1.0.0.2
docker push registry.cn-beijing.aliyuncs.com/exampledev/origin:v1.0.0.2

###======================
#本地运行docker-compose
执行
docker-compose up
或
docker-compose up -d

##origin 本机使用local时测试环境，测试服务器IP地址

##Mysql
2个mysql

3307

3308

##Mongo
2个mongo

27018

27019

use admin
db.auth('admin','admin')

### 插入测试数据
use profile

for(var i=100;i<=200;i++){
    db.bj.insert({
    username: 'zhangsan',
    age:Math.round(Math.random() * 100),
    address:Math.round(Math.random() * 100),
    });
}

####Mongo管理页面
http://localhost:8081

##Redis
2个redis

6380

6381

##Postgres
2个postgres

5433

5434

##Etcd
1个etcd

2381

####Etcd管理页面
http://47.93.163.209:9097

##Neo4j
1个neo4j

7687

####Neo4j操作页面
http://localhost:7474

连接：bolt://localhost:7687

账号：neo4j

密码：12345678

##ClickHouse
1个clickhouse

http端口 8123

tcp端口 9019

####ClickHouse管理操作页面

http://localhost:9020

输入：http://localhost:8123

login: default

password: 空

##Rabbitmq
5672

集群：25672

####rabbitmq管理页面

http://localhost:8672

管理账号/密码：admin/admin

##Kafka
1个kafka

生产：9202

消费：2081

####Kafka管理页面
http://localhost:9093

管理账号/密码：admin/admin

##Nsq
1个nsq

4150

####Nsq管理页面
http://localhost:4171


##ES
1个es

9200
####ES管理页面--打开localhost:9800后，输入http://es:9200 connect
http://localhost:9800

#####管理页面--打开http://localhost:1358，输入http://localhost:9200
后面输入index，connect

http://localhost:1358

####Kibana
http://localhost:5601

####redis 集群，任意节点支持读写
localhost:7001

localhost:7002

localhost:7003

localhost:7004

localhost:7005

localhost:7006

####1个portainer--用于查看所有docker中的资源
http://localhost:9000

账号: admin

密码: 12345678