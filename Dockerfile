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
WORKDIR $GOPATH/src/visource/
#ADD ./ $GOPATH/src/visource/
ADD visource $GOPATH/src/visource/

ADD config $GOPATH/src/visource/config
ADD views $GOPATH/src/visource/views
ADD public $GOPATH/src/visource/public
ADD deploy $GOPATH/src/visource/deploy


RUN ["chmod", "+x", "visource"]

EXPOSE 80

ENTRYPOINT ["./visource","--env","dev"]