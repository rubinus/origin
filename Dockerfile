FROM alpine:latest
LABEL maintainer="zhuhonglei@zhuge.com"


#设置国内镜像
RUN echo 'https://mirrors.ustc.edu.cn/alpine/latest-stable/community' > /etc/apk/repositories
RUN echo 'https://mirrors.ustc.edu.cn/alpine/latest-stable/main' >> /etc/apk/repositories
RUN apk update

#访问https
RUN apk add --no-cache ca-certificates

#设置东八区，北京时间
ENV TZ=Asia/Shanghai
RUN apk add --no-cache tzdata && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


#设置工作目录
WORKDIR $GOPATH/src/base-to-base-wait-copy/
#ADD ./ $GOPATH/src/base-to-base-wait-copy/
ADD base-to-base-wait-copy $GOPATH/src/base-to-base-wait-copy/

ADD config $GOPATH/src/base-to-base-wait-copy/config
ADD views $GOPATH/src/base-to-base-wait-copy/views
ADD public $GOPATH/src/base-to-base-wait-copy/public
ADD deploy $GOPATH/src/base-to-base-wait-copy/deploy


RUN ["chmod", "+x", "base-to-base-wait-copy"]

EXPOSE 80
EXPOSE 50051

ENTRYPOINT ["./base-to-base-wait-copy","--env","dev"]