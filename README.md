# visource


##how to use the zgo engine

##deploy文件目录是用运维用来部署k8s和istio的，其中的yaml文件需要由开发人员编写


//前端ajax-->main.go(Run)-->routes-->(实际业务处理handler)-->services-->zgo.组件(mysql/mongo/redis/pika)-->models(库)


git clone这个项目后，改名成自己开发的项目名字，然后删除掉.git目录，这是一个模板，内含有samples目录

安装docker,在本地一次性跑起redis,mongodb,mysql,nsq,kafka

###zgo start测试方法使用：进入到比如samples/demo_mongo目录下执行，生成相应的.out，并通过go tool pprof查看

// 查看测试代码覆盖率

go test -coverprofile=c.out

go tool cover -html=c.out

// 查看cpu使用

go test -bench . -cpuprofile cpu.out

go tool pprof -http=":8081" cpu.out

// 查看内存使用

go test -memprofile mem.out

go tool pprof -http=":8081" mem.out

执行pprof后，然后输入web  或是quit 保证下载了svg

https://graphviz.gitlab.io/_pages/Download/Download_source.html

下载graphviz-2.40.1后进入目录

./configure

make

make install

###======================
执行
docker-compose up
或
docker-compose up -d

选项一：在当前目录下编译mac运行的二进制文件，仅适用于本机运行
go build -o visource

选项二：在当前目录下编译linux运行的二进制文件，适用于服务器linux环境
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o visource

用docker制作image(dck.zhuge.test是任意一个标识，如果愿意你可以改为visource，每一次v1.0.0需要递增)
本机build
docker build -t dck.zhuge.test/visource:v1.0.8 .

服务器build
docker build -t registry.cn-beijing.aliyuncs.com/zhuge/visource:v1.1.6 .

push到阿里云的私有镜像仓库
docker push registry.cn-beijing.aliyuncs.com/zhuge/visource:v1.1.6


##visource 本机使用local时测试环境
阿里云内网
10.45.146.41
阿里云公网
123.56.173.28

数字是端口号，供测试zgo admin使用，跑在docker里

2个mysql
3307
3308

2个mongo
27018
27019

2个redis
6380
6381

2个postgres
5433
5434

1个etcd
2381

1个neo4j
7687
####Neo4j操作页面
http://127.0.0.1:7474

####redis监控页面,grafana,prometheus
http://127.0.0.1:3000

http://127.0.0.1:9090

1个kafka
9202

1个nsq
4150
####Nsq管理页面
http://127.0.0.1:4171/

####Etcd管理页面
http://47.93.163.209:9097/

1个es
9200
####ES管理页面--打开localhost:9800后，输入http://es:9200 connect
http://127.0.0.1:9800

#####管理页面--打开http://127.0.0.1:1358，输入http://127.0.0.1:9200
后面输入index，connect

http://127.0.0.1:1358

####1个portainer--用于查看所有docker中的资源
http://127.0.0.1:9000