FROM alpine:latest
LABEL maintainer="zhuhonglei@zhuge.com"

#设置国内镜像
RUN echo 'https://mirrors.ustc.edu.cn/alpine/latest-stable/community' > /etc/apk/repositories
RUN echo 'https://mirrors.ustc.edu.cn/alpine/latest-stable/main' >> /etc/apk/repositories
RUN apk update

#访问外部的https
RUN apk add --no-cache ca-certificates

#设置东八区，北京时间
ENV TZ=Asia/Shanghai
RUN apk add --no-cache tzdata && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


#设置工作目录
WORKDIR $GOPATH/src/origin/

#添加可执行文件
ADD origin $GOPATH/src/origin/
ADD entrypoint.sh $GOPATH/src/origin/

#添加配置及html等
ADD config $GOPATH/src/origin/config
ADD views $GOPATH/src/origin/views
ADD public $GOPATH/src/origin/public
ADD deploy $GOPATH/src/origin/deploy


RUN ["chmod", "+x", "origin"]
RUN ["chmod", "+x", "entrypoint.sh"]

#设置Web端口，一般不用更改
EXPOSE 80

#设置GRPC端口，一般不用更改
EXPOSE 50051

#ENTRYPOINT ["./origin","--env","dev"]
ENTRYPOINT ["./entrypoint.sh"]